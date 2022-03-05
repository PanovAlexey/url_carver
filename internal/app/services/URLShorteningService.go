package services

import (
	"fmt"
	"github.com/PanovAlexey/url_carver/config"
)

type shorteningService struct {
	config config.Config
}

func GetShorteningService(config config.Config) *shorteningService {
	return &shorteningService{config: config}
}

func (service shorteningService) GetShortURLWithDomain(shortURLCode string) (string, error) {
	return service.config.GetBaseURL() + "/" + shortURLCode, nil
}

func (service shorteningService) GetShortURLCode(longURL string) (string, error) {
	return fmt.Sprint(len(longURL) + 1), nil
}
