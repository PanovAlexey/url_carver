package repositories

import (
	"errors"
	"github.com/PanovAlexey/url_carver/internal/app/domain/dto"
	"github.com/PanovAlexey/url_carver/internal/app/services"
	"github.com/PanovAlexey/url_carver/internal/app/services/database"
	"github.com/lib/pq"
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
		QueryRow(query, url.UserID, url.LongURL, url.ShortURL).Scan(&insertedID)

	if err != nil {
		errorService := services.GetErrorService()
		err = errorService.GetActualizedError(err, url)
	}

	return insertedID, err
}

func (repository databaseURLRepository) GetList() ([]dto.DatabaseURL, error) {
	var collection []dto.DatabaseURL

	var resultID, resultUserID int
	var resultURL, resultShortURL string
	var isDeleted bool

	query := "SELECT id, user_id, url, short_url, is_deleted FROM " + database.TableURLsName
	rows, err := repository.databaseService.GetDatabaseConnection().Query(query)

	if err != nil {
		return collection, err
	}

	for rows.Next() {
		err = rows.Scan(&resultID, &resultUserID, &resultURL, &resultShortURL, &isDeleted)

		if err != nil {
			return collection, err
		}

		databaseURL := dto.DatabaseURL{
			LongURL:   resultURL,
			ShortURL:  resultShortURL,
			UserID:    resultUserID,
			IsDeleted: isDeleted,
		}

		collection = append(collection, databaseURL)
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
		err = statement.QueryRow(url.UserID, url.LongURL, url.ShortURL).Scan(&insertedID)

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

func (repository databaseURLRepository) DeleteURLsByShortValueSlice(
	shortURLValuesSlice []string, userID int) ([]dto.DatabaseURL, error) {
	query := "UPDATE " + database.TableURLsName +
		" SET is_deleted = true " +
		"WHERE short_url = any($1) AND user_id=" + strconv.Itoa(userID) +
		"RETURNING id, user_id, url, short_url, is_deleted"
	rows, err := repository.databaseService.GetDatabaseConnection().Query(query, pq.Array(shortURLValuesSlice))

	if err != nil {
		return nil, err
	}

	var resultID int
	var resultUserID int
	var resultURL string
	var resultShortURL string
	var resultIsDeleted bool
	var result = make([]dto.DatabaseURL, 0)

	var errorsText string

	for rows.Next() {
		err = rows.Scan(&resultID, &resultUserID, &resultURL, &resultShortURL, &resultIsDeleted)

		if err != nil {
			errorsText = errorsText + " " + err.Error()
		}

		databaseURL := dto.DatabaseURL{
			LongURL:   resultURL,
			ShortURL:  resultShortURL,
			UserID:    resultUserID,
			IsDeleted: true,
		}

		result = append(result, databaseURL)
	}

	if len(errorsText) == 0 {
		return result, nil
	}

	return result, errors.New(errorsText)
}
