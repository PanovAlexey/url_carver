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
	GetPort() string
	GetServerAddress() string
}

func RunServer(handler handlerInterface, config configInterface) {
	router := handler.NewRouter()

	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe(getServerPort(), router))
	log.Fatal(http.ListenAndServe(getServerPort(*config), router))
	log.Fatal(http.ListenAndServe(getServerPort(config), router))
	log.Println("Server stopped.")
}

func getServerPort(config configInterface) string {
	return config.GetServerAddress() + ":" + config.GetPort()
}
