package domain

type URLInterface interface {
	GetShortURL() string
	GetLongURL() string
	GetUserToken() string
}
