package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type StorageInterface interface {
	AddEmail(key string, email string) bool
	GetEmailByKey(key string) string
	IsExistEmailByKey(key string) bool
}

type httpHandler struct {
	storage StorageInterface
}

func GetHttpHandler(storage StorageInterface) *httpHandler {
	return &httpHandler{storage: storage}
}

func (h *httpHandler) HandleGetUrl(w http.ResponseWriter, r *http.Request) {
	queryParamArray := strings.Split(r.URL.Path, "/")

	if len(queryParamArray) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := queryParamArray[1]
	if len(id) == 0 || !h.storage.IsExistEmailByKey(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add("location", h.storage.GetEmailByKey(id))
	w.WriteHeader(http.StatusTemporaryRedirect)
	w.Write([]byte(""))
}

func (h *httpHandler) HandleAddUrl(w http.ResponseWriter, r *http.Request) {
	queryParamArray := strings.Split(r.URL.Path, "/")

	defer r.Body.Close()
	longURL, err := io.ReadAll(r.Body)

	w.Header().Set("Content-Type", "text/plain;charset=utf-8")

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if len(queryParamArray) > 2 || queryParamArray[1] != "" || len(longURL) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	shortURLCode := getShortURLCode(string(longURL))

	if !h.storage.IsExistEmailByKey(shortURLCode) {
		h.storage.AddEmail(shortURLCode, string(longURL))
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(getShortURLByLongURL(string(longURL))))
}

func getShortURLCode(longURL string) string {
	return fmt.Sprint(len(longURL) + 1)
}

func getShortURLByLongURL(longURL string) string {
	return "http://localhost:8080/" + getShortURLCode(longURL)
}
