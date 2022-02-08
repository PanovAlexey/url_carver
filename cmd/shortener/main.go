package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const serverPort uint16 = 8080

type MainHandler struct {
	URL map[string]string
}

func (h MainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryParamArray := strings.Split(r.URL.Path, "/")

	switch r.Method {
	case http.MethodPost:
		handlePostRequest(queryParamArray, w, r, h)
	case http.MethodGet:
		handleGetRequest(queryParamArray, w, r, h)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(""))
	}
}

func main() {
	mainHandler := MainHandler{
		URL: getInitialURLMap(),
	}
	http.ListenAndServe(getServerPort(), mainHandler)
}

func handleGetRequest(queryParamArray []string, w http.ResponseWriter, r *http.Request, h MainHandler) {
	if len(queryParamArray) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := queryParamArray[1]

	if len(id) == 0 || h.URL[(id)] == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add("location", h.URL[id])
	w.WriteHeader(http.StatusTemporaryRedirect)
	w.Write([]byte(""))
}

func handlePostRequest(queryParamArray []string, w http.ResponseWriter, r *http.Request, h MainHandler) {
	longURL, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if len(queryParamArray) > 2 || queryParamArray[1] != "" || len(longURL) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	shortURLCode := getShortURLCode(string(longURL))

	if len(h.URL[shortURLCode]) == 0 {
		h.URL[shortURLCode] = string(longURL)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(getShortURLByLongURL(string(longURL))))
}

func getInitialURLMap() map[string]string {
	urlMap := map[string]string{
		"yandex": "https://www.yandex.ru",
		"google": "https://www.google.com",
		"meta":   "https://about.facebook.com/meta/",
	}

	return urlMap
}

func getServerPort() string {
	return ":" + fmt.Sprint(serverPort)
}

func getShortURLCode(longURL string) string {
	return fmt.Sprint(len(longURL) + 1)
}

func getShortURLByLongURL(longURL string) string {
	return "https://clck.ru/" + getShortURLCode(longURL)
}
