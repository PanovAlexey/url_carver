package services

import (
	"errors"
	"fmt"
	databaseErrors "github.com/PanovAlexey/url_carver/internal/app/services/database/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetErrorService(t *testing.T) {
	t.Run("Test error service creating", func(t *testing.T) {
		errorService := GetErrorService()
		structType := fmt.Sprintf("%T", errorService)
		assert.Equal(t, structType, "services.ErrorService")
	})
}

func Test_GetActualizedError(t *testing.T) {
	t.Run("Test actualized error getting", func(t *testing.T) {
		errorService := GetErrorService()
		errorTestText := "test"
		err := errors.New(errorTestText)
		actualizedError := errorService.GetActualizedError(err, errorTestText)

		assert.Equal(t, actualizedError.Error(), errorTestText)
	})
}

func Test_IsKeyDuplicated(t *testing.T) {
	t.Run("Test is key duplicated", func(t *testing.T) {
		errorService := GetErrorService()
		errorDuplicate := fmt.Errorf("%v: %w", "", databaseErrors.ErrorDuplicateKey)

		assert.Equal(t, true, errorService.IsKeyDuplicated(errorDuplicate))
	})
}

func Test_IsDeleted(t *testing.T) {
	t.Run("Test is deleted", func(t *testing.T) {
		errorService := GetErrorService()
		errorDuplicate := fmt.Errorf("%v: %w", "", databaseErrors.ErrorDuplicateKey)
		errorIsDeleted := fmt.Errorf("%v: %w", "", databaseErrors.ErrorIsDeleted)

		assert.Equal(t, false, errorService.IsDeleted(errorDuplicate))
		assert.Equal(t, true, errorService.IsDeleted(errorIsDeleted))
	})
}
