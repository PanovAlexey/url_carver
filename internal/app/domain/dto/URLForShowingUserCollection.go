package dto

type URLForShowingUserCollection struct {
	LongURL  string `json:"short_url"`
	ShortURL string `json:"original_url"`
	UserID   string `json:"-"`
}

func New(longURL, shortURL string) URLForShowingUserCollection {
	return URLForShowingUserCollection{
		LongURL:  longURL,
		ShortURL: shortURL,
	}
}

func (u URLForShowingUserCollection) GetLongURL() string {
	return u.LongURL
}

func (u URLForShowingUserCollection) GetShortURL() string {
	return u.ShortURL
}

func (u URLForShowingUserCollection) GetUserID() string {
	return u.ShortURL
}
