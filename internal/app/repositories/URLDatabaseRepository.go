package repositories

import (
	"github.com/PanovAlexey/url_carver/internal/app/domain/dto"
	"github.com/PanovAlexey/url_carver/internal/app/services/database"
	"log"
	"strconv"
)

type databaseURLRepository struct {
	databaseService database.DatabaseInterface
}

func GetDatabaseURLRepository(databaseService database.DatabaseInterface) *databaseURLRepository {
	return &databaseURLRepository{databaseService: databaseService}
}

func (repository databaseURLRepository) SaveURL(url dto.DatabaseURL) (int, error) {
	var insertedID int

	query := "INSERT INTO " + database.TableURLsName + " (user_id, url, short_url) VALUES ($1, $2, $3) RETURNING id"
	err := repository.databaseService.GetDatabaseConnection().
		QueryRow(query, url.GetUserID(), url.GetLongURL(), url.GetShortURL()).Scan(&insertedID)

	if err != nil {
		return 0, err // ToDo: 0 - is a crutch
	}

	return insertedID, err
}

func (repository databaseURLRepository) GetList() ([]dto.DatabaseURL, error) {
	var collection []dto.DatabaseURL

	var resultID, resultUserID int
	var resultURL, resultShortURL string

	query := "SELECT id, user_id, url, short_url FROM " + database.TableURLsName
	rows, err := repository.databaseService.GetDatabaseConnection().Query(query)

	if err != nil {
		return collection, err
	}

	for rows.Next() {
		err = rows.Scan(&resultID, &resultUserID, &resultURL, &resultShortURL)

		if err != nil {
			return collection, err
		}

		collection = append(collection, dto.NewDatabaseURL(resultURL, resultShortURL, resultUserID))
	}

	return collection, nil
}

func (repository databaseURLRepository) SaveBatchURLs(collection []dto.DatabaseURL) error {
	dbConnection := repository.databaseService.GetDatabaseConnection()
	tx, err := dbConnection.Begin()

	if err != nil {
		return err
	}

	defer tx.Rollback()
	statement, err := tx.Prepare(
		"INSERT INTO " + database.TableURLsName + "(user_id, url, short_url) VALUES($1,$2,$3) RETURNING id",
	)

	if err != nil {
		return err
	}

	defer statement.Close()

	var insertedID int
	var resultString string

	for _, url := range collection {
		err = statement.QueryRow(url.GetUserID(), url.GetLongURL(), url.GetShortURL()).Scan(&insertedID)

		if len(resultString) > 0 {
			resultString = resultString + ", "
		}

		resultString = resultString + strconv.Itoa(insertedID)

		if err != nil {
			return err
		}
	}

	log.Println("created URLs with id: ", resultString)

	return tx.Commit()
}
