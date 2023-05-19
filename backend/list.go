package budget

import (
	"database/sql"
	"encoding/json"
	"time"
)

const listTable = "lists"

func (List) TableName() string {
	return listTable

}

func (ListsGetter) TableName() string {
	return listTable
}

type CustomTime struct {
	sql.NullTime
}

func (ct *CustomTime) UnmarshalJSON(data []byte) error {
	// Implement your custom unmarshaling logic here
	// Parse the string value from `data` and set it in `ct.NullTime.Time`

	// Example implementation assuming the string is in RFC3339 format
	str := string(data)
	t, err := time.Parse(`"2006-01-02T15:04:05Z"`, str)
	if err != nil {
		return err
	}

	ct.Time = t
	ct.Valid = true

	return nil
}

func (lt *List) UnmarshalJSON(data []byte) error {
	type Alias List // Create an alias of the struct to avoid infinite recursion

	aux := &struct {
		CreatedAt string `json:"createdAt,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(lt),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Handle the conversion from string to sql.NullTime
	if aux.CreatedAt != "" {
		t, err := time.Parse("2006-01-02T15:04:05Z", aux.CreatedAt)
		if err != nil {
			return err
		}
		lt.CreatedAt.Time = t
		lt.CreatedAt.Valid = true
	} else {
		lt.CreatedAt.Valid = false
	}

	return nil
}

type List struct {
	ID         int          `json:"id,omitempty" gorm:"column:id"`
	UserId     int          `json:"-" gorm:"column:user_id"`
	Type       string       `json:"type" gorm:"column:type" binding:"required"`
	CategoryID int          `json:"category" gorm:"column:category_id" binding:"required"`
	Amount     float64      `json:"amount" gorm:"column:amount" binding:"required"`
	Comment    *string      `json:"comment,omitempty" gorm:"column:comment"`
	CreatedAt  sql.NullTime `json:"createdAt,omitempty" gorm:"column:created_at"`
}

type ListsGetter struct {
	Id         int          `json:"id,omitempty" gorm:"column:id"`
	UserId     int          `json:"-" gorm:"column:user_id"`
	Type       string       `json:"type,omitempty" gorm:"column:type"`
	Category   int          `json:"categoryId" gorm:"column:category_id"`
	Categories Categories   `json:"category" gorm:"foreignKey:Category;references:Id"`
	Amount     float64      `json:"amount" gorm:"column:amount"`
	Comment    *string      `json:"comment,omitempty" gorm:"column:comment"`
	CreatedAt  sql.NullTime `json:"createdAt,omitempty" gorm:"column:created_at"`
}

type FinancialData struct {
	Month     string
	IncomeNet float64
}
