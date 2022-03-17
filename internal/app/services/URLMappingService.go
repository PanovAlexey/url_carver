package services

import (
	"github.com/PanovAlexey/url_carver/internal/app/domain/dto"
)

type mappingService struct {
}

func GetURLMappingService() *mappingService {
	return &mappingService{}
}

func (service mappingService) MapURLEntityCollectionToDTO(collection dto.URLCollection) dto.URLCollection {
	collectionDto := dto.GetURLCollection()

	for _, url := range collection.GetCollection() {
		collectionDto.AppendURL(dto.New(url.GetLongURL(), url.GetShortURL(), ``))
	}

	return *collectionDto
}
