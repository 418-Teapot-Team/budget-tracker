package app

import (
	"budget-tracker/pkg/service"
)

type App struct {
	s *service.AuthService
}

func NewApi(s *service.AuthService) *App {
	return &App{s: s}
}
