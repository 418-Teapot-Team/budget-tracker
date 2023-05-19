package budget

import (
	"time"
)

const (
	accountTable = "accounts"
)

func (Account) TableName() string {
	return accountTable
}

type Account struct {
	ID              int       `gorm:"column:id" json:"id,omitempty"`
	Type            string    `gorm:"column:type" json:"type" binding:"required"`
	UserID          int       `gorm:"column:user_id" json:"-"`
	Name            string    `gorm:"column:name" json:"name" binding:"required"`
	MonthAmount     int       `gorm:"column:month_amount" json:"monthAmount" binding:"required"`
	Percent         float64   `gorm:"column:percent" json:"percent" binding:"required"`
	Date            string    `gorm:"column:date" json:"date" binding:"required"`
	Sum             float64   `gorm:"column:sum" json:"sum" binding:"required"`
	CreatedAt       time.Time `gorm:"column:created_at" json:"createdAt,omitempty"`
	AlreadyReceived float64   `gorm:"column:already_received" json:"-"`
	FinishDate      string    `gorm:"column:date_opened" json:"-"`
	MonthlyPayment  float64   `gorm:"column:monthly_payment" json:"-"`
	GoalSum         float64   `gorm:"column:goal_sum"`
}
