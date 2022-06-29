package http

import (
	"encoding/json"
	"fmt"
	"github.com/PanovAlexey/url_carver/internal/app/domain/dto"
	"github.com/PanovAlexey/url_carver/internal/app/services"
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

	longURLDto := dto.GetLongURLByValue("")
	err = json.Unmarshal(bodyJSON, &longURLDto)

	if err != nil || len(longURLDto.Value) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	url, err := h.shorteningService.GetURLEntityByLongURL(longURLDto.Value)

	if err != nil || len(url.LongURL) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	url.UserID = h.contextStorageService.GetUserTokenFromContext(r.Context())
	h.memoryService.SaveURL(url)
	h.storageService.SaveURL(url)
	_, err = h.databaseURLService.SaveURL(url)

	if err != nil {
		errorService := services.GetErrorService()

		if errorService.IsKeyDuplicated(err) {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusCreated)
			// database errors should be ignored
			// w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusCreated)
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
