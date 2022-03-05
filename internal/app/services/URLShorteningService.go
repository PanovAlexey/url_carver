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

func (service shorteningService) GetShortURLWithDomain(shortURLCode string) string {
	return service.config.GetBaseURL() + "/" + shortURLCode
}

func (service shorteningService) GetShortURLCode(longURL string) string {
	return fmt.Sprint(len(longURL) + 1)
}
