package database

import "database/sql"

type DatabaseInterface interface {
	GetDatabaseConnection() *sql.DB
	CheckDatabaseAvailability() error
	MigrateUp()
	MigrateDown()
}
