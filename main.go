package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/grpc-microservice/handler"
)

func main() {

	l := log.New(os.Stdout, "GRPC practise", log.LstdFlags)

	p := handler.NewProduct(l)

	server := mux.NewRouter()

	pr := handler.NewProduct(l)

	getRouter := server.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", pr.GetProduct)

	putRouter := server.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", pr.UpdateProduct)
	putRouter.Use(p.MiddlewareValidateProduct)

	postRouter := server.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/add", pr.AddProduct)
	postRouter.Use(p.MiddlewareValidateProduct)

	//http.Handle("/", handler.LoggingMiddleware(server))

	createServer := http.Server{
		Addr:              ":9000",
		Handler:           server,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
	}

	createServer.ListenAndServe()
}
