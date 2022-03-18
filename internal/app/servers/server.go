package servers

import (
	"github.com/PanovAlexey/url_carver/config"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type HandlerInterface interface {
	NewRouter() chi.Router
}

func RunServer(handler HandlerInterface, config config.Config) {
	router := handler.NewRouter()

	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe(config.GetServerAddress(), router))
	log.Println("Server stopped.")
}
