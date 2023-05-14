package main

import (
	budget "budget-tracker"
	"budget-tracker/internal/app"
	"budget-tracker/internal/config"
	"budget-tracker/pkg/repository"
	"budget-tracker/pkg/service"
	"context"
	"github.com/subosito/gotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	if err := gotenv.Load(); err != nil {
		log.Fatal(err)
	}

	cfg := config.NewConfig()

	db, err := repository.NewDB(cfg.DatabaseUrl)
	if err != nil {
		log.Fatal("err while trying to initialize db", err)
	}

	repo := repository.NewRepository(db)
	authService := service.NewAuthService(repo)
	api := app.NewApi(authService)

	routers := app.Routers(api)
	srv := new(budget.Server)

	go func() {
		if err = srv.Run(cfg.Port, routers); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	log.Print("App Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("App Shutting Down")

	if err = srv.Shutdown(context.Background()); err != nil {
		log.Printf("error occured on server shutting down: %s", err.Error())
	}

}
