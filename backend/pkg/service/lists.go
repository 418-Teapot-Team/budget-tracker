package service

import (
	budget "budget-tracker"
	"budget-tracker/pkg/repository"
	"database/sql"
	"time"
)

type ListsService struct {
	repo repository.Lists
}

func (l *ListsService) CreateList(input *budget.List, userId int) (err error) {
	input.CreatedAt = sql.NullTime{Time: time.Now(), Valid: true}
	input.UserId = userId
	return l.repo.CreateList(input)
}

func (l *ListsService) DeleteList(listId, userId int) (err error) {
	return l.repo.DeleteList(listId, userId)
}

func (l *ListsService) GetList(userId int, budgetType, orderBy, sortedBy string, takeAmount int, skipAmount int, categoryId int) (lists []budget.ListsGetter, err error) {
	return l.repo.GetList(userId, budgetType, orderBy, sortedBy, takeAmount, skipAmount, categoryId)
}

func (l *ListsService) GetTopListsCategories(userId int, listType string, takeAmount int, months int) (lists []budget.ListsGetter, err error) {
	return l.repo.GetTopCategories(userId, listType, takeAmount, months)
}

func (l *ListsService) EditList(input budget.List) error {
	return l.repo.EditList(input)
}

func (l *ListsService) GetCurrentMonthSavings(userId int) (incomeNet float64, err error) {
	return l.repo.GetCurrentMonthSavings(userId)
}

func (l *ListsService) GetSavingsStats(userId int) (data []budget.FinancialData, err error) {
	return l.repo.GetSavingsStats(userId)
}

func (l *ListsService) GetTotalAmount(userId int, budgetType string, months int) (data float64, err error) {
	return l.repo.GetTotalAmount(userId, budgetType, months)
}

func (l *ListsService) GetStats(userId int, budgetType string, months int) (data []budget.FinancialData, err error) {
	return l.repo.GetStats(userId, budgetType, months)
}
