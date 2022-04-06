package handlers

import (
	"net/http"

	"github.com/danielmachado86/contracts/product-api/data"
)

// swagger:route GET /products products listProducts
// Returs a list of products
// responses:
//  200: productsResponse

// GetProducts returns the products from the data store
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] Get list of records")

	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
