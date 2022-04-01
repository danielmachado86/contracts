// Package classification of Product API
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

	"github.com/danielmachado86/contracts/data"
	"github.com/gorilla/mux"
)

// A list of products in the response
// swagger:response productsResponse
type productsResponse struct {
	// All products in the system
	// in: Body
	Body []data.Product
}

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
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
