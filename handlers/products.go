package handlers

import (
	"log"
	"net/http"

	"github.com/ellofae/RESTful-API-Gorilla/data"
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

func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("POST method")

	productObj := &data.Product{}

	err := productObj.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Didn't manage to decode product's data", http.StatusInternalServerError)
		return
	}

	data.AddProduct(productObj)
}

func (p *Products) updateData(id int, rw http.ResponseWriter, r *http.Request) {
	productObj := &data.Product{}

	err := productObj.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Didn't manage to decode data", http.StatusInternalServerError)
		return
	}

	err = data.UpdateData(id, productObj)
	if err == data.ErrProductNotFound {
		http.Error(rw, "The product was not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "The product was not found", http.StatusNotFound)
		return
	}
}
