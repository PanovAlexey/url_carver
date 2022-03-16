package services

import (
	"fmt"
	"github.com/PanovAlexey/url_carver/config"
	"github.com/PanovAlexey/url_carver/internal/app/domain"
	"github.com/PanovAlexey/url_carver/internal/app/domain/dto"
	"github.com/PanovAlexey/url_carver/internal/app/domain/entity/url"
)

type repositoryInterface interface {
	AddURL(url domain.URLInterface) bool
	GetURLByKey(key string) string
	IsExistURLByKey(key string) bool
	GetURLsByUserId(userId string) dto.URLCollection
}

type shorteningServiceInterface interface {
	GetShortURLWithDomain(shortURLCode string) (string, error)
	GetURLEntityByLongURL(longURL string) (url.URL, error)
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

func (service memoryService) GetShortURLDtoByURL(url url.URL) dto.ShortURL {
	shortURLWithDomain, err := service.shorteningService.GetShortURLWithDomain(url.ShortURL)

	if err != nil {
		shortURLWithDomain = ""
		fmt.Println("impossible to build a short url with domain.")
	}

	return dto.GetShortURLByValue(shortURLWithDomain)
}

func (service memoryService) LoadURLs(collection dto.URLCollection) {
	for _, url := range collection.GetCollection() {
		service.SaveURL(url)
	}
}

func (service memoryService) SaveURL(url domain.URLInterface) bool {
	return service.repository.AddURL(url)
}

func (service memoryService) GetURLsByUserId(userId string) dto.URLCollection {
	return service.repository.GetURLsByUserId(userId)
}
