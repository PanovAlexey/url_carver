package services

import (
	"github.com/PanovAlexey/url_carver/internal/app/domain/dto"
	"github.com/PanovAlexey/url_carver/internal/app/domain/entity/url"
)

type mappingService struct {
}

func GetURLMappingService() *mappingService {
	return &mappingService{}
}

func (service mappingService) MapURLEntityCollectionToDTO(collection []url.URL) []dto.URLForShowingUser {
	collectionDto := []dto.URLForShowingUser{}

	for _, url := range collection {
		collectionDto = append(
			collectionDto,
			dto.NewURLForShowingUser(url.GetLongURL(), url.GetShortURL(), ``),
		)
	}

	return collectionDto
}
