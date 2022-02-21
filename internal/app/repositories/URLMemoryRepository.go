package repositories

type shortURLs struct {
	urlMap map[string]string
}

var globalURLs = shortURLs{
	urlMap: make(map[string]string),
}

func GetURLMemoryRepository() *shortURLs {
	return &globalURLs
}

func (u *shortURLs) AddURL(key string, url string) bool {
	if u.IsExistURLByKey(key) {
		return false
	}

	u.urlMap[key] = url

	return true
}

func (u shortURLs) GetURLByKey(key string) string {
	return u.urlMap[key]
}

func (u *shortURLs) IsExistURLByKey(key string) bool {
	_, ok := u.urlMap[key]
	return ok
}
