package services

import (
	"encoding/json"
	"github.com/PanovAlexey/url_carver/internal/app/domain/entity/url"
	"log"
)

type URLStorageService struct {
	config            FileStorageConfigInterface
	storageRepository storageRepositoryInterface
}

type FileStorageConfigInterface interface {
	GetFileStoragePath() string
}

type storageRepositoryInterface interface {
	WriteLine(data []byte) error
}

func GetURLStorageService(
	config FileStorageConfigInterface,
	storageRepository storageRepositoryInterface,
) *URLStorageService {
	return &URLStorageService{config: config, storageRepository: storageRepository}
}

func (service URLStorageService) SaveURL(url url.URL) {
	data, err := json.Marshal(url)

	if err != nil {
		log.Fatalln("error occurred while marshalling URL to JSON format: " + err.Error())
	}

	service.storageRepository.WriteLine(data)
}
