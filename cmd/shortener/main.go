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
	URLMemoryRepository := repositories.GetURLMemoryRepository()
	fileStorageRepository, error := repositories.GetFileStorageRepository(config)

	if error != nil {
		log.Fatalln("error creating file repository by config:" + error.Error())
	}

	URLShorteningService := services.GetURLShorteningService(config)
	URLMemoryService := services.GetURLMemoryService(config, URLMemoryRepository, URLShorteningService)
	URLStorageService := services.GetURLStorageService(config, fileStorageRepository)
	httpHandler := http.GetHTTPHandler(URLMemoryService, URLStorageService)
	servers.RunServer(httpHandler, config)
}

func initialize() {
	if err := godotenv.Load(); err != nil {
		log.Printf("error loading env variables: %s", err.Error())
	}
}
