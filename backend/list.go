package budget

import "time"

const listTable = "lists"

func (List) TableName() string {
	return listTable

}

type List struct {
	Id        int       `json:"id,omitempty" gorm:"column:id"`
	UserId    int       `json:"-" gorm:"column:user_id"`
	Type      string    `json:"type,omitempty" gorm:"column:type"`
	CategoryId  string  `json:"category" gorm:"column:category_id"`
	Amount    float64   `json:"amount" gorm:"column:amount"`
	Comment   *string   `json:"comment,omitempty" gorm:"column:comment"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
}
