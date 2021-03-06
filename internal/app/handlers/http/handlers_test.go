package http

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/PanovAlexey/url_carver/config"
	"github.com/PanovAlexey/url_carver/internal/app/repositories"
	"github.com/PanovAlexey/url_carver/internal/app/services"
	"github.com/PanovAlexey/url_carver/internal/app/services/database"
	"github.com/PanovAlexey/url_carver/internal/app/services/encryption"
	"github.com/PanovAlexey/url_carver/internal/app/tests"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testRequest(
	t *testing.T, ts *httptest.Server, method, path string, body []byte, headers map[string]string, cookies []http.Cookie,
) (*http.Response, string) {
	bodyIoReader := bytes.NewBuffer(body)
	req, err := http.NewRequest(method, ts.URL+path, bodyIoReader)
	require.NoError(t, err)

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	for _, cookie := range cookies {
		req.AddCookie(&cookie)
	}

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
		contentEncoding   string
	}

	tests := []struct {
		name               string
		urlPath            string
		method             string
		headers            map[string]string
		cookies            []http.Cookie
		isCookiesFinalized bool
		body               []byte
		want               want
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
				response:          config.GetBaseURL() + `/1b556b44a4ee73524fb009e11918fb4f`,
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
				response:          config.GetBaseURL() + `/6e8040b1b452170cc569a17ad1c23e13`,
				contentTypeHeader: "text/plain;charset=utf-8",
				locationHeader:    "",
			},
		},
		{
			name:    "Positive test. Get short url for Pikabu site after it was added.",
			urlPath: "/6e8040b1b452170cc569a17ad1c23e13",
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
				response:          `{"result":"` + config.GetBaseURL() + `/f325fe609335839151c83c065975e95e"}`,
				contentTypeHeader: "application/json",
				locationHeader:    "",
			},
		},
		{
			name:    "Positive test. Get short url for 3dnews site after it was added.",
			urlPath: "/f325fe609335839151c83c065975e95e",
			method:  http.MethodGet,
			body:    nil,
			want: want{
				code:              http.StatusTemporaryRedirect,
				response:          ``,
				contentTypeHeader: "text/plain;charset=utf-8",
				locationHeader:    "https://3dnews.com",
			},
		},
		{
			name:    "Positive test. Get short url for 3dnews site after it was added with GZIP compression.",
			urlPath: "/f325fe609335839151c83c065975e95e",
			method:  http.MethodGet,
			body:    nil,
			headers: map[string]string{
				"Accept-Encoding": "gzip",
			},
			want: want{
				code:              http.StatusTemporaryRedirect,
				response:          "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00",
				contentTypeHeader: "text/plain;charset=utf-8",
				locationHeader:    "https://3dnews.com",
				contentEncoding:   "gzip",
			},
		},
		{
			name:    "Positive test. Add Pikabu site url with GZIP compression.",
			urlPath: "/",
			method:  http.MethodPost,
			body:    []byte(`http://pikabu.com`),
			headers: map[string]string{
				"Accept-Encoding": "gzip",
			},
			want: want{
				code:              http.StatusCreated,
				contentTypeHeader: "text/plain;charset=utf-8",
				locationHeader:    "",
				contentEncoding:   "gzip",
			},
		},
		{
			name:    "API. Positive test. Add 3dnews site url with GZIP compression.",
			urlPath: "/api/shorten",
			method:  http.MethodPost,
			body:    []byte(`{"url": "https://3dnews.com"}`),
			headers: map[string]string{
				"Accept-Encoding": "gzip",
			},
			want: want{
				code:              http.StatusCreated,
				contentTypeHeader: "application/json",
				locationHeader:    "",
				contentEncoding:   "gzip",
			},
		},
		{
			name:    "API. Positive test. Add batch URLs.",
			urlPath: "/api/shorten/batch",
			method:  http.MethodPost,
			body: []byte(`[
							   {
								  "correlation_id": "39324b8f-cc0b-439c-8ae3",
								  "original_url": "http://twokw5qvxcc8.net/rlnnq"
							   },
							   {
								  "correlation_id": "1d272046-3115-47ba-be1b",
								  "original_url": "http://id204rs9w3crhe.net"
							   }
						  ]`,
			),
			headers: map[string]string{},
			want: want{
				code:              http.StatusCreated,
				contentTypeHeader: "application/json",
				response:          "[{\"correlation_id\":\"39324b8f-cc0b-439c-8ae3\",\"short_url\":\"" + config.GetBaseURL() + "/8848d687af0fa66c0f43cc1568df1630\"},{\"correlation_id\":\"1d272046-3115-47ba-be1b\",\"short_url\":\"" + config.GetBaseURL() + "/4810e91a0c836cb8ad4121310b2d42ce\"}]",
			},
		},
		{
			name:    "API. Negative test. Add batch URLs by empty body.",
			urlPath: "/api/shorten/batch",
			method:  http.MethodPost,
			body:    []byte(``),
			headers: map[string]string{},
			want: want{
				code:              http.StatusBadRequest,
				contentTypeHeader: "application/json",
				response:          "",
			},
		},
		{
			name:    "Positive test. Add codeblog.pro site url.",
			urlPath: "/",
			method:  http.MethodPost,
			body:    []byte(`http://codeblog.pro`),
			want: want{
				code:              http.StatusCreated,
				response:          config.GetBaseURL() + `/a3035b1ecf2ee793bfe63dc13a051da6`,
				contentTypeHeader: "text/plain;charset=utf-8",
			},
		},
		{
			name:    "API. Positive test. Delete batch URLs by correct IDs.",
			urlPath: "/api/user/urls",
			method:  http.MethodDelete,
			body: []byte(`
				["a3035b1ecf2ee793bfe63dc13a051da6"]
			`),
			headers: map[string]string{
				"Content-Type": "application/json",
			},
			want: want{
				code:              http.StatusAccepted,
				contentTypeHeader: "application/json",
				response:          "",
			},
		},
		{
			name:    "API. Negative test. Delete batch URLs by empty body.",
			urlPath: "/api/user/urls",
			method:  http.MethodDelete,
			body:    []byte(``),
			headers: map[string]string{
				"Content-Type": "application/json",
			},
			want: want{
				code:              http.StatusBadRequest,
				contentTypeHeader: "application/json",
				response:          "",
			},
		},
		{
			name:    "API. Negative test. Delete batch URLs with wrong content type header.",
			urlPath: "/api/user/urls",
			method:  http.MethodDelete,
			body: []byte(`
				["a3035b1ecf2ee793bfe63dc13a051da6"]
			`),
			headers: map[string]string{
				"Content-Type": "text/plain; charset=utf-8",
			},
			want: want{
				code:              http.StatusBadRequest,
				contentTypeHeader: "text/plain; charset=utf-8",
				response:          "",
			},
		},
		{
			name:    "API. Positive test. Ping handler.",
			urlPath: "/ping",
			method:  http.MethodGet,
			body:    nil,
			headers: map[string]string{
				"Content-Type": "application/json",
			},
			want: want{
				code:              http.StatusOK,
				contentTypeHeader: "application/json",
				response:          "",
			},
		},
		{
			name:    "API. Negative test. Ping handler with wrong method.",
			urlPath: "/ping",
			method:  http.MethodPost,
			body:    nil,
			headers: map[string]string{
				"Content-Type": "application/json",
			},
			want: want{
				code:              http.StatusMethodNotAllowed,
				contentTypeHeader: "text/plain;charset=utf-8",
				response:          "",
			},
		},
		{
			name:    "API. Positive test. Get user URLs by user ID.",
			urlPath: "/api/user/urls",
			method:  http.MethodGet,
			body:    nil,
			headers: map[string]string{},
			want: want{
				code:              http.StatusOK,
				contentTypeHeader: "application/json",
				response:          "",
			},
		},
		{
			name:               "API. Negative test. Get user URLs by user ID without user token header.",
			urlPath:            "/api/user/urls",
			method:             http.MethodGet,
			body:               nil,
			headers:            map[string]string{},
			isCookiesFinalized: true,
			want: want{
				code:              http.StatusNoContent,
				contentTypeHeader: "application/json",
				response:          "",
			},
		},
	}

	var userTokenCookie http.Cookie

	for _, testData := range tests {
		if !testData.isCookiesFinalized && userTokenCookie.Value != "" {
			testData.cookies = append(testData.cookies, userTokenCookie)
		}

		response, bodyString := testRequest(t, server, testData.method, testData.urlPath, testData.body, testData.headers, testData.cookies)

		for _, cookie := range response.Cookies() {
			if cookie.Name == string(services.UserTokenName) && len(cookie.Value) > 0 {
				userTokenCookie = *cookie
			}
		}

		if response != nil {
			defer response.Body.Close()
		}

		if testData.want.code > 0 {
			assert.Equal(t, testData.want.code, response.StatusCode)
		}

		if len(testData.want.response) > 0 {
			assert.Equal(t, bodyString, testData.want.response)
		}

		if len(testData.want.contentTypeHeader) > 0 {
			assert.Equal(t, response.Header.Get("Content-Type"), testData.want.contentTypeHeader)
		}

		if len(testData.want.locationHeader) > 0 {
			assert.Equal(t, response.Header.Get("location"), testData.want.locationHeader)
		}

		if len(testData.want.contentEncoding) > 0 {
			assert.Equal(t, response.Header.Get("Content-Encoding"), testData.want.contentEncoding)
		}
	}
}

func Compress(data []byte) ([]byte, error) {
	var b bytes.Buffer
	w, err := gzip.NewWriterLevel(&b, gzip.BestCompression)

	if err != nil {
		return nil, fmt.Errorf("failed init compress writer: %v", err)
	}

	_, err = w.Write(data)

	if err != nil {
		return nil, fmt.Errorf("failed write data to compress temporary buffer: %v", err)
	}

	err = w.Close()

	if err != nil {
		return nil, fmt.Errorf("failed compress data: %v", err)
	}

	return b.Bytes(), nil
}

func getRouterForRouteTest() chi.Router {
	URLMemoryRepository := repositories.GetURLMemoryRepository()
	config := config.New()
	shorteningService := services.GetShorteningService(config)
	memoryService := &services.MemoryService{Config: config, Repository: *URLMemoryRepository, ShorteningService: *shorteningService}

	fileStorageRepository, err := repositories.GetFileStorageRepository(config)

	if err != nil {
		log.Fatalln("error creating file repository by config:" + err.Error())
	}

	errorService := services.ErrorService{}
	databaseService := database.GetDatabaseService(config)
	databaseUserRepository := repositories.GetDatabaseUserRepository(databaseService.GetDatabaseConnection())
	databaseUserService := services.GetDatabaseUserService(*databaseUserRepository)
	databaseURLRepository := tests.GetDatabaseURLRepositoryMock(databaseService)
	databaseURLService := services.GetDatabaseURLService(databaseURLRepository, *databaseUserService)
	storageService := services.GetStorageService(config, fileStorageRepository)
	encryptionService, _ := encryption.NewEncryptionService(config)
	contextStorageService := services.GetContextStorageService()
	userTokenAuthorizationService := services.GetUserTokenAuthorizationService()
	httpHandler := GetHTTPHandler(
		errorService,
		*memoryService,
		*storageService,
		encryptionService,
		*shorteningService,
		contextStorageService,
		*userTokenAuthorizationService,
		databaseService,
		*databaseURLService,
		*databaseUserService,
	)

	return httpHandler.NewRouter()
}
