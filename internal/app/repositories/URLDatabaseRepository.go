package repositories

import (
	"github.com/PanovAlexey/url_carver/internal/app/domain/dto"
	"github.com/PanovAlexey/url_carver/internal/app/domain/entity/url"
	"github.com/PanovAlexey/url_carver/internal/app/services/database"
)

type databaseURLRepository struct {
	databaseService database.DatabaseInterface
}

func GetDatabaseURLRepository(databaseService database.DatabaseInterface) *databaseURLRepository {
	return &databaseURLRepository{databaseService: databaseService}
}

func (repository databaseURLRepository) SaveURL(url dto.DatabaseURLInterface) (int, error) {
	var insertedID int

	query := "INSERT INTO " + database.TableURLsName + " (user_id, url, short_url) VALUES ($1, $2, $3) RETURNING id"
	err := repository.databaseService.GetDatabaseConnection().
		QueryRow(query, url.GetUserID(), url.GetLongURL(), url.GetShortURL()).Scan(&insertedID)

	if err != nil {
		return 0, err // ToDo: 0 - is a crutch
	}

	return insertedID, err
}

func (repository databaseURLRepository) GetURLByKey(key string) url.URL {
	return url.URL{} // todo
}

func (repository databaseURLRepository) IsExistURLByKey(key string) bool {
	return true // todo
}

func (repository databaseURLRepository) GetList() (dto.URLDatabaseCollection, error) {
	collection := dto.GetURLDatabaseCollection()

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

		collection.AppendURL(dto.NewDatabaseURL(resultURL, resultShortURL, resultUserID))
	}

	return collection, nil
}
