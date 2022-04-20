package services

import (
	"errors"
	"fmt"
	databaseErrors "github.com/PanovAlexey/url_carver/internal/app/services/database/errors"
	"strings"
)

const keyDuplicate = "23505"

type ErrorService struct {
}

func GetErrorService() ErrorService {
	return ErrorService{}
}

func (service ErrorService) GetActualizedError(err error, additionalInfo interface{}) error {
	if strings.Contains(err.Error(), keyDuplicate) {
		return fmt.Errorf("%v: %w", additionalInfo, databaseErrors.ErrorDuplicateKey)
	}

	return err
}

func (service ErrorService) IsKeyDuplicated(err error) bool {
	return errors.Is(err, databaseErrors.ErrorDuplicateKey)
}

func (service ErrorService) IsDeleted(err error) bool {
	return errors.Is(err, databaseErrors.ErrorIsDeleted)
}
