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
