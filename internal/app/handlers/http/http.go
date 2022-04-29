package http

import (
	"context"
	"github.com/PanovAlexey/url_carver/internal/app/domain/dto"
	"github.com/PanovAlexey/url_carver/internal/app/domain/entity/url"
	"github.com/PanovAlexey/url_carver/internal/app/domain/entity/user"
	internalMiddleware "github.com/PanovAlexey/url_carver/internal/app/handlers/http/middleware"
	"github.com/PanovAlexey/url_carver/internal/app/services/database"
	"github.com/PanovAlexey/url_carver/internal/app/services/encryption"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type memoryService interface {
	GetURLEntityByShortURL(shortURL string) (string, error)
	IsExistURLEntityByShortURL(shortURL string) bool
	CreateLongURLDto() dto.LongURL
	GetShortURLDtoByURL(url url.URL) dto.ShortURL
	SaveURL(url url.URL) bool
	GetURLsByUserToken(userToken string) []url.URL
	DeleteURLsByShortValueSlice([]string)
}

type storageServiceInterface interface {
	SaveURL(url url.URL)
}

type shorteningServiceInterface interface {
	GetURLEntityByLongURL(longURL string) (url.URL, error)
	GetShortURLWithDomain(shortURLCode string) (string, error)
}

type contextStorageServiceInterface interface {
	GetUserTokenFromContext(ctx context.Context) string
}

type userTokenAuthorizationServiceInterface interface {
	IsUserTokenValid(userToken string) bool
}

type URLMappingServiceInterface interface {
	MapURLEntityCollectionToDTO(collection []url.URL) []dto.URLForShowingUser
}

type DatabaseURLServiceInterface interface {
	SaveURL(url url.URL) (int, error)
	SaveBatchURLs(collection []url.URL)
	RemoveByShortURLSlice(URLSlice []string, userToken string) error
}
type DatabaseUserServiceInterface interface {
	SaveUser(user user.User) (int, error)
}

type httpHandler struct {
	memoryService                 memoryService
	storageService                storageServiceInterface
	encryptionService             encryption.EncryptorInterface
	shorteningService             shorteningServiceInterface
	contextStorageService         contextStorageServiceInterface
	userTokenAuthorizationService userTokenAuthorizationServiceInterface
	URLMappingService             URLMappingServiceInterface
	databaseService               database.DatabaseInterface
	databaseURLService            DatabaseURLServiceInterface
	databaseUserService           DatabaseUserServiceInterface
}

func GetHTTPHandler(
	memoryService memoryService,
	storageService storageServiceInterface,
	encryptionService encryption.EncryptorInterface,
	shorteningService shorteningServiceInterface,
	contextStorageService contextStorageServiceInterface,
	userTokenAuthorizationService userTokenAuthorizationServiceInterface,
	URLMappingService URLMappingServiceInterface,
	databaseService database.DatabaseInterface,
	databaseURLService DatabaseURLServiceInterface,
	databaseUserService DatabaseUserServiceInterface,
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
		databaseURLService:            databaseURLService,
		databaseUserService:           databaseUserService,
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

	router.With(internalMiddleware.JSON).Post("/api/shorten", h.HandleAddURLByJSON)

	router.Get("/api/user/urls", h.HandleGetURLsByUserToken)
	router.Post("/api/shorten/batch", h.HandleAddBatchURLs)

	router.With(internalMiddleware.JSON).Delete("/api/user/urls", h.HandleDeleteBatchURLs)

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
