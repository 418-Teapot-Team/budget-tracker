package repository

import (
	budget "budget-tracker"
	"gorm.io/gorm"
)

type listsSql struct {
	db *gorm.DB
}

func newListsSQL(db *gorm.DB) Lists {
	return &listsSql{db: db}
}

func (db *listsSql) CreateList(input *budget.List) (err error) {
	tx := db.db.Begin()

	err = tx.Create(input).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error

}

func (db *listsSql) DeleteList(listId, userId int) (err error) {
	tx := db.db.Begin()

	err = tx.Where("id = ?", listId).Where("user_id = ?", userId).Delete(&budget.List{}).Error
	if err != nil {
		return err
	}
	return tx.Commit().Error
}

func (db *listsSql) GetList(userId int, budgetType string) (lists []budget.List, err error) {
	err = db.db.Where("user_id = ?", userId).Where("type = ?", budgetType).Find(&lists).Error
	return
}
