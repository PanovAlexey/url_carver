package http

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (h *httpHandler) HandleDeleteBatchURLs(w http.ResponseWriter, r *http.Request) {
	userToken := h.contextStorageService.GetUserTokenFromContext(r.Context())

	if !h.userTokenAuthorizationService.IsUserTokenValid(userToken) {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	defer r.Body.Close()
	bodyJSON, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var URLsCollection []string
	err = json.Unmarshal(bodyJSON, &URLsCollection)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.databaseURLService.RemoveByShortURLSlice(URLsCollection, userToken)
	h.memoryService.DeleteURLsByShortValueSlice(URLsCollection)

	var message string

	if err == nil {
		w.WriteHeader(http.StatusAccepted)
		message = "URL batch sent for deletion"
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		message = "error while deleting batch with URLs " + err.Error()
	}

	log.Println(message)
	w.Write([]byte(message))
}
