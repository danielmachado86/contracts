package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Chocolate",
		Price: 1.00,
		SKU:   "asd-dads-hhg",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
