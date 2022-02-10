package services

type StorageInterface interface {
	AddEmail(key string, email string) bool
	GetEmailByKey(key string) string
	IsExistEmailByKey(key string) bool
}

type shortURLService struct {
	storage StorageInterface
}

func GetShortURLService(storage StorageInterface) *shortURLService {
	return &shortURLService{storage: storage}
}

func (service shortURLService) AddEmail(key string, email string) bool {
	return service.storage.AddEmail(key, email)
}

func (service shortURLService) GetEmailByKey(key string) string {
	return service.storage.GetEmailByKey(key)
}

func (service shortURLService) IsExistEmailByKey(key string) bool {
	return service.storage.IsExistEmailByKey(key)
}
