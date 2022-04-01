package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *httpHandler) HandleGetURLsByUserToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userToken := h.contextStorageService.GetUserTokenFromContext(r.Context())

	if !h.userTokenAuthorizationService.IsUserTokenValid(userToken) {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	URLCollection := h.memoryService.GetURLsByUserToken(userToken)

	if len(URLCollection) < 1 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	URLCollectionForShowingUser := h.URLMappingService.MapURLEntityCollectionToDTO(URLCollection)
	URLCollectionJSON, err := json.Marshal(URLCollectionForShowingUser)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("get URL collection by user id:" + userToken)

	w.WriteHeader(http.StatusOK)
	w.Write(URLCollectionJSON)
}
