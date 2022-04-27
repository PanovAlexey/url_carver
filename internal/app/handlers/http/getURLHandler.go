package http

import (
	"github.com/PanovAlexey/url_carver/internal/app/services"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *httpHandler) HandleGetURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain;charset=utf-8")

	urlID := chi.URLParam(r, "id")

	if len(urlID) == 0 || !h.memoryService.IsExistURLEntityByShortURL(urlID) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	urlFull, err := h.memoryService.GetURLEntityByShortURL(urlID)

	if err != nil {
		errorService := services.GetErrorService()

		if errorService.IsDeleted(err) {
			w.WriteHeader(http.StatusGone)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}

		return
	}

	w.Header().Add("location", urlFull)
	w.WriteHeader(http.StatusTemporaryRedirect)
	w.Write([]byte(""))
}
