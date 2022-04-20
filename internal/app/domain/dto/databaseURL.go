package dto

type DatabaseURL struct {
	longURL   string
	shortURL  string
	userID    int
	IsDeleted bool
}

func NewDatabaseURL(longURL, shortURL string, userID int, isDeleted bool) DatabaseURL {
	return DatabaseURL{
		longURL:   longURL,
		shortURL:  shortURL,
		userID:    userID,
		IsDeleted: isDeleted,
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

func (databaseURL DatabaseURL) GetIsDeleted() bool {
	return databaseURL.IsDeleted
}
