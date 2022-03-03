package services

import (
	"encoding/json"
	"github.com/PanovAlexey/url_carver/internal/app/domain/dto"
	"github.com/PanovAlexey/url_carver/internal/app/domain/entity/url"
	"log"
)

type storageService struct {
	config            FileStorageConfigInterface
	storageRepository storageRepositoryInterface
}

type FileStorageConfigInterface interface {
	GetFileStoragePath() string
}

type storageRepositoryInterface interface {
	WriteLine(data []byte) error
	ReadLine() ([]byte, error)
	IsStorageExist() bool
}

func GetStorageService(
	config FileStorageConfigInterface,
	storageRepository storageRepositoryInterface,
) *storageService {
	return &storageService{config: config, storageRepository: storageRepository}
}

func (service storageService) GetURLCollectionFromStorage() dto.URLCollection {
	collection := dto.GetURLCollection()

	if service.storageRepository.IsStorageExist() {
		return *collection
	}

	for {
		data, err := service.storageRepository.ReadLine()

		if err != nil {
			log.Println("an error was found while parsing records from the storage: " + err.Error())
			break
		}

		if len(data) == 0 {
			log.Println("storage parsing successfully completed")
			break
		}

		url := url.URL{}
		err = json.Unmarshal(data, &url)

		if err != nil {
			log.Println("error while JSON parsing URL in storage: " + err.Error())
			break
		}

		collection.AppendURL(url)
	}

	return *collection
}

func (service storageService) SaveURL(url url.URL) {
	if !service.storageRepository.IsStorageExist() {
		return
	}

	data, err := json.Marshal(url)

	if err != nil {
		log.Fatalln("error occurred while marshalling URL to JSON format: " + err.Error())
	}

	err = service.storageRepository.WriteLine(data)

	if err != nil {
		log.Println("error occurred while saving URL to storage: " + err.Error())
	}
}