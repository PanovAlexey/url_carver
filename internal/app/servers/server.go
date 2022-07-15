// Package servers that run an HTTP server to provide endpoints for interacting with the service
package servers

import (
	"context"
	"github.com/PanovAlexey/url_carver/config"
	grpcServices "github.com/PanovAlexey/url_carver/internal/app/services/grpc"
	"github.com/PanovAlexey/url_carver/internal/app/services/grpc/interceptors"
	pb "github.com/PanovAlexey/url_carver/pkg/shortener_grpc"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
	"log"
	"net"
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

type mainServer struct {
	httpServer *http.Server
}

type grpcServer struct {
	server *grpc.Server
}

func RunServer(handler HandlerInterface, grpcService grpcServices.ShortenerService, config config.Config) {
	go runDocumentationServer(config)
	go runURLCarverGRpcServer(grpcService, config)
	runURLCarverServer(handler, config)
}

func runDocumentationServer(config config.Config) {
	if config.IsDebug() {
		log.Println("Documentation starting server...")
		err := http.ListenAndServe(config.GetServerDebugAddress(), nil)

		if err != nil {
			log.Printf("error occurred while running http documentation server: %s", err.Error())
		}

		log.Println("Documentation server stopped.")
	}
}

func runURLCarverServer(handler HandlerInterface, config config.Config) {
	log.Println("UrlCarver starting server...")
	router := handler.NewRouter()
	srv := new(mainServer)
	srv.httpServer = &http.Server{
		Addr:    config.GetServerAddress(),
		Handler: router,
	}
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

func runURLCarverGRpcServer(grpcService grpcServices.ShortenerService, config config.Config) {
	log.Println("UrlCarver gRPC starting server...")
	listen, err := net.Listen("tcp", ":"+config.GetGRpcServerPort())

	if err != nil {
		log.Fatal(err)
	}

	grpcServer := new(grpcServer)
	grpcServer.server = grpc.NewServer(grpc.UnaryInterceptor(interceptors.AuthorizationByToken))
	pb.RegisterShortenerServer(grpcServer.server, &grpcService)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := grpcServer.server.Serve(listen); err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("Signal detected: ", <-sigs)

	grpcServer.server.GracefulStop()
	log.Println("UrlCarver gRPC server is shutdowning...")

	if err != nil {
		log.Println(err)
	}

	log.Println("UrlCarver gRPC server stopped.")

}
