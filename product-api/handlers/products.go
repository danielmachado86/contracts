// Product API
//
// Documentation for Product API
// Schemes: http
// BasePath: /
// Version 1.0.0
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/danielmachado86/contracts/currency/protos"
	"github.com/danielmachado86/contracts/product-api/data"
	"github.com/gorilla/mux"
)

// A list of products in the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// swagger:parameters deleteProduct
type productIDParameterWrapper struct {
	// The id of the product to delete from the database
	// in: path
	// required: true
	ID int `json:"id"`
}

// swagger:response noContent
type productsNoContent struct {
}

type Products struct {
	l  *log.Logger
	cc protos.CurrencyClient
}

func NewProducts(l *log.Logger, cc protos.CurrencyClient) *Products {
	return &Products{l, cc}
}

type KeyProduct struct{}

func GetProductID(rw http.ResponseWriter, r *http.Request) int {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	return id
}

type GenericError struct {
	Message string `json:"messages"`
}
