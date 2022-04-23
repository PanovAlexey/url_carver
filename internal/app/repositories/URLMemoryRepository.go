package repositories

import (
	"github.com/PanovAlexey/url_carver/internal/app/domain/entity/url"
)

type shortURLs struct {
	urlMap map[string]url.URL
}

var globalURLs = shortURLs{
	urlMap: make(map[string]url.URL),
}

func GetURLMemoryRepository() *shortURLs {
	return &globalURLs
}

func (u *shortURLs) AddURL(url url.URL) bool {
	u.urlMap[url.GetShortURL()] = url

	return true
}

func (u shortURLs) GetURLByKey(key string) url.URL {
	return u.urlMap[key]
}

func (u shortURLs) GetURLsByShortValueSlice(urlShortValuesSlice []string) []url.URL {
	urlCollection := make([]url.URL, 0)

	for _, urlShortValue := range urlShortValuesSlice {
		if u.IsExistURLByKey(urlShortValue) {
			urlCollection = append(urlCollection, u.GetURLByKey(urlShortValue))
		}
	}

	return urlCollection
}

func (u *shortURLs) IsExistURLByKey(key string) bool {
	_, ok := u.urlMap[key]
	return ok
}

func (u *shortURLs) GetURLsByUserToken(userToken string) []url.URL {
	var collection []url.URL

	for _, url := range u.urlMap {
		if url.GetUserToken() == userToken {
			collection = append(collection, url)
		}
	}

	return collection
}
