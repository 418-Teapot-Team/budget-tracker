package budget

import (
	"time"
)

const listTable = "lists"

func (List) TableName() string {
	return listTable

}

type List struct {
	Id        int       `json:"id,omitempty" gorm:"column:id"`
	UserId    int       `json:"-" gorm:"column:user_id"`
	Type      string    `json:"type" gorm:"column:type"`
	Category  string    `json:"category" gorm:"column:category"`
	Amount    float64   `json:"amount" gorm:"column:amount"`
	Comment   *string   `json:"comment,omitempty" gorm:"column:comment"`
	CreatedAt time.Time `json:"-" gorm:"column:created_at"`
}
