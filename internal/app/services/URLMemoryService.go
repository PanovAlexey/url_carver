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

type shorteningServiceInterface interface {
	GetShortURLWithDomain(shortURLCode string) (string, error)
	GetURLEntityByLongURL(longURL string) (urlEntity.URL, error)
}

type memoryService struct {
	repository        repositoryInterface
	config            config.Config
	shorteningService shorteningServiceInterface
}

func GetMemoryService(
	config config.Config,
	repository repositoryInterface,
	shorteningService shorteningServiceInterface,
) *memoryService {
	return &memoryService{config: config, repository: repository, shorteningService: shorteningService}
}

func (service memoryService) GetURLEntityByShortURL(shortURL string) (string, error) {
	url := service.repository.GetURLByKey(shortURL)

	if url.IsDeleted {
		return "", fmt.Errorf("%v: %w", url, databaseErrors.ErrorIsDeleted)
	}

	return url.LongURL, nil
}

func (service memoryService) IsExistURLEntityByShortURL(shortURL string) bool {
	return service.repository.IsExistURLByKey(shortURL)
}

func (service memoryService) CreateLongURLDto() dto.LongURL {
	return dto.GetLongURLByValue("")
}

func (service memoryService) GetShortURLDtoByURL(url urlEntity.URL) dto.ShortURL {
	shortURLWithDomain, err := service.shorteningService.GetShortURLWithDomain(url.ShortURL)

	if err != nil {
		shortURLWithDomain = ""
		fmt.Println("impossible to build a short url with domain.")
	}

	return dto.GetShortURLByValue(shortURLWithDomain)
}

func (service memoryService) LoadURLs(collection []urlEntity.URL) {
	for _, url := range collection {
		service.SaveURL(url)
	}
}

func (service memoryService) SaveURL(url urlEntity.URL) bool {
	return service.repository.AddURL(url)
}

func (service memoryService) DeleteURLsByShortValueSlice(urlShortValuesSlice []string) {
	urlCollection := service.repository.GetURLsByShortValueSlice(urlShortValuesSlice)

	for _, url := range urlCollection {
		url.IsDeleted = true
		service.SaveURL(url)
	}
}

func (service memoryService) GetURLsByUserToken(userToken string) []urlEntity.URL {
	inputCollection := service.repository.GetURLsByUserToken(userToken)
	resultCollection := []urlEntity.URL{}

	for _, URL := range inputCollection {
		shortURLWithDomain, err := service.shorteningService.GetShortURLWithDomain(URL.ShortURL)

		if err != nil {
			shortURLWithDomain = ""
			fmt.Println("impossible to build a short url with domain.")
		}

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
