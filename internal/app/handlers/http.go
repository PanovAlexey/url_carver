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

func GetHTTPHandler(shortURLService shortURLServiceInterface) *httpHandler {
	return &httpHandler{shortURLService: shortURLService}
}

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

	longURLDto := h.shortURLService.CreateLongURLDto()
	longURLDto.SetValue(string(body))
	url := h.shortURLService.GetURLByLongURLDto(longURLDto)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(url.ShortURL))
}

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
