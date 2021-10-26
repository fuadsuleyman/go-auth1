package main

import (
	"log"

	"github.com/fuadsuleyman/go-auth1"
	"github.com/fuadsuleyman/go-auth1/pkg/handler"
	"github.com/fuadsuleyman/go-auth1/pkg/repository"
	"github.com/fuadsuleyman/go-auth1/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(auth.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}