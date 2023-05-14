package app

import (
	"budget-tracker/internal/client"
	"budget-tracker/pkg/service"
	"github.com/BoryslavGlov/logrusx"
)

type App struct {
	s      *service.Service
	logg   logrusx.Logging
	client *client.Client
}

func NewApi(services *service.Service, logg logrusx.Logging, client *client.Client) *App {
	return &App{s: services, logg: logg, client: client}
}
