package repositories

type shortURLs struct {
	urlMap map[string]string
}

var globalURLs = shortURLs{}

func GetEmailRepository() *shortURLs {
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
