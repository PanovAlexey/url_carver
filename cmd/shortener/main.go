package main

import (
	"github.com/PanovAlexey/url_carver/config"
	"github.com/PanovAlexey/url_carver/internal/app/handlers/http"
	"github.com/PanovAlexey/url_carver/internal/app/repositories"
	"github.com/PanovAlexey/url_carver/internal/app/servers"
	"github.com/PanovAlexey/url_carver/internal/app/services"
	"github.com/PanovAlexey/url_carver/internal/app/services/encryption"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	config := config.New()
	URLMemoryRepository := repositories.GetURLMemoryRepository()
	fileStorageRepository, err := repositories.GetFileStorageRepository(config)

	if err != nil {
		log.Fatalln("error creating file repository by config:" + err.Error())
	} else {
		defer fileStorageRepository.Close()
	}

	shorteningService := services.GetShorteningService(config)
	storageService := services.GetStorageService(config, fileStorageRepository)
	memoryService := services.GetMemoryService(config, URLMemoryRepository, shorteningService)
	memoryService.LoadURLs(storageService.GetURLCollectionFromStorage())
	contextStorageService := services.GetContextStorageService()
	userTokenAuthorizationService := services.GetUserTokenAuthorizationService()
	encryptionService, err := encryption.NewEncryptionService(config)

	if err != nil {
		log.Println("error with encryption service initialization: " + err.Error())
	}

	httpHandler := http.GetHTTPHandler(
		memoryService,
		storageService,
		encryptionService,
		shorteningService,
		contextStorageService,
		userTokenAuthorizationService,
	)

	servers.RunServer(httpHandler, config)
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("error loading env variables: %s", err.Error())
	}
}
