package main

import (
	"fmt"
	"net/http"
)

const serverPort uint16 = 8080

func runServer() {
	mainHandler := MainHandler{
		URL: getInitialURLMap(),
	}
	http.ListenAndServe(getServerPort(), mainHandler)
}

func getServerPort() string {
	return ":" + fmt.Sprint(serverPort)
}
