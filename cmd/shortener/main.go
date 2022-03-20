package main

import (
	"database/sql"
	"fmt"
	"github.com/PanovAlexey/url_carver/config"
	"github.com/PanovAlexey/url_carver/internal/app/handlers/http"
	"github.com/PanovAlexey/url_carver/internal/app/repositories"
	"github.com/PanovAlexey/url_carver/internal/app/servers"
	"github.com/PanovAlexey/url_carver/internal/app/services"
	"github.com/PanovAlexey/url_carver/internal/app/services/encryption"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	config := config.New()
	httpHandler := getHttpHandler(config)

	db := getDatabaseConnection(config)
	defer db.Close()
	checkDatabaseAvailability(db)

	servers.RunServer(httpHandler, config)
}

func checkDatabaseAvailability(db *sql.DB) {
	err := db.Ping()

	if err != nil {
		log.Println("Database connection error", err.Error())
	} else {
		log.Println("Database connection successfully")
	}
}

func getDatabaseConnection(config config.Config) *sql.DB {
	db, err := sql.Open("pgx", config.GetDatabaseDsn())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return db
}

func getHttpHandler(config config.Config) servers.HandlerInterface {
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
	)

	return httpHandler
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("error loading env variables: %s", err.Error())
	}
}
