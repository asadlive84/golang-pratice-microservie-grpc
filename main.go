package main

import (
	"log"
	"net/http"
	"os"

	"github.com/grpc-microservice/handler"
)

func main() {

	l := log.New(os.Stdout, "GRPC practise", log.LstdFlags)

	server := http.NewServeMux()

	//wc := welcome.NewWelcome(l)
	pr := handler.NewProduct(l)

	//server.Handle("/", wc)
	server.Handle("/", pr)

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
