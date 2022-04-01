package handlers

import (
	"net/http"

	"github.com/danielmachado86/contracts/data"
)

func (p *Products) DeleteProducts(rw http.ResponseWriter, r *http.Request) {

	id := GetProductID(rw, r)

	p.l.Println("[DEBUG] deleting record id", id)

	err := data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		p.l.Println("[Error] deleting error id does not exist")
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	if err != nil {
		p.l.Println("[Error] deleting record ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)

}
