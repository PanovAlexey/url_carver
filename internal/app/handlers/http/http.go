package http

import (
	"github.com/PanovAlexey/url_carver/internal/app/domain/dto"
	"github.com/PanovAlexey/url_carver/internal/app/domain/entity/url"
	internalMiddleware "github.com/PanovAlexey/url_carver/internal/app/handlers/http/middleware"
	"github.com/PanovAlexey/url_carver/internal/app/services/encryption"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type memoryServiceInterface interface {
	GetURLByKey(key string) string
	IsExistURLByKey(key string) bool
	CreateLongURLDto() dto.LongURL
	GetURLByLongURLDto(dto.LongURL) url.URL
	GetShortURLDtoByURL(url url.URL) dto.ShortURL
}

type storageServiceInterface interface {
	SaveURL(url url.URL)
}

type httpHandler struct {
	memoryService     memoryServiceInterface
	storageService    storageServiceInterface
	encryptionService encryption.EncryptorInterface
}

func GetHTTPHandler(
	memoryService memoryServiceInterface,
	storageService storageServiceInterface,
	encryptionService encryption.EncryptorInterface,
) *httpHandler {
	return &httpHandler{
		memoryService:     memoryService,
		storageService:    storageService,
		encryptionService: encryptionService,
	}
}

func (h *httpHandler) NewRouter() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(internalMiddleware.Authorization(h.encryptionService))
	router.Use(internalMiddleware.GZip)

	router.Get("/{id}", h.HandleGetURL)
	router.Post("/", h.HandleAddURL)

	router.Post("/api/shorten", h.HandleAddURLByJSON)

	router.Get("/api/user/urls", h.HandleGetURLsByUserId)

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
