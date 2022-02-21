package services

import (
	"fmt"
	"github.com/PanovAlexey/url_carver/internal/app/domain/dto"
	"github.com/PanovAlexey/url_carver/internal/app/domain/entity/url"
)

type RepositoryInterface interface {
	AddEmail(key string, email string) bool
	GetEmailByKey(key string) string
	IsExistEmailByKey(key string) bool
}

type configInterface interface {
	GetBaseURL() string
	GetPort() string
}

type shortURLService struct {
	repository RepositoryInterface
	config     configInterface
}

func GetShortURLService(repository RepositoryInterface, config configInterface) *shortURLService {
	return &shortURLService{repository: repository, config: config}
}

func (service shortURLService) GetEmailByKey(key string) string {
	return service.repository.GetEmailByKey(key)
}

func (service shortURLService) IsExistEmailByKey(key string) bool {
	return service.repository.IsExistEmailByKey(key)
}

func (service shortURLService) CreateLongURLDto() dto.LongURL {
	return dto.GetLongURLByValue("")
}

func (service shortURLService) GetURLByLongURLDto(longURLDto dto.LongURL) url.URL {
	return url.New(longURLDto.Value, service.cutAndAddEmail(longURLDto.Value))
}

func (service shortURLService) GetShortURLDtoByURL(url url.URL) dto.ShortURL {
	return dto.GetShortURLByValue(url.ShortURL)
}

func (service shortURLService) cutAndAddEmail(longURL string) string {
	shortURLCode := getShortURLCode(longURL)
	service.repository.AddEmail(shortURLCode, longURL)

	return service.getShortEmailWithDomain(shortURLCode)
}

func (service shortURLService) getShortEmailWithDomain(shortURLCode string) string {
	return service.config.GetBaseURL() + ":" + service.config.GetPort() + "/" + shortURLCode
}

func getShortURLCode(longURL string) string {
	return fmt.Sprint(len(longURL) + 1)
}
