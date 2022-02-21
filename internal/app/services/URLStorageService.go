package services

import "github.com/PanovAlexey/url_carver/internal/app/domain/entity/url"

type URLStorageService struct {
	config URLStorageConfigInterface
}

type URLStorageConfigInterface interface {
	GetURLStoragePath() string
}

func GetURLStorageService(config URLStorageConfigInterface) *URLStorageService {
	return &URLStorageService{config: config}
}

func (service URLStorageService) SaveURL(url url.URL) {

}
