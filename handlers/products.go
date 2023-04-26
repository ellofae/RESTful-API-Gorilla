package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/ellofae/RESTful-API-Gorilla/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("GET method")

	lp := data.GetProducts()

	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Didn't manage to encode products data", http.StatusInternalServerError)
		return
	}
}

type MiddlewareDataKey struct{}

func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("POST method")

	prodObj := r.Context().Value(MiddlewareDataKey{}).(*data.Product)

	data.AddProduct(prodObj)
}

func (p *Products) UpdateData(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("PUT method")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		p.l.Println("Bad Request Error: didn't manage to covnert string to int")
		http.Error(rw, "Incorrect URI", http.StatusBadRequest)
		return
	}

	prodObj := r.Context().Value(MiddlewareDataKey{}).(*data.Product)

	err = data.UpdateData(id, prodObj)
	if err == data.ErrProductNotFound {
		http.Error(rw, "The product was not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "The product was not found", http.StatusNotFound)
		return
	}
}

func (p *Products) MiddlewareValidationForDatatransfer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		productObj := &data.Product{}

		err := productObj.FromJSON(r.Body)
		if err != nil {
			p.l.Println("Internal Server Error: didn't manage to unmarshall data")
			http.Error(rw, "Didn't manage to decode product's data", http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(r.Context(), MiddlewareDataKey{}, productObj)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}
