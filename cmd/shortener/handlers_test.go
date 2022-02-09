package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_getShortURLByLongURL(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{
			name:  "Test by vk.com service without scheme",
			value: "vk.com",
			want:  "http://localhost:8080/7",
		},
		{
			name:  "Test by mamba service with http scheme",
			value: "http://mamba.ru",
			want:  "http://localhost:8080/16",
		},
		{
			name:  "Test by facebook service with https scheme",
			value: "https://facebook.com",
			want:  "http://localhost:8080/21",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, getShortURLByLongURL(tt.value), tt.want)
		})
	}
}

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


func Test_handlePostRequest(t *testing.T) {
	type want struct {
		code        int
		response    string
		contentType string
	}
	tests := []struct {
		name    string
		urlPath string
		method  string
		body    []byte
		want    want
	}{
		{
			name:    "Positive test. Cut long facebook URL",
			urlPath: "/",
			method:  http.MethodPost,
			body:    []byte(`http://ya.ru`),
			want: want{
				code:        http.StatusCreated,
				response:    `http://localhost:8080/13`,
				contentType: "text/plain;charset=utf-8",
			},
		},
		{
			name:    "Negative test. Empty body.",
			urlPath: "/",
			method:  http.MethodPost,
			body:    nil,
			want: want{
				code:        http.StatusBadRequest,
				response:    ``,
				contentType: "text/plain;charset=utf-8",
			},
		},
		{
			name:    "Negative test. Not main route.",
			urlPath: "/a/b/c",
			method:  http.MethodPost,
			body:    []byte(`http://ya.ru`),
			want: want{
				code:        http.StatusBadRequest,
				response:    ``,
				contentType: "text/plain;charset=utf-8",
			},
		},
	}
	for _, testData := range tests {
		t.Run(testData.name, func(t *testing.T) {
			request := httptest.NewRequest(testData.method, testData.urlPath, bytes.NewBuffer(testData.body))

			mainHandler := MainHandler{
				URL: getInitialURLMap(),
			}
			w := httptest.NewRecorder()

			mainHandler.ServeHTTP(w, request)
			res := w.Result()

			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)

			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, res.StatusCode, testData.want.code)
			assert.Equal(t, res.Header.Get("Content-Type"), testData.want.contentType)
			assert.Equal(t, string(resBody), testData.want.response)
		})
	}
}
