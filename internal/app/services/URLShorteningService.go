package services

import (
	"fmt"
	"github.com/PanovAlexey/url_carver/config"
	"github.com/PanovAlexey/url_carver/internal/app/domain/entity/url"
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

func (service shorteningService) GetURLEntityByLongURL(longURL string) (url.URL, error) {
	shortURL := fmt.Sprint(len(longURL) + 1)

	return url.New(longURL, shortURL, ``), nil
}
