package service

import "budget-tracker/pkg/repository"

type Service struct {
	AuthService
	ListsService
	CategoriesService
	AccountsService
}

func NewService(reps *repository.Repository) *Service {
	return &Service{
		AuthService{repo: reps.Authorization},
		ListsService{repo: reps.Lists},
		CategoriesService{repo: reps.Categories},
		AccountsService{repo: reps.Accounts},
	}
}
