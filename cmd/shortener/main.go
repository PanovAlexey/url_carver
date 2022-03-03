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
	config := config.New()
	URLMemoryRepository := repositories.GetURLMemoryRepository()
	fileStorageRepository, err := repositories.GetFileStorageRepository(config)

	if err != nil {
		log.Fatalf("error creating file repository by config:" + err.Error())
	} else {
		defer fileStorageRepository.Close()
	}

	URLShorteningService := services.GetURLShorteningService(config)
	URLStorageService := services.GetURLStorageService(config, fileStorageRepository)
	URLMemoryService := services.GetURLMemoryService(config, URLMemoryRepository, URLShorteningService)
	URLMemoryService.LoadURLs(URLStorageService.GetURLCollectionFromStorage())
	httpHandler := http.GetHTTPHandler(URLMemoryService, URLStorageService)

	servers.RunServer(httpHandler, config)
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("error loading env variables: %s", err.Error())
	}
}
