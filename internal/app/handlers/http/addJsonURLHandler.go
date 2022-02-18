package http

import (
	"encoding/json"
	"io"
	"net/http"
)

func (h *httpHandler) HandleAddJsonURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()
	bodyJson, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	longURLDto := h.shortURLService.CreateLongURLDto()
	err = json.Unmarshal(bodyJson, &longURLDto)

	url := h.shortURLService.GetURLByLongURLDto(longURLDto)

	if err != nil || len(url.LongURL) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	shortURLJson, err := json.Marshal(h.shortURLService.GetShortURLDtoByURL(url))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(shortURLJson))
}
