package url

type URL struct {
	LongURL  string
	ShortURL string
	UserID   string
}

func New(longURL, shortURL, userToken string) URL {
	return URL{
		LongURL:  longURL,
		ShortURL: shortURL,
		UserID:   userToken,
	}
}

func (url URL) GetLongURL() string {
	return url.LongURL
}

func (url URL) GetShortURL() string {
	return url.ShortURL
}

func (url URL) GetUserToken() string {
	return url.UserID
}

func (url *URL) SetLongURL(value string) {
	url.LongURL = value
}

func (url *URL) SetShortURL(value string) {
	url.ShortURL = value
}

func (url *URL) SetUserToken(value string) {
	url.UserID = value
}
