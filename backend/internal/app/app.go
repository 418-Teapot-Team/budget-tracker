package app

import (
	"budget-tracker/pkg/service"
	"github.com/BoryslavGlov/logrusx"
)

type App struct {
	s    *service.Service
	logg logrusx.Logging
}

func NewApi(services *service.Service, logg logrusx.Logging) *App {
	return &App{s: services, logg: logg}
}
