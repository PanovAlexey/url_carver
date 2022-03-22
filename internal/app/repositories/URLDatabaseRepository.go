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
	var insertedId int

	query := "INSERT INTO urls (user_id, url, short_url) VALUES ($1, $2, $3) RETURNING id"
	err := repository.databaseService.GetDatabaseConnection().
		QueryRow(query, url.GetUserID(), url.GetLongURL(), url.GetShortURL()).Scan(&insertedId)

	if err != nil {
		return 0, err // ToDo: 0 - is a crutch
	}

	return insertedId, err
}

func (repository databaseURLRepository) GetURLByKey(key string) url.URL {
	return url.URL{} // todo
}

func (repository databaseURLRepository) IsExistURLByKey(key string) bool {
	return true // todo
}
