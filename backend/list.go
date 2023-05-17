package budget

import "time"

const listTable = "lists"

func (List) TableName() string {
	return listTable

}

func (ListsGetter) TableName() string {
	return listTable
}

type List struct {
	Id        int       `json:"id,omitempty" gorm:"column:id"`
	UserId    int       `json:"-" gorm:"column:user_id"`
	Type      string    `json:"type" gorm:"column:type"`
	Category  int       `json:"category" gorm:"column:category_id"`
	Amount    float64   `json:"amount" gorm:"column:amount"`
	Comment   *string   `json:"comment,omitempty" gorm:"column:comment"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
}

type ListsGetter struct {
	Id         int        `json:"id,omitempty" gorm:"column:id"`
	UserId     int        `json:"-" gorm:"column:user_id"`
	Type       string     `json:"type,omitempty" gorm:"column:type"`
	Categories Categories `json:"category" gorm:"foreignKey:id;references:Id"`
	Amount     float64    `json:"amount" gorm:"column:amount"`
	Comment    *string    `json:"comment,omitempty" gorm:"column:comment"`
	CreatedAt  time.Time  `json:"created_at,omitempty" gorm:"column:created_at"`
}
