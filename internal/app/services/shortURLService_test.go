package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getShortURLCode(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{
			name:  "Test by vk.com service with http scheme",
			value: "vk.com",
			want:  "7",
		},
		{
			name:  "Test by mamba service with http scheme",
			value: "http://mamba.ru",
			want:  "16",
		},
		{
			name:  "Test by facebook service with https scheme",
			value: "https://facebook.com",
			want:  "21",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, getShortURLCode(tt.value), tt.want)
		})
	}
}

func Test_getShortEmailWithDomain(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{
			name:  "Test by num",
			value: "7",
			want:  "http://localhost:8080/7",
		},
		{
			name:  "Test by string",
			value: "azazam",
			want:  "http://localhost:8080/azazam",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, getShortEmailWithDomain(tt.value), tt.want)
		})
	}
}
