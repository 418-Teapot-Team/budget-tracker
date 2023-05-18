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
	Id        int        `json:"id,omitempty" gorm:"column:id"`
	UserId    int        `json:"-" gorm:"column:user_id"`
	Type      string     `json:"type" gorm:"column:type" binding:"required"`
	Category  int        `json:"category" gorm:"column:category_id" binding:"required"`
	Amount    float64    `json:"amount" gorm:"column:amount" binding:"required"`
	Comment   *string    `json:"comment,omitempty" gorm:"column:comment"`
	CreatedAt *time.Time `json:"createdAt,omitempty" gorm:"column:created_at"`
}

type ListsGetter struct {
	Id         int        `json:"id,omitempty" gorm:"column:id"`
	UserId     int        `json:"-" gorm:"column:user_id"`
	Type       string     `json:"type,omitempty" gorm:"column:type"`
	Category   int        `json:"categoryId" gorm:"column:category_id"`
	Categories Categories `json:"category" gorm:"foreignKey:Category;references:Id"`
	Amount     float64    `json:"amount" gorm:"column:amount"`
	Comment    *string    `json:"comment,omitempty" gorm:"column:comment"`
	CreatedAt  *time.Time `json:"createdAt,omitempty" gorm:"column:created_at"`
}
