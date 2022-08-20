package services

import (
	"fmt"
	"github.com/PanovAlexey/url_carver/config"
	"github.com/PanovAlexey/url_carver/internal/app/domain/entity/url"
	"github.com/PanovAlexey/url_carver/internal/app/domain/entity/user"
	"github.com/PanovAlexey/url_carver/internal/app/services/database"
	"github.com/PanovAlexey/url_carver/internal/app/tests"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetDatabaseURLService(t *testing.T) {
	t.Run("Test getting database URL service", func(t *testing.T) {
		config := config.New()
		databaseService := database.GetDatabaseService(config)
		databaseUserRepository := GetDatabaseUserRepositoryMock(databaseService)
		databaseUserService := GetDatabaseUserService(databaseUserRepository)
		databaseURLService := GetDatabaseURLService(tests.GetDatabaseURLRepositoryMock(databaseService), *databaseUserService)
		structType := fmt.Sprintf("%T", databaseURLService)

		assert.Equal(t, structType, "*services.databaseURLService")
	})
}

func Test_SaveUser(t *testing.T) {
	t.Run("Test saving URL by database URL service", func(t *testing.T) {
		config := config.New()
		databaseService := database.GetDatabaseService(config)
		databaseUserRepository := GetDatabaseUserRepositoryMock(databaseService)
		databaseUserService := GetDatabaseUserService(databaseUserRepository)
		databaseURLService := GetDatabaseURLService(tests.GetDatabaseURLRepositoryMock(databaseService), *databaseUserService)
		url := url.URL{UserID: "1", LongURL: "www.google.ru", ShortURL: "123456"}
		ID, err := databaseURLService.SaveURL(url)

		assert.Equal(t, 777, ID)
		assert.Equal(t, nil, err)
	})
}

func Test_SaveBatchURLs(t *testing.T) {
	t.Run("Test saving batch URLs by database URL service", func(t *testing.T) {
		config := config.New()
		databaseService := database.GetDatabaseService(config)
		databaseUserRepository := GetDatabaseUserRepositoryMock(databaseService)
		databaseUserService := GetDatabaseUserService(databaseUserRepository)
		databaseURLService := GetDatabaseURLService(tests.GetDatabaseURLRepositoryMock(databaseService), *databaseUserService)
		urlExample := url.URL{UserID: "1", LongURL: "www.google.ru", ShortURL: "123456"}
		databaseURLService.SaveBatchURLs([]url.URL{urlExample})

	})
}

func Test_RemoveByShortURLSlice(t *testing.T) {
	t.Run("Test removing by short URL slice by database URL service", func(t *testing.T) {
		config := config.New()
		databaseService := database.GetDatabaseService(config)
		databaseUserRepository := GetDatabaseUserRepositoryMock(databaseService)
		databaseUserService := GetDatabaseUserService(databaseUserRepository)
		databaseURLService := GetDatabaseURLService(tests.GetDatabaseURLRepositoryMock(databaseService), *databaseUserService)
		var URLsCollection []string

		err := databaseURLService.RemoveByShortURLSlice(URLsCollection, "test_user_token")

		assert.Equal(t, nil, err)
	})
}

func Test_GetURLCollectionFromStorage(t *testing.T) {
	t.Run("Test getting URL collection from storage by database URL service", func(t *testing.T) {
		config := config.New()
		databaseService := database.GetDatabaseService(config)
		databaseUserRepository := GetDatabaseUserRepositoryMock(databaseService)
		databaseUserService := GetDatabaseUserService(databaseUserRepository)
		databaseURLService := GetDatabaseURLService(tests.GetDatabaseURLRepositoryMock(databaseService), *databaseUserService)
		URLsCollection := databaseURLService.GetURLCollectionFromStorage()

		assert.Equal(t, 0, len(URLsCollection))
	})
}

type databaseUserRepositoryMock struct {
	databaseService database.DatabaseInterface
}

func GetDatabaseUserRepositoryMock(databaseService database.DatabaseInterface) databaseUserRepositoryInterface {
	return databaseUserRepositoryMock{databaseService: databaseService}
}

func (repository databaseUserRepositoryMock) SaveUser(user user.User) (int, error) {
	return 1, nil
}

func (repository databaseUserRepositoryMock) GetUserByID(userID int) (user.User, error) {
	return user.User{}, nil
}

func (repository databaseUserRepositoryMock) GetUserByGUID(guid string) (user.User, error) {
	return user.User{}, nil
}

func (repository databaseUserRepositoryMock) IsExistUserByGUID(guid string) bool {
	return true
}
