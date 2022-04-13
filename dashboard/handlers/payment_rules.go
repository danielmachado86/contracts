package handlers

import (
	"github.com/danielmachado86/contracts/dashboard/data"
)

type PaymentRule interface {
	Calculate(pm *data.PaymentManager, c *data.Contract) data.Payment
}

type PaymentValueRule struct {
	last bool
}

func (pr *PaymentValueRule) Calculate(pm *data.PaymentManager, c *data.Contract) *data.Payment {

	pValue := pm.PaymentValue(c, pr.last)

	p := &data.Payment{
		Name:  "Payment",
		Value: pValue,
	}

	return p
}
