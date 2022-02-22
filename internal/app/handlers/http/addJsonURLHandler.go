package http

import (
	"encoding/json"
	"fmt"
	"io"
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

	longURLDto := h.URLMemoryService.CreateLongURLDto()
	err = json.Unmarshal(bodyJSON, &longURLDto)

	url := h.URLMemoryService.GetURLByLongURLDto(longURLDto)
	h.URLStorageService.SaveURL(url)

	if err != nil || len(url.LongURL) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	shortURLJson, err := json.Marshal(h.URLMemoryService.GetShortURLDtoByURL(url))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("URL " + url.LongURL + " added by " + url.ShortURL)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(shortURLJson))
}
