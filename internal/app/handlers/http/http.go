package http

import (
	internalMiddleware "github.com/PanovAlexey/url_carver/internal/app/handlers/http/middleware"
	"github.com/PanovAlexey/url_carver/internal/app/services"
	"github.com/PanovAlexey/url_carver/internal/app/services/database"
	"github.com/PanovAlexey/url_carver/internal/app/services/encryption"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type httpHandler struct {
	errorService                  services.ErrorService
	memoryService                 services.MemoryService
	storageService                services.StorageService
	encryptionService             encryption.EncryptorInterface
	shorteningService             services.ShorteningService
	contextStorageService         services.ContextStorageService
	userTokenAuthorizationService services.UserTokenAuthorizationService
	URLMappingService             services.MappingService
	databaseService               database.DatabaseInterface
	databaseURLService            services.DatabaseURLService
	databaseUserService           services.DatabaseUserService
}

func GetHTTPHandler(
	errorService services.ErrorService,
	memoryService services.MemoryService,
	storageService services.StorageService,
	encryptionService encryption.EncryptorInterface,
	shorteningService services.ShorteningService,
	contextStorageService services.ContextStorageService,
	userTokenAuthorizationService services.UserTokenAuthorizationService,
	databaseService database.DatabaseInterface,
	databaseURLService services.DatabaseURLService,
	databaseUserService services.DatabaseUserService,
) *httpHandler {
	return &httpHandler{
		errorService:                  errorService,
		memoryService:                 memoryService,
		storageService:                storageService,
		encryptionService:             encryptionService,
		shorteningService:             shorteningService,
		contextStorageService:         contextStorageService,
		userTokenAuthorizationService: userTokenAuthorizationService,
		URLMappingService:             services.MappingService{},
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

	router.Post("/api/shorten", h.HandleAddURLByJSON)

	router.Get("/api/user/urls", h.HandleGetURLsByUserToken)

	router.Post("/api/shorten/batch", h.HandleAddBatchURLs)

	router.Get("/api/internal/stats", h.HandleGetAppStats)

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
