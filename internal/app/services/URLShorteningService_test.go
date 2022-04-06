package services

import (
	"github.com/PanovAlexey/url_carver/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetURLEntityByLongURL(t *testing.T) {
	config := config.New()
	shorteningService := GetShorteningService(config)

	tests := []struct {
		name  string
		value string
		want  string
	}{
		{
			name:  "Test by vk.com service with http scheme",
			value: "vk.com",
			want:  "d2b6043e84b2aebc95a2faf382bde230",
		},
		{
			name:  "Test by mamba service with http scheme",
			value: "http://mamba.ru",
			want:  "1bcfcaf3e204762c138590c951995d49",
		},
		{
			name:  "Test by facebook service with https scheme",
			value: "https://facebook.com",
			want:  "a023cfbf5f1c39bdf8407f28b60cd134",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			URL, err := shorteningService.GetURLEntityByLongURL(tt.value)
			assert.Equal(t, err, nil)
			assert.Equal(t, URL.ShortURL, tt.want)
		})
	}
}

func Test_GetShortURLWithDomain(t *testing.T) {
	config := config.New()
	shorteningService := GetShorteningService(config)

	tests := []struct {
		name  string
		value string
		want  string
	}{
		{
			name:  "Test by num",
			value: "7",
			want:  config.GetBaseURL() + "/7",
		},
		{
			name:  "Test by string",
			value: "azazam",
			want:  config.GetBaseURL() + "/azazam",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shortURLWithDomain, err := shorteningService.GetShortURLWithDomain(tt.value)
			assert.Equal(t, shortURLWithDomain, tt.want)
			assert.Equal(t, err, nil)
		})
	}
}
