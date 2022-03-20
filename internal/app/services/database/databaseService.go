package database

import (
	"database/sql"
	"fmt"
	"github.com/PanovAlexey/url_carver/config"
	_ "github.com/jackc/pgx/stdlib"
	"os"
)

type databaseService struct {
	db sql.DB
}

func GetDatabaseService(config config.Config) *databaseService {
	databaseService := databaseService{}
	databaseService.db = *databaseService.initDatabaseConnection(config)

	return &databaseService
}

func (service databaseService) CheckDatabaseAvailability() error {
	return service.db.Ping()
}

func (service databaseService) GetDatabaseConnection() *sql.DB {
	return &service.db
}

func (service databaseService) initDatabaseConnection(config config.Config) *sql.DB {
	db, err := sql.Open("pgx", config.GetDatabaseDsn())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	return db
}
