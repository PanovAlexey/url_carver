package services

import (
	"encoding/json"
	"github.com/PanovAlexey/url_carver/internal/app/domain/dto"
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
	ReadLine() ([]byte, error)
}

func GetURLStorageService(
	config FileStorageConfigInterface,
	storageRepository storageRepositoryInterface,
) *URLStorageService {
	return &URLStorageService{config: config, storageRepository: storageRepository}
}

func (service URLStorageService) GetURLCollectionFromStorage() dto.URLCollection {
	collection := dto.GetURLCollection()

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

func (service URLStorageService) SaveURL(url url.URL) {
	data, err := json.Marshal(url)

	if err != nil {
		log.Fatalln("error occurred while marshalling URL to JSON format: " + err.Error())
	}

	err = service.storageRepository.WriteLine(data)

	if err != nil {
		log.Println("error occurred while saving URL to storage: " + err.Error())
	}
}
