package dto

type DatabaseURL struct {
	longURL  string
	shortURL string
	userID   int
}

func NewDatabaseURL(longURL, shortURL string, userID int) DatabaseURL {
	return DatabaseURL{
		longURL:  longURL,
		shortURL: shortURL,
		userID:   userID,
	}
}

func (databaseURL DatabaseURL) GetLongURL() string {
	return databaseURL.longURL
}

func (databaseURL DatabaseURL) GetShortURL() string {
	return databaseURL.shortURL
}

func (databaseURL DatabaseURL) GetUserID() int {
	return databaseURL.userID
}

func (databaseURL DatabaseURL) GetUserToken() string {
	return ``
}
