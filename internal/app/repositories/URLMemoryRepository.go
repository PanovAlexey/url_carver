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

func (u *shortURLs) AddURL(url urlPackage.URL) bool {
	if u.IsExistURLByKey(url.ShortURL) {
		return false
	}

	u.urlMap[url.ShortURL] = url


	return true
}

func (u shortURLs) GetURLByKey(key string) string {
	return u.urlMap[key].LongURL
}

func (u *shortURLs) IsExistURLByKey(key string) bool {
	_, ok := u.urlMap[key]
	return ok
}
