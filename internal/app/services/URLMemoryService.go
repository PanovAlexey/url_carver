package services

import (
	"fmt"
	"github.com/PanovAlexey/url_carver/config"
	"github.com/PanovAlexey/url_carver/internal/app/domain/dto"
	urlEntity "github.com/PanovAlexey/url_carver/internal/app/domain/entity/url"
	databaseErrors "github.com/PanovAlexey/url_carver/internal/app/services/database/errors"
)

type repositoryInterface interface {
	AddURL(url urlEntity.URL) bool
	GetURLByKey(key string) urlEntity.URL
	IsExistURLByKey(key string) bool
	GetURLsByUserToken(userToken string) []urlEntity.URL
	GetURLsByShortValueSlice(urlShortValuesSlice []string) []urlEntity.URL
}

type MemoryService struct {
	Repository        repositoryInterface
	Config            config.Config
	ShorteningService ShorteningService
}

func (service MemoryService) GetURLEntityByShortURL(shortURL string) (string, error) {
	url := service.Repository.GetURLByKey(shortURL)

	if url.IsDeleted {
		return "", fmt.Errorf("%v: %w", url, databaseErrors.ErrorIsDeleted)
	}

	return url.LongURL, nil
}

func (service MemoryService) IsExistURLEntityByShortURL(shortURL string) bool {
	return service.Repository.IsExistURLByKey(shortURL)
}

func (service MemoryService) GetShortURLDtoByURL(url urlEntity.URL) dto.ShortURL {
	shortURLWithDomain := service.ShorteningService.GetShortURLWithDomain(url.ShortURL)

	return dto.GetShortURLByValue(shortURLWithDomain)
}

func (service MemoryService) LoadURLs(collection []urlEntity.URL) {
	for _, url := range collection {
		service.SaveURL(url)
	}
}

func (service MemoryService) SaveURL(url urlEntity.URL) bool {
	return service.Repository.AddURL(url)
}

func (service MemoryService) DeleteURLsByShortValueSlice(urlShortValuesSlice []string) {
	urlCollection := service.Repository.GetURLsByShortValueSlice(urlShortValuesSlice)

	for _, url := range urlCollection {
		url.IsDeleted = true
		service.SaveURL(url)
	}
}

func (service MemoryService) GetURLsByUserToken(userToken string) []urlEntity.URL {
	inputCollection := service.Repository.GetURLsByUserToken(userToken)
	resultCollection := []urlEntity.URL{}

	for _, URL := range inputCollection {
		shortURLWithDomain := service.ShorteningService.GetShortURLWithDomain(URL.ShortURL)

		resultCollection = append(
			resultCollection,
			urlEntity.URL{
				LongURL:   URL.LongURL,
				ShortURL:  shortURLWithDomain,
				UserID:    URL.UserID,
				IsDeleted: URL.IsDeleted,
			},
		)
	}

	return resultCollection
}
