package main

import (
	"github.com/spf13/viper"
	"log"
	"github.com/fuadsuleyman/go-auth1"
	"github.com/fuadsuleyman/go-auth1/pkg/handler"
	"github.com/fuadsuleyman/go-auth1/pkg/repository"
	"github.com/fuadsuleyman/go-auth1/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(auth.Server)
	// asagida 8000 evezine deyishen yazilmalidi
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil { 
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
} 