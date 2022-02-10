package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

const serverPort uint16 = 8080

type HandlerInterface interface {
	HandleGetUrl(w http.ResponseWriter, r *http.Request)
	HandleAddUrl(w http.ResponseWriter, r *http.Request)
}

func runServer(handler HandlerInterface) {
	InitialURLMap()

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/{id}", handler.HandleGetUrl)
	router.Post("/", handler.HandleAddUrl)

	log.Println("Starting server...")
	http.ListenAndServe(getServerPort(), router)
	log.Println("Server stopped.")
}

func getServerPort() string {
	return ":" + fmt.Sprint(serverPort)
}
