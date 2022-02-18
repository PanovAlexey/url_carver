package servers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

const serverPort uint16 = 8080

type handlerInterface interface {
	NewRouter() chi.Router
}

func RunServer(handler handlerInterface) {
	router := handler.NewRouter()

	log.Println("Starting server...")

	log.Fatal(http.ListenAndServe(getServerPort(), router))

	log.Println("Server stopped.")
}

func getServerPort() string {
	return ":" + fmt.Sprint(serverPort)
}
