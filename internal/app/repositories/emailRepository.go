package repositories

type shortURLs struct {
	urlMap map[string]string
}

var globalURLs = shortURLs{}

func GetEmailRepository() *shortURLs {
	InitialURLMap()

	return &globalURLs
}

func (u *shortURLs) AddEmail(key string, email string) bool {
	if u.urlMap == nil {
		u.urlMap = make(map[string]string)
	}

	if u.IsExistEmailByKey(key) {
		return false
	}

	u.urlMap[key] = email

	return true
}

func (u shortURLs) GetEmailByKey(key string) string {
	return u.urlMap[key]
}

func (u *shortURLs) IsExistEmailByKey(key string) bool {
	if _, ok := u.urlMap[key]; ok {
		return true
	}

	return false
}

func InitialURLMap() {
	globalURLs.AddEmail("yandex", "http://www.yandex.ru")
	globalURLs.AddEmail("google", "http://www.google.com")
	globalURLs.AddEmail("meta", "http://about.facebook.com/meta/")
}
