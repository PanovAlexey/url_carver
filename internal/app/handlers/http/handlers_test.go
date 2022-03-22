package http

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/PanovAlexey/url_carver/config"
	"github.com/PanovAlexey/url_carver/internal/app/domain/dto"
	"github.com/PanovAlexey/url_carver/internal/app/repositories"
	"github.com/PanovAlexey/url_carver/internal/app/services"
	"github.com/PanovAlexey/url_carver/internal/app/services/database"
	"github.com/PanovAlexey/url_carver/internal/app/services/encryption"
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
	t *testing.T, ts *httptest.Server, method, path string, body []byte, headers map[string]string,
) (*http.Response, string) {
	bodyIoReader := bytes.NewBuffer(body)
	req, err := http.NewRequest(method, ts.URL+path, bodyIoReader)
	require.NoError(t, err)

	for key, value := range headers {
		req.Header.Set(key, value)
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
		name    string
		urlPath string
		method  string
		headers map[string]string
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
		{
			name:    "Positive test. Get short url for 3dnews site after it was added with GZIP compression.",
			urlPath: "/19",
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
				response:          "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\xca())\xb0\xd2\xd7\xcf\xc9ON\xcc\xc9\xc8/.\xb1\xb20\xb00\xd07\xb4\x00\x04\x00\x00\xff\xff\xc6@\x82\xbe\x18\x00\x00\x00",
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
				response:          "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\xaaV*J-.\xcd)Q\xb2R\xca())\xb0\xd2\xd7\xcf\xc9ON\xcc\xc9\xc8/.\xb1\xb20\xb00\xd07\xb4T\xaa\x05\x04\x00\x00\xff\xff\x00V\x8b\xae%\x00\x00\x00",
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
				response:          "[{\"correlation_id\":\"39324b8f-cc0b-439c-8ae3\",\"short_url\":\"" + config.GetBaseURL() + "/30\"},{\"correlation_id\":\"1d272046-3115-47ba-be1b\",\"short_url\":\"" + config.GetBaseURL() + "/26\"}]",
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
	}

	for _, testData := range tests {
		response, bodyString := testRequest(t, server, testData.method, testData.urlPath, testData.body, testData.headers)

		if response != nil {
			defer response.Body.Close()
		}

		assert.Equal(t, testData.want.code, response.StatusCode)
		assert.Equal(t, bodyString, testData.want.response)
		assert.Equal(t, response.Header.Get("Content-Type"), testData.want.contentTypeHeader)
		assert.Equal(t, response.Header.Get("location"), testData.want.locationHeader)
		assert.Equal(t, response.Header.Get("Content-Encoding"), testData.want.contentEncoding)
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
	memoryService := services.GetMemoryService(config, URLMemoryRepository, shorteningService)
	fileStorageRepository, err := repositories.GetFileStorageRepository(config)

	if err != nil {
		log.Fatalln("error creating file repository by config:" + err.Error())
	}

	databaseService := database.GetDatabaseService(config)
	databaseUserRepository := repositories.GetDatabaseUserRepository(databaseService)
	databaseUserService := services.GetDatabaseUserService(databaseUserRepository)
	databaseURLRepository := getDatabaseURLRepositoryMock()
	databaseURLService := services.GetDatabaseURLService(databaseURLRepository, *databaseUserService)
	storageService := services.GetStorageService(config, fileStorageRepository)
	encryptionService, _ := encryption.NewEncryptionService(config)
	contextStorageService := services.GetContextStorageService()
	userTokenAuthorizationService := services.GetUserTokenAuthorizationService()
	URLMappingService := services.GetURLMappingService()
	httpHandler := GetHTTPHandler(
		memoryService,
		storageService,
		encryptionService,
		shorteningService,
		contextStorageService,
		userTokenAuthorizationService,
		URLMappingService,
		databaseService,
		databaseURLService,
		databaseUserService,
	)

	return httpHandler.NewRouter()
}

// Mock dependencies
type databaseURLRepositoryMock struct {
}

func (repository databaseURLRepositoryMock) SaveURL(url dto.DatabaseURL) (int, error) {
	return 777, nil
}

func (repository databaseURLRepositoryMock) GetList() ([]dto.DatabaseURL, error) {
	return []dto.DatabaseURL{}, nil
}

func (repository databaseURLRepositoryMock) SaveBatchURLs(collection []dto.DatabaseURL) error {
	return nil
}

func getDatabaseURLRepositoryMock() databaseURLRepositoryMock {
	return databaseURLRepositoryMock{}
}
