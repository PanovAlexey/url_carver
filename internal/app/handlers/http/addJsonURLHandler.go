package http

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (h *httpHandler) HandleAddURLByJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()
	bodyJSON, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	longURLDto := h.memoryService.CreateLongURLDto()
	err = json.Unmarshal(bodyJSON, &longURLDto)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	url, err := h.shorteningService.GetURLEntityByLongURL(longURLDto.Value)
	url.SetUserToken(h.contextStorageService.GetUserTokenFromContext(r.Context()))

	if h.memoryService.IsExistURLByKey(url.GetShortURL()) {
		w.WriteHeader(http.StatusConflict)
	} else {
		h.memoryService.SaveURL(url)
		h.storageService.SaveURL(url)
		_, err = h.databaseURLService.SaveURL(url)

		if err != nil {
			log.Println("failed to save url to database")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusCreated)
		}
	}

	if err != nil || len(url.LongURL) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	shortURL := h.memoryService.GetShortURLDtoByURL(url)
	shortURLJSON, err := json.Marshal(shortURL)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("URL " + url.LongURL + " added by " + url.ShortURL)

	w.Write(shortURLJSON)
}
