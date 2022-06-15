package tests

import (
	"github.com/PanovAlexey/url_carver/internal/app/domain/dto"
	"github.com/PanovAlexey/url_carver/internal/app/services/database"
)

type databaseURLRepositoryMock struct {
	databaseService database.DatabaseInterface
}

func GetDatabaseURLRepositoryMock(databaseService database.DatabaseInterface) databaseURLRepositoryMock {
	return databaseURLRepositoryMock{databaseService: databaseService}
}

func (repository databaseURLRepositoryMock) SaveURL(url dto.DatabaseURL) (int, error) {
	return 777, nil
}

func (repository databaseURLRepositoryMock) GetList() ([]dto.DatabaseURL, error) {
	return []dto.DatabaseURL{}, nil
}

func (repository databaseURLRepositoryMock) SaveBatchURLs(collection []dto.DatabaseURL) error {
	return nil
}

func (repository databaseURLRepositoryMock) DeleteURLsByShortValueSlice([]string, int) ([]dto.DatabaseURL, error) {
	return []dto.DatabaseURL{}, nil
}
