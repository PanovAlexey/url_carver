// Package servers that run an HTTP server to provide endpoints for interacting with the service
package servers

import (
	"context"
	"github.com/PanovAlexey/url_carver/config"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type HandlerInterface interface {
	NewRouter() chi.Router
}

type server struct {
	httpServer *http.Server
}

func RunServer(handler HandlerInterface, config config.Config) {
	runDocumentationServer(config)
	runUrlCarverServer(handler, config)
}

func runDocumentationServer(config config.Config) {
	if config.IsDebug() {
		go func() {
			log.Println("Starting documentation server...")
			err := http.ListenAndServe(config.GetServerDebugAddress(), nil)

			if err != nil {
				log.Printf("error occurred while running http documentation server: %s", err.Error())
			}

			log.Println("Documentation server stopped.")
		}()
	}
}

func runUrlCarverServer(handler HandlerInterface, config config.Config) {
	log.Println("UrlCarver starting server...")
	router := handler.NewRouter()
	srv := newServer(config.GetServerAddress(), router)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if config.IsEnableHTTPS() {
			log.Println(
				srv.httpServer.ListenAndServeTLS("./config/ssl/cert.pem", "./config/ssl/key.pem"),
			)
		} else {
			log.Println(srv.httpServer.ListenAndServe())
		}
	}()

	log.Println("Signal detected: ", <-sigs)

	ctx, cancel := context.WithTimeout(
		context.Background(), time.Duration(config.Server.TimeoutShutdownInSeconds)*time.Second,
	)
	defer cancel()

	err := srv.httpServer.Shutdown(ctx)
	log.Println("UrlCarver server is shutdowning...")

	if err != nil {
		log.Println(err)
	}

	log.Println("UrlCarver server stopped.")
}

func newServer(address string, handler http.Handler) server {
	srv := new(server)
	srv.httpServer = &http.Server{
		Addr:    address,
		Handler: handler,
	}

	return *srv
}
