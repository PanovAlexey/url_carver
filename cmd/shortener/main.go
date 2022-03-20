package main

import (
	"github.com/PanovAlexey/url_carver/config"
	"github.com/PanovAlexey/url_carver/internal/app/handlers/http"
	"github.com/PanovAlexey/url_carver/internal/app/repositories"
	"github.com/PanovAlexey/url_carver/internal/app/servers"
	"github.com/PanovAlexey/url_carver/internal/app/services"
	"github.com/PanovAlexey/url_carver/internal/app/services/database"
	"github.com/PanovAlexey/url_carver/internal/app/services/encryption"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	config := config.New()
	httpHandler := getHTTPHandler(config)

	servers.RunServer(httpHandler, config)
}

func getHTTPHandler(config config.Config) servers.HandlerInterface {
	URLMemoryRepository := repositories.GetURLMemoryRepository()
	fileStorageRepository, err := repositories.GetFileStorageRepository(config)

	if err != nil {
		log.Fatalln("error creating file repository by config:" + err.Error())
	} else {
		defer fileStorageRepository.Close()
	}

	databaseService := getDatabaseService(config)
	shorteningService := services.GetShorteningService(config)
	storageService := services.GetStorageService(config, fileStorageRepository)
	memoryService := services.GetMemoryService(config, URLMemoryRepository, shorteningService)
	memoryService.LoadURLs(storageService.GetURLCollectionFromStorage())
	contextStorageService := services.GetContextStorageService()
	userTokenAuthorizationService := services.GetUserTokenAuthorizationService()
	URLMappingService := services.GetURLMappingService()
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
		URLMappingService,
		databaseService,
	)

	return httpHandler
}

func getDatabaseService(config config.Config) database.DatabaseInterface {
	databaseService := database.GetDatabaseService(config)
	db := databaseService.GetDatabaseConnection()
	defer db.Close()

	err := databaseService.CheckDatabaseAvailability()

	if err != nil {
		log.Println("Database connection error", err.Error())
	} else {
		log.Println("Database connection successfully")
	}

	return databaseService
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("error loading env variables: %s", err.Error())
	}
}
