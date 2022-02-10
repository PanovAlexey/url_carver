package main

var GlobalURLs = URLs{}

type URLs struct {
	urlMap map[string]string
}

func (u *URLs) AddEmail(key string, email string) bool {
	if u.urlMap == nil {
		u.urlMap = make(map[string]string)
	}

	if u.IsExistEmailByKey(key) {
		return false
	}

	u.urlMap[key] = email

	return true
}

func (u URLs) GetEmailByKey(key string) string {
	return u.urlMap[key]
}

func (u *URLs) IsExistEmailByKey(key string) bool {
	if _, ok := u.urlMap[key]; ok {
		return true
	}

	return false
}

func InitialURLMap() {
	GlobalURLs.AddEmail("yandex", "http://www.yandex.ru")
	GlobalURLs.AddEmail("google", "http://www.google.com")
	GlobalURLs.AddEmail("meta", "http://about.facebook.com/meta/")
}
