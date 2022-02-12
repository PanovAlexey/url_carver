package repositories

type shortURLs struct {
	urlMap map[string]string
}

var globalURLs = shortURLs{
	urlMap: make(map[string]string),
}

func GetEmailRepository() *shortURLs {
	return &globalURLs
}

func (u *shortURLs) AddEmail(key string, email string) bool {
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
	_, ok := u.urlMap[key]
	return ok
}
