package main

import (
	"fmt"
	"net/http"
)

const serverPort uint16 = 8080

type MainHandler struct {
	Template []byte
	URL      []string
}

func (h MainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
}

func main() {
	mainHandler := MainHandler{
		URL: getInitialURLArray(),
	}
	http.ListenAndServe(getServerPort(), mainHandler)
}

func getInitialURLArray() []string {
	urlArray := []string{
		"www.yandex.ru",
		"www.google.com",
		"www.codeblog.pro",
	}

	return urlArray
}

func getServerPort() string {
	return ":" + fmt.Sprint(serverPort)
}
