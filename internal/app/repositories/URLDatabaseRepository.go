package repositories

import (
	"database/sql"
	"github.com/PanovAlexey/url_carver/internal/app/domain/dto"
	"github.com/lib/pq"
	"log"
	"strconv"
)

type errorServiceInterface interface {
	GetActualizedError(err error, additionalInfo interface{}) error
}

type DatabaseURLRepository struct {
	DB           *sql.DB
	ErrorService errorServiceInterface
}

func (repository DatabaseURLRepository) SaveURL(url dto.DatabaseURL) (int, error) {
	var insertedID int

	query := "INSERT INTO " + `urls` + " (user_id, url, short_url) VALUES ($1, $2, $3) RETURNING id"
	err := repository.DB.QueryRow(query, url.UserID, url.LongURL, url.ShortURL).Scan(&insertedID)

	if err != nil {
		err = repository.ErrorService.GetActualizedError(err, url)
		log.Println(err)
	}

	return insertedID, err
}

func (repository DatabaseURLRepository) GetList() ([]dto.DatabaseURL, error) {
	var collection []dto.DatabaseURL

	var resultID, resultUserID int
	var resultURL, resultShortURL string
	var isDeleted bool

	query := "SELECT id, user_id, url, short_url, is_deleted FROM " + `urls`
	rows, err := repository.DB.Query(query)

	if err != nil || rows.Err() != nil {
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

func (repository DatabaseURLRepository) SaveBatchURLs(collection []dto.DatabaseURL) error {
	tx, err := repository.DB.Begin()

	if err != nil {
		return err
	}

	defer tx.Rollback()
	statement, err := tx.Prepare(
		"INSERT INTO " + `urls` + "(user_id, url, short_url) VALUES($1,$2,$3) RETURNING id",
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

func (repository DatabaseURLRepository) DeleteURLsByShortValueSlice(
	shortURLValuesSlice []string, userID int) ([]dto.DatabaseURL, error,
) {
	query := "UPDATE " + `urls` + " SET is_deleted = true " +
		"WHERE short_url = any($1) AND user_id=($2) RETURNING id, user_id, url, short_url, is_deleted"

	rows, err := repository.DB.Query(
		query,
		pq.Array(shortURLValuesSlice),
		strconv.Itoa(userID),
	)

	if err != nil || rows.Err() != nil {
		return nil, err
	}

	var result = make([]dto.DatabaseURL, 0)

	for rows.Next() {
		var u dto.DatabaseURL

		if err := rows.Scan(&u.ID, &u.UserID, &u.LongURL, &u.ShortURL, &u.IsDeleted); err != nil {
			if err == sql.ErrNoRows {
				return result, nil
			}

			return nil, err
		}

		result = append(result, u)
	}

	return result, nil
}
