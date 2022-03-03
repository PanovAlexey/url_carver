package dto

// LongURL - entity to pass the source URL
type LongURL struct {
	Value string `json:"url"`
}

func GetLongURLByValue(value string) LongURL {
	return LongURL{Value: value}
}

func (longUrl *LongURL) SetValue(value string) {
	longUrl.Value = value
}
