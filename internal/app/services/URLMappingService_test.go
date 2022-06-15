package services

import (
	"fmt"
	"github.com/PanovAlexey/url_carver/internal/app/domain/entity/url"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetURLMappingService(t *testing.T) {
	t.Run("Test URL mapping creating", func(t *testing.T) {
		URLMappingService := GetURLMappingService()
		structType := fmt.Sprintf("%T", URLMappingService)

		assert.Equal(t, "*services.mappingService", structType)
	})
}

func Test_MapURLEntityCollectionToDTO(t *testing.T) {
	t.Run("Test map URL entity collection to DTO", func(t *testing.T) {
		URLMappingService := GetURLMappingService()

		URLSlice := make([]url.URL, 0)
		URLSlice = append(URLSlice, url.URL{
			LongURL:   "http://google.ru",
			ShortURL:  "38b92f",
			UserID:    `1`,
			IsDeleted: false,
		})

		URLForShowingUserDTOSlice := URLMappingService.MapURLEntityCollectionToDTO(URLSlice)

		assert.Equal(t, len(URLForShowingUserDTOSlice), 1)
	})
}
