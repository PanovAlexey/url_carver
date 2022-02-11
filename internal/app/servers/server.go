package servers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

const serverPort uint16 = 8080

type handlerInterface interface {
	HandleGetUrl(w http.ResponseWriter, r *http.Request)
	HandleAddUrl(w http.ResponseWriter, r *http.Request)
}

func RunServer(handler handlerInterface) {
	router := NewRouter(handler)

	log.Println("Starting server...")
	err := http.ListenAndServe(getServerPort(), router)

	if err != nil {
		log.Println(err)
	}

	log.Println("Server stopped.")
}

func NewRouter(handler handlerInterface) chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)

	router.Get("/{id}", handler.HandleGetUrl)
	router.Post("/", handler.HandleAddUrl)

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain;charset=utf-8")
		w.WriteHeader(404)
	})
	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain;charset=utf-8")
		w.WriteHeader(405)
	})

	return router
}

func getServerPort() string {
	return ":" + fmt.Sprint(serverPort)
}