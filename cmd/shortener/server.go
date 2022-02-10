package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

const serverPort uint16 = 8080

func runServer() {
	InitialURLMap()

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/{id}", handleGetUrl)
	router.Post("/", handleAddUrl)
	log.Println("Starting server...")
	http.ListenAndServe(getServerPort(), router)
	log.Println("Server stopped.")
}

func getServerPort() string {
	return ":" + fmt.Sprint(serverPort)
}
