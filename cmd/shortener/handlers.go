package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type HttpHandler struct {
}

func GetHttpHandler() *HttpHandler {
	return &HttpHandler{}
}

func (h *HttpHandler) HandleGetUrl(w http.ResponseWriter, r *http.Request) {
	queryParamArray := strings.Split(r.URL.Path, "/")

	if len(queryParamArray) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := queryParamArray[1]
	if len(id) == 0 || !globalURLs.IsExistEmailByKey(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add("location", globalURLs.GetEmailByKey(id))
	w.WriteHeader(http.StatusTemporaryRedirect)
	w.Write([]byte(""))
}

func (h *HttpHandler) HandleAddUrl(w http.ResponseWriter, r *http.Request) {
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

	if !globalURLs.IsExistEmailByKey(shortURLCode) {
		globalURLs.AddEmail(shortURLCode, string(longURL))
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
