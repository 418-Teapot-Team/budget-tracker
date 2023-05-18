package service

import (
	budget "budget-tracker"
	"budget-tracker/pkg/repository"
	"time"
)

type ListsService struct {
	repo repository.Lists
}

func (l *ListsService) CreateList(input *budget.List, userId int) (err error) {
	*input.CreatedAt = time.Now()
	input.UserId = userId
	return l.repo.CreateList(input)
}

func (l *ListsService) DeleteList(listId, userId int) (err error) {
	return l.repo.DeleteList(listId, userId)
}

func (l *ListsService) GetList(userId int, budgetType, orderBy, sortedBy string) (lists []budget.ListsGetter, err error) {
	return l.repo.GetList(userId, budgetType, orderBy, sortedBy)
}

func (l *ListsService) GetTopExpenses(userId int) (lists []budget.ListsGetter, err error) {
	return l.repo.GetTopExpenses(userId)
}

func (l *ListsService) EditList(input budget.List) error {
	return l.repo.EditList(input)
}

func (l *ListsService) GetCurrentMonthSavings(userId int) (incomeNet int64, err error) {
	return l.repo.GetCurrentMonthSavings(userId)
}
