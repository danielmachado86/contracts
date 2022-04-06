package handlers

import (
	"net/http"

	"github.com/danielmachado86/contracts/product-api/data"
)

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	p.l.Printf("Prod: %#v", prod)

	data.AddProduct(prod)
}
