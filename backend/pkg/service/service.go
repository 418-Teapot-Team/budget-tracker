package service

import "budget-tracker/pkg/repository"

type Service struct {
	AuthService
}

func NewService(reps *repository.Repository) *Service {
	return &Service{AuthService{repo: reps.Authorization}}
}
