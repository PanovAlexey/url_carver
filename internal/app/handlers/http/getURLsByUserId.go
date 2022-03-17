package http

import (
	"encoding/json"
	"fmt"
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

	collection = h.URLMappingService.MapURLEntityCollectionToDTO(collection)

	URLCollectionJSON, err := json.Marshal(collection.GetCollection())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("get URL collection by user id:" + userToken)

	w.WriteHeader(http.StatusOK)
	w.Write(URLCollectionJSON)
}
