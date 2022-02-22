package http

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *httpHandler) HandleGetURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain;charset=utf-8")

	urlID := chi.URLParam(r, "id")
	if len(urlID) == 0 || !h.shortURLService.IsExistEmailByKey(urlID) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add("location", h.shortURLService.GetEmailByKey(urlID))
	w.WriteHeader(http.StatusTemporaryRedirect)
	w.Write([]byte(""))
}