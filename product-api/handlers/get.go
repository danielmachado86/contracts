package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/danielmachado86/contracts/currency/protos"
	"github.com/danielmachado86/contracts/product-api/data"
	"github.com/gorilla/mux"
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

func (p *Products) GetProductByID(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}

	prod, err := data.GetProductByID(id)

	if err != nil {
		http.Error(rw, "Error getting product", http.StatusNotFound)
	}

	rr := &protos.RateRequest{
		Base:        protos.Currencies_EUR,
		Destination: protos.Currencies_GBP,
	}

	resp, err := p.cc.GetRate(context.Background(), rr)
	if err != nil {
		p.l.Println("[Error] error getting new rate", err)
	}

	prod.Price = prod.Price * resp.Rate

	err = data.ToJSON(prod, rw)
	if err != nil {
		p.l.Println("[Error] error serializing product", err)
	}
}
