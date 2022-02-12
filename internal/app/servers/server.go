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
	HandleGetURL(w http.ResponseWriter, r *http.Request)
	HandleAddURL(w http.ResponseWriter, r *http.Request)
}

func RunServer(handler handlerInterface) {
	router := NewRouter(handler)

	log.Println("Starting server...")
	err := http.ListenAndServe(getServerPort(), router)

	log.Fatal(http.ListenAndServe(getServerPort(), router))
	if err != nil {
		log.Println(err)
	}

	log.Println("Server stopped.")
}

func NewRouter(handler handlerInterface) chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)

	router.Get("/{id}", handler.HandleGetURL)
	router.Post("/", handler.HandleAddURL)

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain;charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		w.WriteHeader(404)
	})
	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain;charset=utf-8")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.WriteHeader(405)
	})

	return router
}

func getServerPort() string {
	return ":" + fmt.Sprint(serverPort)
}
