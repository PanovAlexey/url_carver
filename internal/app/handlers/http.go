package handlers

import (
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
)

type shortURLServiceInterface interface {
	CutAndAddEmail(email string) string
	GetEmailByKey(key string) string
	IsExistEmailByKey(key string) bool
}

type httpHandler struct {
	shortURLService shortURLServiceInterface
}

func GetHttpHandler(shortURLService shortURLServiceInterface) *httpHandler {
	return &httpHandler{shortURLService: shortURLService}
}

func (h *httpHandler) HandleGetUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain;charset=utf-8")

	urlId := chi.URLParam(r, "id")
	if len(urlId) == 0 || !h.shortURLService.IsExistEmailByKey(urlId) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add("location", h.shortURLService.GetEmailByKey(urlId))
	w.WriteHeader(http.StatusTemporaryRedirect)
	w.Write([]byte(""))
}

func (h *httpHandler) HandleAddUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain;charset=utf-8")

	defer r.Body.Close()
	longURL, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(longURL) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	key := h.shortURLService.CutAndAddEmail(string(longURL))

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(key))
}
