package services

import (
	"fmt"
	"github.com/PanovAlexey/url_carver/config"
	"github.com/PanovAlexey/url_carver/internal/app/domain/dto"
	"github.com/PanovAlexey/url_carver/internal/app/domain/entity/url"
)

type repositoryInterface interface {
	AddURL(key string, url string) bool
	GetURLByKey(key string) string
	IsExistURLByKey(key string) bool
}

type shorteningServiceInterface interface {
	GetShortURLWithDomain(shortURLCode string) (string, error)
	GetShortURLCode(longURL string) (string, error)
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

func (service memoryService) GetURLByLongURLDto(longURLDto dto.LongURL) url.URL {
	return url.New(longURLDto.Value, service.cutAndAddURL(longURLDto.Value))
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
		service.repository.AddURL(url.ShortURL, url.LongURL)
	}
}

func (service memoryService) cutAndAddURL(longURL string) string {
	shortURLCode, err := service.shorteningService.GetShortURLCode(longURL)

	if err != nil {
		shortURLCode = ""
		fmt.Println("an error occurred while getting short URL code by long URL")
	}

	service.repository.AddURL(shortURLCode, longURL)

	return shortURLCode
}
