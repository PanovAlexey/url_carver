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
		structType := fmt.Sprintf("%T", ErrorService{})
		assert.Equal(t, structType, "services.ErrorService")
	})
}

func Test_GetActualizedError(t *testing.T) {
	t.Run("Test actualized error getting", func(t *testing.T) {
		errorTestText := "test"
		err := errors.New(errorTestText)
		actualizedError := ErrorService{}.GetActualizedError(err, errorTestText)

		assert.Equal(t, actualizedError.Error(), errorTestText)
	})
}

func Test_IsKeyDuplicated(t *testing.T) {
	t.Run("Test is key duplicated", func(t *testing.T) {
		errorDuplicate := fmt.Errorf("%v: %w", "", databaseErrors.ErrorDuplicateKey)

		assert.Equal(t, true, ErrorService{}.IsKeyDuplicated(errorDuplicate))
	})
}

func Test_IsDeleted(t *testing.T) {
	t.Run("Test is deleted", func(t *testing.T) {
		errorService := ErrorService{}
		errorDuplicate := fmt.Errorf("%v: %w", "", databaseErrors.ErrorDuplicateKey)
		errorIsDeleted := fmt.Errorf("%v: %w", "", databaseErrors.ErrorIsDeleted)

		assert.Equal(t, false, errorService.IsDeleted(errorDuplicate))
		assert.Equal(t, true, errorService.IsDeleted(errorIsDeleted))
	})
}
