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

type URLMemoryServiceConfigInterface interface {
	GetBaseURL() string
}

type shorteningServiceInterface interface {
	GetShortURLWithDomain(shortURLCode string) string
	GetShortURLCode(longURL string) string
}

type URLMemoryService struct {
	repository        repositoryInterface
	config            URLMemoryServiceConfigInterface
	shorteningService shorteningServiceInterface
}

func GetURLMemoryService(
	config URLMemoryServiceConfigInterface,
	repository repositoryInterface,
	shorteningService shorteningServiceInterface,
) *URLMemoryService {
	return &URLMemoryService{config: config, repository: repository, shorteningService: shorteningService}
}

func (service URLMemoryService) GetURLByKey(key string) string {
	return service.repository.GetURLByKey(key)
}

func (service URLMemoryService) IsExistURLByKey(key string) bool {
	return service.repository.IsExistURLByKey(key)
}

func (service URLMemoryService) CreateLongURLDto() dto.LongURL {
	return dto.GetLongURLByValue("")
}

func (service URLMemoryService) GetURLByLongURLDto(longURLDto dto.LongURL) url.URL {
	return url.New(longURLDto.Value, service.cutAndAddURL(longURLDto.Value))
}

func (service URLMemoryService) GetShortURLDtoByURL(url url.URL) dto.ShortURL {
	return dto.GetShortURLByValue(url.ShortURL)
}

func (service URLMemoryService) cutAndAddURL(longURL string) string {
	shortURLCode := service.shorteningService.GetShortURLCode(longURL)
	service.repository.AddURL(shortURLCode, longURL)

	return service.shorteningService.GetShortURLWithDomain(shortURLCode)
}
