// Package servers that run an HTTP server to provide endpoints for interacting with the service
package servers

import (
	"github.com/PanovAlexey/url_carver/config"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	_ "net/http/pprof"
)

type HandlerInterface interface {
	NewRouter() chi.Router
}

func RunServer(handler HandlerInterface, config config.Config) {
	router := handler.NewRouter()

	if config.IsDebug() {
		go func() {
			err := http.ListenAndServe(config.GetServerDebugAddress(), nil)

			if err != nil {
				log.Fatalln("error occurred while running http documentation server: %s", err.Error())
			}
		}()
	}

	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe(config.GetServerAddress(), router))
	log.Println("Server stopped.")
}
