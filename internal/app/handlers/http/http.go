package http

import (
	"context"
	"github.com/PanovAlexey/url_carver/internal/app/domain/dto"
	"github.com/PanovAlexey/url_carver/internal/app/domain/entity/url"
	internalMiddleware "github.com/PanovAlexey/url_carver/internal/app/handlers/http/middleware"
	"github.com/PanovAlexey/url_carver/internal/app/services/database"
	"github.com/PanovAlexey/url_carver/internal/app/services/encryption"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type memoryServiceInterface interface {
	GetURLByKey(key string) string
	IsExistURLByKey(key string) bool
	CreateLongURLDto() dto.LongURL
	GetShortURLDtoByURL(url url.URL) dto.ShortURL
	SaveURL(url url.URL) bool
	GetURLsByUserID(userID string) dto.URLCollection
}

type storageServiceInterface interface {
	SaveURL(url url.URL)
}

type shorteningServiceInterface interface {
	GetURLEntityByLongURL(longURL string) (url.URL, error)
}

type contextStorageServiceInterface interface {
	GetUserIDFromContext(ctx context.Context) string
}

type userTokenAuthorizationServiceInterface interface {
	IsUserTokenValid(userToken string) bool
}

type URLMappingServiceInterface interface {
	MapURLEntityCollectionToDTO(collection dto.URLCollection) dto.URLCollection
}

type httpHandler struct {
	memoryService                 memoryServiceInterface
	storageService                storageServiceInterface
	encryptionService             encryption.EncryptorInterface
	shorteningService             shorteningServiceInterface
	contextStorageService         contextStorageServiceInterface
	userTokenAuthorizationService userTokenAuthorizationServiceInterface
	URLMappingService             URLMappingServiceInterface
	databaseService               database.DatabaseInterface
}

func GetHTTPHandler(
	memoryService memoryServiceInterface,
	storageService storageServiceInterface,
	encryptionService encryption.EncryptorInterface,
	shorteningService shorteningServiceInterface,
	contextStorageService contextStorageServiceInterface,
	userTokenAuthorizationService userTokenAuthorizationServiceInterface,
	URLMappingService URLMappingServiceInterface,
	databaseService database.DatabaseInterface,
) *httpHandler {
	return &httpHandler{
		memoryService:                 memoryService,
		storageService:                storageService,
		encryptionService:             encryptionService,
		shorteningService:             shorteningService,
		contextStorageService:         contextStorageService,
		userTokenAuthorizationService: userTokenAuthorizationService,
		URLMappingService:             URLMappingService,
		databaseService:               databaseService,
	}
}

func (h *httpHandler) NewRouter() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(internalMiddleware.Authorization(h.encryptionService))
	router.Use(internalMiddleware.GZip)

	router.Get("/ping", h.HandlePingDatabase)
	router.Get("/{id}", h.HandleGetURL)
	router.Post("/", h.HandleAddURL)

	router.Post("/api/shorten", h.HandleAddURLByJSON)

	router.Get("/api/user/urls", h.HandleGetURLsByUserID)

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
