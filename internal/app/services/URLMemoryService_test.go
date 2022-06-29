package services

import (
	"github.com/PanovAlexey/url_carver/config"
	"github.com/PanovAlexey/url_carver/internal/app/domain/entity/url"
	"github.com/PanovAlexey/url_carver/internal/app/repositories"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetURLEntityByShortURL(t *testing.T) {
	memoryService := getMemoryService()
	URLEntity, err := memoryService.GetURLEntityByShortURL("http://www.google.ru")

	t.Run("Negative test GetURLEntityByShortURL by missing URL", func(t *testing.T) {
		assert.Equal(t, nil, err)
		assert.Equal(t, "", URLEntity)
	})
}

func Test_IsExistURLEntityByShortURL(t *testing.T) {
	memoryService := getMemoryService()
	isExist := memoryService.IsExistURLEntityByShortURL("http://www.google.ru")

	t.Run("Negative test IsExistURLEntityByShortURL by missing URL", func(t *testing.T) {
		assert.Equal(t, false, isExist)
	})
}

func Test_GetShortURLDtoByURL(t *testing.T) {
	config := config.New()
	memoryService := getMemoryService()

	urlExample := url.URL{UserID: "1", LongURL: "www.google.ru", ShortURL: "123456"}
	dtoShortURL := memoryService.GetShortURLDtoByURL(urlExample)

	t.Run("Positive test GetShortURLDtoByURL by correct URL", func(t *testing.T) {
		assert.Equal(t, config.GetBaseURL()+"/"+urlExample.ShortURL, dtoShortURL.Value)
	})
}

func Test_LoadURLs(t *testing.T) {
	config := config.New()
	memoryService := getMemoryService()
	fileStorageRepository, err := repositories.GetFileStorageRepository(config)
	storageService := GetStorageService(config, fileStorageRepository)
	memoryService.LoadURLs(storageService.GetURLCollectionFromStorage())

	collection := []url.URL{}
	URLExample := url.URL{LongURL: "http://codeblog.pro", ShortURL: "896ABC", UserID: "1", IsDeleted: false}
	collection = append(collection, URLExample)
	memoryService.LoadURLs(collection)

	dtoShortURL := memoryService.GetShortURLDtoByURL(URLExample)

	t.Run("Positive test LoadURLs by correct URL", func(t *testing.T) {
		assert.Equal(t, nil, err)
		assert.Equal(t, dtoShortURL.Value, config.GetBaseURL()+"/"+URLExample.ShortURL)
	})
}

func Test_SaveURL(t *testing.T) {
	memoryService := getMemoryService()

	URLExample := url.URL{LongURL: "http://codeblog2.pro", ShortURL: "896ABCAZ", UserID: "1", IsDeleted: false}
	isSaved := memoryService.SaveURL(URLExample)

	var URLsCollection []string
	URLsCollection = append(URLsCollection, "896ABCAZ")

	memoryService.DeleteURLsByShortValueSlice(URLsCollection)

	t.Run("Positive test SaveURL", func(t *testing.T) {
		assert.Equal(t, true, isSaved)
	})
}

func Test_GetURLsByUserToken(t *testing.T) {
	memoryService := getMemoryService()
	URLsCollection := memoryService.GetURLsByUserToken("WRONG_USER_TOKEN")

	t.Run("Negative test GetURLsByUserToken by wrong user key", func(t *testing.T) {
		assert.Equal(t, 0, len(URLsCollection))
	})
}

func getMemoryService() *MemoryService {
	config := config.New()
	URLMemoryRepository := repositories.GetURLMemoryRepository()
	shorteningService := GetShorteningService(config)

	return &MemoryService{Config: config, Repository: URLMemoryRepository, ShorteningService: *shorteningService}
}
