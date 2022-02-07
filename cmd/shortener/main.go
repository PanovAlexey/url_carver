package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const serverPort uint16 = 8080

type MainHandler struct {
	URL map[int64]string
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

	id, err := strconv.Atoi(queryParamArray[1])

	if id == 0 || err != nil || h.URL[int64(id)] == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add("location", h.URL[int64(id)])
	w.WriteHeader(http.StatusTemporaryRedirect)
	w.Write([]byte(""))
}

func handlePostRequest(queryParamArray []string, w http.ResponseWriter, r *http.Request, h MainHandler) {
	longUrl, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if len(queryParamArray) > 2 || queryParamArray[1] != "" || longUrl == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h.URL[int64(len(h.URL)+1)] = string(longUrl)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(getShortUrlByLongUrl(string(longUrl))))
}

func getInitialURLMap() map[int64]string {
	urlMap := map[int64]string{
		1: "https://www.yandex.ru",
		2: "https://www.google.com",
		3: "http://www.codeblog.pro",
	}

	return urlMap
}

func getServerPort() string {
	return ":" + fmt.Sprint(serverPort)
}

func getShortUrlByLongUrl(longUrl string) string {
	return "https://temp_" + fmt.Sprint(len((longUrl))) + ".com"
}
