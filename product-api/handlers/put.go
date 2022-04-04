package handlers

import (
	"net/http"

	"github.com/danielmachado86/contracts/product-api/data"
)

func (p *Products) Update(rw http.ResponseWriter, r *http.Request) {

	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	p.l.Println("[DEBUG] deleting record id", prod.ID)

	err := data.UpdateProduct(prod)
	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] product not found", err)
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Product not found in database"}, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
