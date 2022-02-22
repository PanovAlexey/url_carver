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

	longURLDto := h.URLMemoryService.CreateLongURLDto()
	longURLDto.SetValue(string(body))
	url := h.URLMemoryService.GetURLByLongURLDto(longURLDto)
	h.URLStorageService.SaveURL(url)

	shortURLJson := h.URLMemoryService.GetShortURLDtoByURL(url)

	fmt.Println("URL " + url.LongURL + " added by " + url.ShortURL)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(shortURLJson.Value))
}
