package main

import (
	"github.com/PanovAlexey/url_carver/internal/app/handlers"
	"github.com/PanovAlexey/url_carver/internal/app/repositories"
	"github.com/PanovAlexey/url_carver/internal/app/servers"
	"github.com/PanovAlexey/url_carver/internal/app/services"
)

func main() {
	emailRepository := repositories.GetEmailRepository()
	shortURLService := services.GetShortURLService(emailRepository)
	httpHandler := handlers.GetHttpHandler(shortURLService)
	servers.RunServer(httpHandler)
}
