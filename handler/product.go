package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/grpc-microservice/data"
)

type Product struct {
	log *log.Logger
}
type KeyProduct struct{}

func NewProduct(l *log.Logger) *Product {
	return &Product{
		log: l,
	}
}

//Get product
func (p *Product) GetProduct(rw http.ResponseWriter, r *http.Request) {
	d := data.GetProducts()
	err := d.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable to unmarshal", http.StatusInternalServerError)
	}
}

//update product
func (p *Product) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	p.log.Println("Method PUT")
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		p.log.Println("Id parse error")
		http.Error(rw, "ID error", http.StatusInternalServerError)
	}
	id32 := int32(id)
	p.updateProduct(rw, r, id32)
}

//add product
func (p *Product) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.log.Println("POST method")
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	// prod := &data.Product{}
	// err := prod.FromJSON(r.Body)
	// if err != nil {
	// 	p.log.Println("Decode error")
	// 	http.Error(rw, "Err: Not decode", http.StatusInternalServerError)
	// 	return
	// }
	// p.log.Printf("%+v", prod)
	data.AddProduct(&prod)
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	p.log.Println("welcome to data")

	rw.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Product) updateProduct(rw http.ResponseWriter, r *http.Request, id int32) {
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		p.log.Println("update decode error")
		http.Error(rw, "error", http.StatusBadRequest)
	}
	err = data.UpdateProduct(int32(id), prod)
	if err == data.ErrorProductNotFound {
		http.Error(rw, "Error product not found", http.StatusInternalServerError)
	}
	if err != nil {
		http.Error(rw, "update error", http.StatusBadRequest)
	}
}

func (p Product) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.log.Println("[ERROR] deserializing product", err)
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}

		err = prod.Validate()
		if err != nil {
			p.log.Println("[ERROR] validating product", err)
			http.Error(rw, fmt.Sprintf("Error validating product %+v", err), http.StatusBadRequest)
			return
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
