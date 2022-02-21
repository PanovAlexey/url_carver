package http

import (
	"github.com/PanovAlexey/url_carver/internal/app/domain/dto"
	"github.com/PanovAlexey/url_carver/internal/app/domain/entity/url"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type URLMemoryServiceInterface interface {
	GetURLByKey(key string) string
	IsExistURLByKey(key string) bool
	CreateLongURLDto() dto.LongURL
	GetURLByLongURLDto(dto.LongURL) url.URL
	GetShortURLDtoByURL(url url.URL) dto.ShortURL
}

type URLStorageServiceInterface interface {
	SaveURL(url url.URL)
}

type httpHandler struct {
	URLMemoryService  URLMemoryServiceInterface
	URLStorageService URLStorageServiceInterface
}

func GetHTTPHandler(
	URLMemoryService URLMemoryServiceInterface,
	URLStorageService URLStorageServiceInterface,
) *httpHandler {
	return &httpHandler{URLMemoryService: URLMemoryService, URLStorageService: URLStorageService}
}

func (h *httpHandler) NewRouter() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)

	router.Get("/{id}", h.HandleGetURL)
	router.Post("/", h.HandleAddURL)

	router.Post("/api/shorten", h.HandleAddURLByJSON)

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain;charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
	})
	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain;charset=utf-8")
		w.WriteHeader(http.StatusMethodNotAllowed)
	})

	return router
}
