package repositories

import (
	urlPackage "github.com/PanovAlexey/url_carver/internal/app/domain/entity/url"
)

type shortURLs struct {
	urlMap map[string]urlPackage.URL
}

var globalURLs = shortURLs{
	urlMap: make(map[string]urlPackage.URL),
}

func GetURLMemoryRepository() *shortURLs {
	return &globalURLs
}

func (u *shortURLs) AddURL(key string, url string) bool {
	if u.IsExistURLByKey(key) {
		return false
	}

	u.urlMap[key] = urlPackage.New(url, key)


	return true
}

func (u shortURLs) GetURLByKey(key string) string {
	return u.urlMap[key].LongURL
}

func (u *shortURLs) IsExistURLByKey(key string) bool {
	_, ok := u.urlMap[key]
	return ok
}
