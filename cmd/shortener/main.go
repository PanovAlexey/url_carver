package main

import "github.com/PanovAlexey/url_carver/internal/app/services"

func main() {
	storage := &GlobalURLs
	shortURLService := services.GetShortURLService(storage)
	httpHandler := GetHttpHandler(shortURLService)
	runServer(httpHandler)
}
