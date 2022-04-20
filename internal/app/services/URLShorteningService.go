package services

import (
	"crypto/md5"
	"encoding/hex"
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
	shortURLHash := md5.Sum([]byte(longURL))
	shortURLHashString := hex.EncodeToString(shortURLHash[:])

	return url.New(longURL, shortURLHashString, ``, false), nil
}
