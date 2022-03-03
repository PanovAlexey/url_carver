package services

import (
	"github.com/PanovAlexey/url_carver/internal/app/domain/dto"
	"github.com/PanovAlexey/url_carver/internal/app/domain/entity/url"
)

type repositoryInterface interface {
	AddURL(key string, url string) bool
	GetURLByKey(key string) string
	IsExistURLByKey(key string) bool
}

type memoryServiceConfigInterface interface {
	GetBaseURL() string
}

type shorteningServiceInterface interface {
	GetShortURLWithDomain(shortURLCode string) string
	GetShortURLCode(longURL string) string
}

type memoryService struct {
	repository        repositoryInterface
	config            memoryServiceConfigInterface
	shorteningService shorteningServiceInterface
}

func GetMemoryService(
	config memoryServiceConfigInterface,
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
	return dto.GetShortURLByValue(service.shorteningService.GetShortURLWithDomain(url.ShortURL))
}

func (service memoryService) LoadURLs(collection dto.URLCollection) {
	for _, url := range collection.GetCollection() {
		service.repository.AddURL(url.ShortURL, url.LongURL)
	}
}

func (service memoryService) cutAndAddURL(longURL string) string {
	shortURLCode := service.shorteningService.GetShortURLCode(longURL)
	service.repository.AddURL(shortURLCode, longURL)

	return shortURLCode
}
