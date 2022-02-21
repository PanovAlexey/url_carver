package services

import "github.com/PanovAlexey/url_carver/internal/app/domain/entity/url"

type fileStorageService struct {
	config fileStorageConfigInterface
}

type fileStorageConfigInterface interface {
	GetFileStoragePath() string
}

func GetFileStorageService(config fileStorageConfigInterface) *fileStorageService {
	return &fileStorageService{config: config}
}

func (service fileStorageService) SaveURL(url url.URL) {

}
