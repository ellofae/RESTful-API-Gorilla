// Package classification for Bakery API
//
// # Documentation for Bakery API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
// Contact: Sergei Bykovskiy<bykovskiy.sergei.dev@gmail.com>
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/ellofae/RESTful-API-Gorilla/data"
)

// ProductsResponse is a satisfied response to the call of data from the data storage
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the data storage
	// in: body
	Body []data.Product
}

// ProductsResponseError is an error response to an unsatisfied request to call data from the storage
// swagger:response productsResponseError
type productsResponseErrorWrapper struct {
	// Error: didn't manage to decide data sent to the user
	// in: body
	Body []data.Product
}

// UpdateData is a satisfied response to the call to update a product in the data storage
// swagger:response updateData
type updateDataWrapper struct {
	// A product object in the data storage has been updated
	// in: body

	Body data.Product
}

// UpdateDataBadRequstWrapper is an error response the incorrect/invalid request to update the data
// swagger:response updateDataBadRequest
type updateDataBadRequestWrapper struct {
}

// UpdateDataNotFound is an error response to the call to update data because of the non-existing object
// swagger:response updateDataNotFound
type updateDataNotFoundWrapper struct {
}

// ProductIDParameter is a required parameter to a request to update a data in the data storage
// swagger:parameters updateProducts
type productIDParameter struct {
	// The ID of the product to update in the data storage
	// in: path
	// Required: true
	ID int `json:"id"`
}

// AddData is a satisfied resposne to the request to add new data to the data storage
// swagger:response addData
type addDataWrapper struct {
	// Adds new data to the storage
	// in: body
	Body data.Product
}

// AddDataServerError is an error resposne to the internal server error while decoding data
// swagger:response addDataServerError
type addDataServerErrorWrapper struct {
	// in: body
	Body data.Product
}

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

type MiddlewareDataKey struct{}

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
