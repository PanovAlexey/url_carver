package services

import (
	"fmt"
	"github.com/PanovAlexey/url_carver/internal/app/domain/dto"
	"github.com/PanovAlexey/url_carver/internal/app/domain/entity/url"
)

type RepositoryInterface interface {
	AddURL(key string, url string) bool
	GetURLByKey(key string) string
	IsExistURLByKey(key string) bool
}

type shortURLServiceConfigInterface interface {
	GetBaseURL() string
}

type shortURLService struct {
	repository RepositoryInterface
	config     shortURLServiceConfigInterface
}

func GetShortURLService(repository RepositoryInterface, config shortURLServiceConfigInterface) *shortURLService {
	return &shortURLService{repository: repository, config: config}
}

func (service shortURLService) GetURLByKey(key string) string {
	return service.repository.GetURLByKey(key)
}

func (service shortURLService) IsExistURLByKey(key string) bool {
	return service.repository.IsExistURLByKey(key)
}

func (service shortURLService) CreateLongURLDto() dto.LongURL {
	return dto.GetLongURLByValue("")
}

func (service shortURLService) GetURLByLongURLDto(longURLDto dto.LongURL) url.URL {
	return url.New(longURLDto.Value, service.cutAndAddURL(longURLDto.Value))
}

func (service shortURLService) GetShortURLDtoByURL(url url.URL) dto.ShortURL {
	return dto.GetShortURLByValue(url.ShortURL)
}

func (service shortURLService) cutAndAddURL(longURL string) string {
	shortURLCode := getShortURLCode(longURL)
	service.repository.AddURL(shortURLCode, longURL)

	return service.getShortURLWithDomain(shortURLCode)
}

func (service shortURLService) getShortURLWithDomain(shortURLCode string) string {
	return service.config.GetBaseURL() + "/" + shortURLCode
}

func getShortURLCode(longURL string) string {
	return fmt.Sprint(len(longURL) + 1)
}
