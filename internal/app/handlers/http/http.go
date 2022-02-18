package http

import (
	"github.com/PanovAlexey/url_carver/internal/app/domain/dto"
	"github.com/PanovAlexey/url_carver/internal/app/domain/entity/url"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type shortURLServiceInterface interface {
	GetEmailByKey(key string) string
	IsExistEmailByKey(key string) bool
	CreateLongURLDto() dto.LongURL
	GetURLByLongURLDto(dto.LongURL) url.URL
	GetShortURLDtoByURL(url url.URL) dto.ShortURL
}

type httpHandler struct {
	shortURLService shortURLServiceInterface
}

func GetHTTPHandler(shortURLService shortURLServiceInterface) *httpHandler {
	return &httpHandler{shortURLService: shortURLService}
}

func (h *httpHandler) NewRouter() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)

	router.Get("/{id}", h.HandleGetURL)
	router.Post("/", h.HandleAddURL)

	router.Post("/api/shorten", h.HandleAddJsonURL)

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
