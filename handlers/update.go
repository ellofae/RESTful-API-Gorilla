package handlers

import (
	"net/http"
	"strconv"

	"github.com/ellofae/RESTful-API-Gorilla/data"
	"github.com/gorilla/mux"
)

// swagger:route PUT /products/{id} products updateProducts
//
// # Updates an existing product in the data storage
//
// Responses:
//
//	200: updateData
//	400: updateDataBadRequest
//	404: updateDataNotFound

// UpdateData updates an existing product in the data storage
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
