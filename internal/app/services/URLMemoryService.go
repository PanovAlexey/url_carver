package services

import (
	"fmt"
	"github.com/PanovAlexey/url_carver/config"
	"github.com/PanovAlexey/url_carver/internal/app/domain/dto"
	urlEntity "github.com/PanovAlexey/url_carver/internal/app/domain/entity/url"
)

type repositoryInterface interface {
	AddURL(url urlEntity.URL) bool
	GetURLByKey(key string) string
	IsExistURLByKey(key string) bool
	GetURLsByUserToken(userToken string) dto.URLCollection
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

func (service memoryService) GetURLByKey(key string) string {
	return service.repository.GetURLByKey(key)
}

func (service memoryService) IsExistURLByKey(key string) bool {
	return service.repository.IsExistURLByKey(key)
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

func (service memoryService) LoadURLs(collection dto.URLCollection) {
	for _, url := range collection.GetCollection() {
		service.SaveURL(urlEntity.New(url.GetLongURL(), url.GetShortURL(), url.GetUserToken()))
	}
}

func (service memoryService) SaveURL(url urlEntity.URL) bool {
	return service.repository.AddURL(url)
}

func (service memoryService) GetURLsByUserToken(userToken string) dto.URLCollection {
	inputCollection := service.repository.GetURLsByUserToken(userToken)
	resultCollection := dto.URLCollection{}

	for _, URL := range inputCollection.GetCollection() {
		shortURLWithDomain, err := service.shorteningService.GetShortURLWithDomain(URL.GetShortURL())

		if err != nil {
			shortURLWithDomain = ""
			fmt.Println("impossible to build a short url with domain.")
		}

		resultCollection.AppendURL(
			dto.New(
				URL.GetLongURL(),
				shortURLWithDomain,
				URL.GetUserToken(),
			))
	}

	return resultCollection
}
