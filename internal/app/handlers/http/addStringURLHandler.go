package http

import (
	"fmt"
	"io"
	"net/http"
)

func (h *httpHandler) HandleAddURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain;charset=utf-8")
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(body) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	url, err := h.shorteningService.GetURLEntityByLongURL(string(body))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	url.UserID = h.contextStorageService.GetUserTokenFromContext(r.Context())
	h.memoryService.SaveURL(url)
	h.storageService.SaveURL(url)
	_, err = h.databaseURLService.SaveURL(url)

	if err != nil || len(url.LongURL) == 0 {
		if h.errorService.IsKeyDuplicated(err) {
			w.WriteHeader(http.StatusConflict)
		} else {
			// database errors should be ignored
			w.WriteHeader(http.StatusCreated)
		}
	} else {
		w.WriteHeader(http.StatusCreated)
	}

	shortURLJSON := h.memoryService.GetShortURLDtoByURL(url)

	fmt.Println("URL " + url.LongURL + " added by " + url.ShortURL)
	w.Write([]byte(shortURLJSON.Value))
}
