package database

import (
	"database/sql"
	"github.com/PanovAlexey/url_carver/config"
	_ "github.com/jackc/pgx/stdlib"
	"io/ioutil"
	"log"
	"os"
)

type databaseService struct {
	db *sql.DB
}

const TableURLsName = `urls`
const TableUsersName = `users`
const migrationUpFilesPath = `internal/app/services/database/migrations/up`
const migrationDownFilesPath = `internal/app/services/database/migrations/down`
const migrationFilePermissions = 0777

func GetDatabaseService(config config.Config) *databaseService {
	databaseService := databaseService{}
	databaseService.db = databaseService.initDatabaseConnection(config)

	return &databaseService
}

func (service databaseService) CheckDatabaseAvailability() error {
	return service.db.Ping()
}

func (service databaseService) GetDatabaseConnection() *sql.DB {
	return service.db
}

func (service databaseService) MigrateUp() {
	filesPath := service.getMigrationFilesPath(migrationUpFilesPath)

	for _, filePath := range filesPath {
		migrationFileContent, err := service.getMigrationFileContent(filePath)

		if err != nil {
			log.Println(`impossible to get content of migration file by path: ` + filePath)
		}

		isApplied := service.applyMigrate(migrationFileContent)

		if isApplied {
			log.Println(`migration applied successfully (` + filePath + `)`)
		}
	}
}

func (service databaseService) MigrateDown() {
	filesPath := service.getMigrationFilesPath(migrationDownFilesPath)

	for _, filePath := range filesPath {
		migrationFileContent, err := service.getMigrationFileContent(filePath)

		if err != nil {
			log.Println(`impossible to get content of migration file by path: ` + filePath)
		}

		service.applyMigrate(migrationFileContent)
	}
}

func (service databaseService) initDatabaseConnection(config config.Config) *sql.DB {
	db, err := sql.Open("pgx", config.GetDatabaseDsn())

	if err != nil {
		log.Printf(`Unable to connect to database: %v\n`, err)
	}

	return db
}

func (service databaseService) getMigrationFileContent(filePath string) (string, error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, migrationFilePermissions)

	if err != nil {
		log.Println(`error with opening migration file ` + filePath + ` ` + err.Error())
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Println(`error with closing migrate file: ` + filePath + ` ` + err.Error())
		}
	}()

	contentString := ``
	content, err := ioutil.ReadAll(file)

	if err == nil {
		contentString = string(content)
	}

	return contentString, err
}

func (service databaseService) applyMigrate(sqlQuery string) bool {
	db := service.GetDatabaseConnection()
	_, err := db.Exec(sqlQuery)

	if err != nil {
		log.Println(`error while applying migration. ` + err.Error())
		return false
	}

	return true
}

func (service databaseService) getMigrationFileFullName(path string, fileName string) string {
	return path + `/` + fileName
}

func (service databaseService) getMigrationFilesPath(dirPath string) []string {
	files, err := ioutil.ReadDir(dirPath)
	filesPath := make([]string, len(files))

	if err != nil {
		log.Println("error while getting migration files by dir: " + dirPath + ". " + err.Error())
	}

	for key, file := range files {
		if !file.IsDir() {
			filesPath[key] = service.getMigrationFileFullName(dirPath, file.Name())
		}
	}

	return filesPath
}
