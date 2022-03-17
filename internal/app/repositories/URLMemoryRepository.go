package repositories

import (
	"github.com/PanovAlexey/url_carver/internal/app/domain"
	"github.com/PanovAlexey/url_carver/internal/app/domain/dto"
)

type shortURLs struct {
	urlMap map[string]domain.URLInterface
}

var globalURLs = shortURLs{
	urlMap: make(map[string]domain.URLInterface),
}

func GetURLMemoryRepository() *shortURLs {
	return &globalURLs
}

func (u *shortURLs) AddURL(url domain.URLInterface) bool {
	if u.IsExistURLByKey(url.GetShortURL()) {
		return false
	}

	u.urlMap[url.GetShortURL()] = url

	return true
}

func (u shortURLs) GetURLByKey(key string) string {
	return u.urlMap[key].GetLongURL()
}

func (u *shortURLs) IsExistURLByKey(key string) bool {
	_, ok := u.urlMap[key]
	return ok
}

func (u *shortURLs) GetURLsByUserID(userID string) dto.URLCollection {
	collection := dto.GetURLCollection()

	for _, url := range u.urlMap {
		if url.GetUserID() == userID {
			collection.AppendURL(url)
		}
	}

	return *collection
}
