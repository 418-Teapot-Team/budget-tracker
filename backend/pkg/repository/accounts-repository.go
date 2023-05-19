package repository

import (
	budget "budget-tracker"
	"fmt"
	"gorm.io/gorm"
)

type accountsSql struct {
	db *gorm.DB
}

func newAccountsSQL(db *gorm.DB) Accounts {
	return &accountsSql{db: db}
}

func (db *accountsSql) CreateAccount(input *budget.Account) (err error) {
	tx := db.db.Begin()

	err = tx.Create(input).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (db *accountsSql) GetAllAccounts(userId int, account, orderBy, sortedBy string) (list []budget.Account, err error) {
	query := db.db.Select("id, type, name, month_amount, percent, sum, DATE_FORMAT(date, '%Y-%m-%d') date, created_at").
		Order(orderBy+sortedBy).Where("user_id = ?", userId)
	if account != "" {
		query = query.Where("type = ?", account)
	}
	err = query.Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil

}

func (db *accountsSql) DeleteAccount(listId, userId int) (err error) {
	tx := db.db.Begin()

	err = tx.Where("id = ?", listId).Where("user_id = ?", userId).Delete(&budget.Account{}).Error
	if err != nil {
		return err
	}
	return tx.Commit().Error
}

func (db *accountsSql) EditAccount(input budget.Account) (err error) {

	return db.db.Save(&input).Error

}

func (db *accountsSql) GetTotalsDeposits(userId int) (float64, float64) {

	var sumResult struct {
		Payed   float64 `gorm:"column:already_received"`
		GoalSum float64 `gorm:"column:goal_sum"`
	}

	query := db.db.Table("accounts").
		Select("SUM(already_received) already_received, SUM(goal_sum) goal_sum").
		Where("user_id = ?", userId).
		Where("type = ?", "deposit")

	err := query.Find(&sumResult).Error
	if err != nil {
		fmt.Println(err)
		return 0, 0
	}

	return sumResult.Payed, sumResult.GoalSum

}
