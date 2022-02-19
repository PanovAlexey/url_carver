package dto

type ShortURL struct {
	Value string `json:"result"`
}

func GetShortURLByValue(value string) ShortURL {
	return ShortURL{Value: value}
}

func (shortUrl *ShortURL) SetValue(value string) {
	shortUrl.Value = value
}
