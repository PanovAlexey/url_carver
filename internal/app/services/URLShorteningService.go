package services

import "fmt"

type shorteningServiceConfigInterface interface {
	GetBaseURL() string
}

type shorteningService struct {
	config shorteningServiceConfigInterface
}

func GetShorteningService(config shorteningServiceConfigInterface) *shorteningService {
	return &shorteningService{config: config}
}

func (service shorteningService) GetShortURLWithDomain(shortURLCode string) string {
	return service.config.GetBaseURL() + "/" + shortURLCode
}

func (service shorteningService) GetShortURLCode(longURL string) string {
	return fmt.Sprint(len(longURL) + 1)
}
