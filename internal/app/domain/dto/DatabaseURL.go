package dto

type DatabaseURL struct {
	ID        int
	LongURL   string
	ShortURL  string
	UserID    int
	IsDeleted bool
}
