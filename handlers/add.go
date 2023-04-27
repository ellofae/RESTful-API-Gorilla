package handlers

import (
	"net/http"

	"github.com/ellofae/RESTful-API-Gorilla/data"
)

// swagger:route POST /products products addProducts
//
// # Adds new product in the data storage
//
// Responses:
//	200: addData
// 	500: addDataServerError

// AddProducts adds a new product to the data storage
func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("POST method")

	prodObj := r.Context().Value(MiddlewareDataKey{}).(*data.Product)

	data.AddProduct(prodObj)
}
