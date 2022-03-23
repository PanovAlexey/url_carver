package dto

type URLForShowingUserCollection struct {
	LongURL   string `json:"original_url"`
	ShortURL  string `json:"short_url"`
	UserToken string `json:"-"`
}

func New(longURL, shortURL, UserToken string) URLForShowingUserCollection {
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

func (u URLForShowingUserCollection) GetUserToken() string {
	return u.UserToken
}

func (u *URLForShowingUserCollection) SetShortURL(value string) {
	u.ShortURL = value
}
