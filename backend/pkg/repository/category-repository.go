package repository

import (
	budget "budget-tracker"
	"gorm.io/gorm"
)

type categorySql struct {
	db *gorm.DB
}

func newCategorySQl(db *gorm.DB) Categories {
	return &categorySql{db: db}

}

func (db *categorySql) GetAllCategories() (categories []budget.Categories, err error) {
	err = db.db.Model(&budget.Categories{}).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil

}
