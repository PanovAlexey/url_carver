package url

type URL struct {
	LongURL  string
	ShortURL string
	UserID   string
}

func New(longURL, shortURL string) URL {
	return URL{
		LongURL:  longURL,
		ShortURL: shortURL,
	}
}

func (url URL) GetLongURL() string {
	return url.LongURL
}

func (url URL) GetShortURL() string {
	return url.ShortURL
}

func (url URL) GetUserID() string {
	return url.UserID
}

func (url *URL) SetLongURL(value string) {
	url.LongURL = value
}

func (url *URL) SetShortURL(value string) {
	url.ShortURL = value
}

func (url *URL) SetUserID(value string) {
	url.UserID = value
}
