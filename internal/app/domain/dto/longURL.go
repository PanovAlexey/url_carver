package dto

// LongURL - entity to pass the source URL
type LongURL struct {
	value string `json:"url"`
}

func GetLongURLByValue(value string) LongURL {
	return LongURL{value: value}
}

func (longUrl *LongURL) SetValue(value string) {
	longUrl.value = value
}

func (longUrl *LongURL) GetValue() string {
	return longUrl.value
}
