package dto

type DatabaseURLInterface interface {
	GetShortURL() string
	GetLongURL() string
	GetUserID() int
}
