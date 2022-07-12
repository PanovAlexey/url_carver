package repositories

import (
	"github.com/PanovAlexey/url_carver/internal/app/domain/entity/url"
)

type URLMemoryRepository struct {
	urlMap map[string]url.URL
}

var globalURLs = URLMemoryRepository{
	urlMap: make(map[string]url.URL),
}

func GetURLMemoryRepository() *URLMemoryRepository {
	return &globalURLs
}

func (u *URLMemoryRepository) AddURL(url url.URL) bool {
	u.urlMap[url.ShortURL] = url

	return true
}

func (u URLMemoryRepository) GetURLByKey(key string) url.URL {
	return u.urlMap[key]
}

func (u URLMemoryRepository) GetURLsByShortValueSlice(urlShortValuesSlice []string) []url.URL {
	urlCollection := make([]url.URL, 0)

	for _, urlShortValue := range urlShortValuesSlice {
		if u.IsExistURLByKey(urlShortValue) {
			urlCollection = append(urlCollection, u.GetURLByKey(urlShortValue))
		}
	}

	return urlCollection
}

func (u URLMemoryRepository) GetAllURLs() []url.URL {
	var collection []url.URL

	for _, url := range u.urlMap {
		collection = append(collection, url)
	}

	return collection
}

func (u *URLMemoryRepository) IsExistURLByKey(key string) bool {
	_, ok := u.urlMap[key]
	return ok
}

func (u *URLMemoryRepository) GetURLsByUserToken(userToken string) []url.URL {
	var collection []url.URL

	for _, url := range u.urlMap {
		if url.UserID == userToken {
			collection = append(collection, url)
		}
	}

	return collection
}
