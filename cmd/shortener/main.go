package main

import (
	"github.com/PanovAlexey/url_carver/internal/app/handlers/http"
	"github.com/PanovAlexey/url_carver/internal/app/repositories"
	"github.com/PanovAlexey/url_carver/internal/app/servers"
	"github.com/PanovAlexey/url_carver/internal/app/services"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	initialize()

	emailRepository := repositories.GetEmailRepository()
	shortURLService := services.GetShortURLService(emailRepository)
	httpHandler := http.GetHTTPHandler(shortURLService)
	servers.RunServer(httpHandler)
}

func initialize() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}
}
