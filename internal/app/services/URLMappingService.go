package services

import (
	"github.com/PanovAlexey/url_carver/internal/app/domain/dto"
	"github.com/PanovAlexey/url_carver/internal/app/domain/entity/url"
)

type MappingService struct {
}

func (service MappingService) MapURLEntityCollectionToDTO(collection []url.URL) []dto.URLForShowingUser {
	collectionDto := []dto.URLForShowingUser{}

	for _, url := range collection {
		collectionDto = append(
			collectionDto,
			dto.NewURLForShowingUser(url.LongURL, url.ShortURL, ``),
		)
	}

	return collectionDto
}
