package services

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/PanovAlexey/url_carver/config"
	"github.com/PanovAlexey/url_carver/internal/app/domain/entity/url"
)

type ShorteningService struct {
	config config.Config
}

func GetShorteningService(config config.Config) *ShorteningService {
	return &ShorteningService{config: config}
}

func (service ShorteningService) GetShortURLWithDomain(shortURLCode string) (string, error) {
	return service.config.GetBaseURL() + "/" + shortURLCode, nil
}

func (service ShorteningService) GetURLEntityByLongURL(longURL string) (url.URL, error) {
	shortURLHash := md5.Sum([]byte(longURL))
	shortURLHashString := hex.EncodeToString(shortURLHash[:])

	return url.URL{
		LongURL:   longURL,
		ShortURL:  shortURLHashString,
		UserID:    ``,
		IsDeleted: false,
	}, nil
}
