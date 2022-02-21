package servers

import (
	"github.com/PanovAlexey/url_carver/config"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type handlerInterface interface {
	NewRouter() chi.Router
}

func RunServer(handler handlerInterface, config *config.Config) {
	router := handler.NewRouter()

	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe(getServerPort(), router))
	log.Fatal(http.ListenAndServe(getServerPort(*config), router))
	log.Println("Server stopped.")
}

func getServerPort(config config.Config) string {
	return ":" + config.Server.ServerPort
}
