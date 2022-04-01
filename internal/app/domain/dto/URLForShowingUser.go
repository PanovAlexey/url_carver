package dto

type URLForShowingUser struct {
	LongURL   string `json:"original_url"`
	ShortURL  string `json:"short_url"`
	UserToken string `json:"-"`
}

func NewURLForShowingUser(longURL, shortURL, UserToken string) URLForShowingUser {
	return URLForShowingUser{
		LongURL:  longURL,
		ShortURL: shortURL,
	}
}

func (u URLForShowingUser) GetLongURL() string {
	return u.LongURL
}

func (u URLForShowingUser) GetShortURL() string {
	return u.ShortURL
}

func (u URLForShowingUser) GetUserToken() string {
	return u.UserToken
}

func (u *URLForShowingUser) SetShortURL(value string) {
	u.ShortURL = value
}
