package main

import (
	"github.com/PanovAlexey/url_carver/config"
	"github.com/PanovAlexey/url_carver/internal/app/handlers/http"
	"github.com/PanovAlexey/url_carver/internal/app/repositories"
	"github.com/PanovAlexey/url_carver/internal/app/servers"
	"github.com/PanovAlexey/url_carver/internal/app/services"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	initialize()

	config := config.New()
	URLRepository := repositories.GetURLRepository()
	shortURLService := services.GetShortURLService(URLRepository, config)
	URLStorageService := services.GetURLStorageService(config)
	httpHandler := http.GetHTTPHandler(shortURLService, URLStorageService)
	servers.RunServer(httpHandler, config)
}

func initialize() {
	if err := godotenv.Load(); err != nil {
		log.Printf("error loading env variables: %s", err.Error())
	}
}
