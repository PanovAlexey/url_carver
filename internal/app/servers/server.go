package servers

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type handlerInterface interface {
	NewRouter() chi.Router
}

type configInterface interface {
	GetServerAddress() string
}

func RunServer(handler handlerInterface, config configInterface) {
	router := handler.NewRouter()

	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe(config.GetServerAddress(), router))
	log.Println("Server stopped.")
}
