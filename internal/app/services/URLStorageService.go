package services

import (
	"encoding/json"
	"github.com/PanovAlexey/url_carver/config"
	"github.com/PanovAlexey/url_carver/internal/app/domain/dto"
	"github.com/PanovAlexey/url_carver/internal/app/domain/entity/url"
	"log"
)

type storageService struct {
	config            config.Config
	storageRepository storageRepositoryInterface
}

type storageRepositoryInterface interface {
	WriteLine(data []byte) error
	ReadLine() ([]byte, error)
	IsStorageExist() (bool, error)
}

func GetStorageService(
	config config.Config,
	storageRepository storageRepositoryInterface,
) *storageService {
	return &storageService{config: config, storageRepository: storageRepository}
}

func (service storageService) GetURLCollectionFromStorage() dto.URLCollection {
	collection := dto.GetURLCollection()

	isStorageExist, err := service.storageRepository.IsStorageExist()

	if !isStorageExist || err != nil {
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
	isStorageExist, err := service.storageRepository.IsStorageExist()

	if !isStorageExist || err != nil {
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
