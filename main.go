package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	http.ListenAndServe(":9090", nil)
}
