package handlers

import (
	"net/http"

	"github.com/ellofae/RESTful-API-Gorilla/data"
)

// swagger:route GET /products products listProducts
//
// # Lists all products from the data storage
//
// Responses:
// 	200: productsResponse
//  500: productsResponseError

// GetProducts returns the products from the data storage
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("GET method")

	lp := data.GetProducts()

	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Didn't manage to encode products data", http.StatusInternalServerError)
		return
	}
}
