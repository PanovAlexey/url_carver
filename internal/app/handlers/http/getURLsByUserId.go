package http

import (
	"encoding/json"
	"fmt"
	"github.com/PanovAlexey/url_carver/internal/app/domain/dto"
	"net/http"
)

func (h *httpHandler) HandleGetURLsByUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userToken := h.contextStorageService.GetUserIDFromContext(r.Context())

	if !h.userTokenAuthorizationService.IsUserTokenValid(userToken) {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	collection := h.memoryService.GetURLsByUserID(userToken)

	if len(collection.GetCollection()) < 1 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	collection = mapURLEntityCollectionToDTO(collection)

	URLCollectionJSON, err := json.Marshal(collection.GetCollection())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("get URL collection by user id:" + userToken)

	w.WriteHeader(http.StatusOK)
	w.Write(URLCollectionJSON)
}

func mapURLEntityCollectionToDTO(collection dto.URLCollection) dto.URLCollection {
	collectionDto := dto.GetURLCollection()

	for _, url := range collection.GetCollection() {
		collectionDto.AppendURL(dto.New(url.GetLongURL(), url.GetShortURL()))
	}

	return *collectionDto
}
