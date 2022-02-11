package services

import "fmt"

type RepositoryInterface interface {
	AddEmail(key string, email string) bool
	GetEmailByKey(key string) string
	IsExistEmailByKey(key string) bool
}

type shortURLService struct {
	repository RepositoryInterface
}

func GetShortURLService(repository RepositoryInterface) *shortURLService {
	return &shortURLService{repository: repository}
}

func (service shortURLService) CutAndAddEmail(longURL string) string {
	shortURLCode := getShortURLCode(longURL)
	service.repository.AddEmail(shortURLCode, longURL)

	return getShortEmailWithDomain(shortURLCode)
}

func (service shortURLService) GetEmailByKey(key string) string {
	return service.repository.GetEmailByKey(key)
}

func (service shortURLService) IsExistEmailByKey(key string) bool {
	return service.repository.IsExistEmailByKey(key)
}

func getShortURLCode(longURL string) string {
	return fmt.Sprint(len(longURL) + 1)
}

func getShortEmailWithDomain(shortURLCode string) string {
	return "http://localhost:8080/" + shortURLCode // @ToDo: move to .env
}
