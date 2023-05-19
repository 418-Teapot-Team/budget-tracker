package repository

import (
	budget "budget-tracker"
	"gorm.io/gorm"
	"time"
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

func (db *listsSql) GetList(userId int, budgetType, orderBy, sortedBy string, takeAmount int, skipAmount int) (lists []budget.ListsGetter, err error) {
	query := db.db.Order(orderBy+sortedBy).Where("user_id = ?", userId)

	if budgetType != "" {
		query = query.Where("type = ?", budgetType)
	}

	if takeAmount != 0 {
		query = query.Limit(takeAmount)
	}

	if skipAmount != 0 {
		query = query.Offset(skipAmount)
	}

	err = query.Preload("Categories").Find(&lists).Error

	if len(lists) == 0 {

	}

	return
}

func (db *listsSql) GetTopCategories(userId int, enum string) (lists []budget.ListsGetter, err error) {
	query := db.db.Select("category_id , sum(amount) amount").
		Where("user_id = ?", userId).
		Where("type = ?", enum).
		Group("category_id").
		Order("2 DESC")

	err = query.Preload("Categories").Find(&lists).Error
	return
}

func (db *listsSql) EditList(input budget.List) (err error) {
	return db.db.Save(&input).Error
}

func (db *listsSql) GetCurrentMonthSavings(userId int) (result float64, err error) {
	query := db.db.Table("lists").
		Select("IFNULL(SUM(CASE WHEN type = 'income' THEN amount ELSE -amount END),0)").
		Where("MONTH(created_at) = ? AND YEAR(created_at) = ?", time.Now().Month(), time.Now().Year()).
		Where("user_id = ?", userId)

	err = query.Scan(&result).Error
	return
}

func (db *listsSql) GetSavingsStats(userId int) (data []budget.FinancialData, err error) {
	query := db.db.Table("lists").
		Select("DATE_FORMAT(created_at, '%Y-%m') AS month, SUM(CASE WHEN type = 'income' THEN amount ELSE -amount END) AS income_net").
		Where("created_at >= ?", time.Now().AddDate(-1, 0, 0)).
		Where("user_id = ?", userId).
		Group("month").
		Order("month")

	err = query.Find(&data).Error

	if err != nil {
		return data, err
	}

	return data, nil
}
