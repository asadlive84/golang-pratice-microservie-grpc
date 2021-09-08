package handler

import (
	"log"
	"net/http"

	"github.com/grpc-microservice/data"
)

type Product struct {
	log *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{
		log: l,
	}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	p.log.Println("welcome to data")
	if r.Method == http.MethodGet {
		d := data.GetProducts()
		err := d.ToJSON(rw)
		if err != nil {
			http.Error(rw, "unable to unmarshal", http.StatusInternalServerError)
		}

	}
	if r.Method == http.MethodPost {
		prod := &data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "Not decode", http.StatusInternalServerError)
			return
		}

		data.AddProduct(prod)
	}

	http.Error(rw, "Error", http.StatusMethodNotAllowed)

}
