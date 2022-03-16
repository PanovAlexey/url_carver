package url

type URL struct {
	LongURL  string
	ShortURL string
	UserID   string `json:"-"`
}

func New(longURL, shortURL string) URL {
	return URL{
		LongURL:  longURL,
		ShortURL: shortURL,
	}
}

func (url *URL) SetLongURL(value string) {
	url.LongURL = value
}

func (url *URL) SetShortURL(value string) {
	url.ShortURL = value
}

func (url *URL) SetUserId(value string) {
	url.UserID = value
}
