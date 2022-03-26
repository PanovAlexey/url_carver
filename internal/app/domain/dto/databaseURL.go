package dto

type databaseURL struct {
	longURL  string
	shortURL string
	userID   int
}

func NewDatabaseURL(longURL, shortURL string, userID int) databaseURL {
	return databaseURL{
		longURL:  longURL,
		shortURL: shortURL,
		userID:   userID,
	}
}

func (databaseURL databaseURL) GetLongURL() string {
	return databaseURL.longURL
}

func (databaseURL databaseURL) GetShortURL() string {
	return databaseURL.shortURL
}

func (databaseURL databaseURL) GetUserID() int {
	return databaseURL.userID
}

func (databaseURL databaseURL) GetUserToken() string {
	return ``
}
