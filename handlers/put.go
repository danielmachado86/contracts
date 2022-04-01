package handlers

import (
	"net/http"

	"github.com/danielmachado86/contracts/data"
)

func (p *Products) Update(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Product")

	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	err := data.UpdateProduct(prod)
	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] product not found", err)
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Product not found in database"}, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
