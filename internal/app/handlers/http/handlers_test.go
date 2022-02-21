package http

import (
	"bytes"
	"github.com/PanovAlexey/url_carver/config"
	"github.com/PanovAlexey/url_carver/internal/app/repositories"
	"github.com/PanovAlexey/url_carver/internal/app/services"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testRequest(t *testing.T, ts *httptest.Server, method, path string, body []byte) (*http.Response, string) {
	bodyIoReader := bytes.NewBuffer(body)
	req, err := http.NewRequest(method, ts.URL+path, bodyIoReader)
	require.NoError(t, err)

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	response, err := client.Do(req)
	require.NoError(t, err)

	respBody, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	require.NoError(t, err)

	return response, string(respBody)
}

func Test_handleAddAndGetRequests(t *testing.T) {
	router := getRouterForRouteTest()
	config := config.New()
	server := httptest.NewServer(router)

	defer server.Close()

	type want struct {
		code              int
		response          string
		locationHeader    string
		contentTypeHeader string
	}

	tests := []struct {
		name    string
		urlPath string
		method  string
		body    []byte
		want    want
	}{
		{
			name:    "Negative test. Get short url by wrong url",
			urlPath: "/wrong/url/a",
			method:  http.MethodGet,
			body:    nil,
			want: want{
				code:              http.StatusNotFound,
				response:          ``,
				contentTypeHeader: "text/plain;charset=utf-8",
				locationHeader:    "",
			},
		},
		{
			name:    "Negative test. Get short url without code sending",
			urlPath: "/",
			method:  http.MethodGet,
			body:    nil,
			want: want{
				code:              http.StatusMethodNotAllowed,
				response:          ``,
				contentTypeHeader: "text/plain;charset=utf-8",
				locationHeader:    "",
			},
		},
		{
			name:    "Negative test. Get short url for Pikabu site before it was added.",
			urlPath: "/18",
			method:  http.MethodGet,
			body:    nil,
			want: want{
				code:              http.StatusNotFound,
				response:          ``,
				contentTypeHeader: "text/plain;charset=utf-8",
				locationHeader:    "",
			},
		},
		{
			name:    "Positive test. Add Yandex site url.",
			urlPath: "/",
			method:  http.MethodPost,
			body:    []byte(`http://ya.ru`),
			want: want{
				code:              http.StatusCreated,
				response:          config.GetBaseURL() + `/13`,
				contentTypeHeader: "text/plain;charset=utf-8",
				locationHeader:    "",
			},
		},
		{
			name:    "Positive test. Add Pikabu site url.",
			urlPath: "/",
			method:  http.MethodPost,
			body:    []byte(`http://pikabu.com`),
			want: want{
				code:              http.StatusCreated,
				response:          config.GetBaseURL() + `/18`,
				contentTypeHeader: "text/plain;charset=utf-8",
				locationHeader:    "",
			},
		},
		{
			name:    "Positive test. Get short url for Pikabu site after it was added.",
			urlPath: "/18",
			method:  http.MethodGet,
			body:    nil,
			want: want{
				code:              http.StatusTemporaryRedirect,
				response:          ``,
				contentTypeHeader: "text/plain;charset=utf-8",
				locationHeader:    "http://pikabu.com",
			},
		},
		{
			name:    "API. Negative test. Empty body.",
			urlPath: "/api/shorten",
			method:  http.MethodPost,
			body:    []byte(`{}`),
			want: want{
				code:              http.StatusBadRequest,
				response:          ``,
				contentTypeHeader: "application/json",
				locationHeader:    "",
			},
		},
		{
			name:    "API. Positive test. Add 3dnews site url.",
			urlPath: "/api/shorten",
			method:  http.MethodPost,
			body:    []byte(`{"url": "https://3dnews.com"}`),
			want: want{
				code:              http.StatusCreated,
				response:          `{"result":"` + config.GetBaseURL() + `/19"}`,
				contentTypeHeader: "application/json",
				locationHeader:    "",
			},
		},
		{
			name:    "Positive test. Get short url for 3dnews site after it was added.",
			urlPath: "/19",
			method:  http.MethodGet,
			body:    nil,
			want: want{
				code:              http.StatusTemporaryRedirect,
				response:          ``,
				contentTypeHeader: "text/plain;charset=utf-8",
				locationHeader:    "https://3dnews.com",
			},
		},
	}

	for _, testData := range tests {
		response, bodyString := testRequest(t, server, testData.method, testData.urlPath, testData.body)

		if response != nil {
			defer response.Body.Close()
		}

		assert.Equal(t, testData.want.code, response.StatusCode)
		assert.Equal(t, bodyString, testData.want.response)
		assert.Equal(t, response.Header.Get("Content-Type"), testData.want.contentTypeHeader)
		assert.Equal(t, response.Header.Get("location"), testData.want.locationHeader)
	}
}

func getRouterForRouteTest() chi.Router {
	URLMemoryRepository := repositories.GetURLMemoryRepository()
	config := config.New()
	shortURLService := services.GetShortURLService(URLMemoryRepository, config)
	fileStorageRepository, error := repositories.GetFileStorageRepository(config)

	if error != nil {
		log.Fatalln("error creating file repository by config:" + error.Error())
	}

	URLStorageService := services.GetURLStorageService(config, fileStorageRepository)
	httpHandler := GetHTTPHandler(shortURLService, URLStorageService)

	return httpHandler.NewRouter()
}
