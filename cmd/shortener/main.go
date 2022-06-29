package main

import (
	"fmt"
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

var (
	buildVersion string
	buildDate    string
	buildCommit  string
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("error loading env variables: %s", err.Error())
	}
}

func main() {
	displayVersionInfo()

	config := config.New()
	databaseService, err := getDatabaseService(config)

	if err == nil {
		databaseService.MigrateUp()
		defer databaseService.GetDatabaseConnection().Close()
	}

	fileStorageRepository, err := repositories.GetFileStorageRepository(config)

	if err != nil {
		log.Fatalln("error creating file repository by config:" + err.Error())
	} else {
		defer fileStorageRepository.Close()
	}

	globalURLDeletingChannel := getGlobalURLDeletingChannel()
	defer close(globalURLDeletingChannel)

	httpHandler := getHTTPHandler(config, databaseService, fileStorageRepository)

	servers.RunServer(httpHandler, config)
}

func getGlobalURLDeletingChannel() chan string {
	queueMap := services.GetChannelsMapService()
	return queueMap.GetChannelByName(services.ChannelWithRemovingURLsName)
}

func getHTTPHandler(
	config config.Config,
	databaseService database.DatabaseInterface,
	fileStorageRepository repositories.FileStorageRepositoryInterface,
) servers.HandlerInterface {
	URLMemoryRepository := repositories.GetURLMemoryRepository()
	databaseURLRepository := repositories.GetDatabaseURLRepository(databaseService)
	databaseUserRepository := repositories.GetDatabaseUserRepository(databaseService)

	databaseUserService := services.GetDatabaseUserService(databaseUserRepository)
	databaseURLService := services.GetDatabaseURLService(databaseURLRepository, *databaseUserService)
	shorteningService := services.GetShorteningService(config)
	storageService := services.GetStorageService(config, fileStorageRepository)
	memoryService := services.GetMemoryService(config, URLMemoryRepository, shorteningService)
	contextStorageService := services.GetContextStorageService()
	userTokenAuthorizationService := services.GetUserTokenAuthorizationService()
	URLMappingService := services.GetURLMappingService()
	encryptionService, err := encryption.NewEncryptionService(config)

	if err != nil {
		log.Println("error with encryption service initialization: " + err.Error())
	}

	// Load URLs to memory from other storages
	memoryService.LoadURLs(storageService.GetURLCollectionFromStorage())
	memoryService.LoadURLs(databaseURLService.GetURLCollectionFromStorage())

	httpHandler := http.GetHTTPHandler(
		memoryService,
		storageService,
		encryptionService,
		shorteningService,
		contextStorageService,
		userTokenAuthorizationService,
		URLMappingService,
		databaseService,
		databaseURLService,
		databaseUserService,
	)

	return httpHandler
}

func getDatabaseService(config config.Config) (database.DatabaseInterface, error) {
	databaseService := database.GetDatabaseService(config)
	err := databaseService.CheckDatabaseAvailability()

	if err != nil {
		log.Println("Database connection error", err.Error())
	} else {
		log.Println("Database connection successfully")
	}

	return databaseService, err
}

func displayVersionInfo() {
	const defaultInfo = "N/A"

	if buildVersion == "" {
		buildVersion = defaultInfo
	}

	if buildDate == "" {
		buildDate = defaultInfo
	}

	if buildCommit == "" {
		buildCommit = defaultInfo
	}

	fmt.Printf("Build version: %s\n", buildVersion)
	fmt.Printf("Build date: %s\n", buildDate)
	fmt.Printf("Build commit: %s\n", buildCommit)
	fmt.Println()
}
