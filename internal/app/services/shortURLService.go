package services

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

func (service shortURLService) AddEmail(key string, email string) bool {
	return service.repository.AddEmail(key, email)
}

func (service shortURLService) GetEmailByKey(key string) string {
	return service.repository.GetEmailByKey(key)
}

func (service shortURLService) IsExistEmailByKey(key string) bool {
	return service.repository.IsExistEmailByKey(key)
}
