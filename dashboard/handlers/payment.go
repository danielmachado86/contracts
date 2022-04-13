package handlers

import (
	"math"

	"github.com/danielmachado86/contracts/dashboard/data"
)

// Defines payments and schedules
type PaymentManager struct {
}

func NewPaymentManager() *PaymentManager {
	return &PaymentManager{}
}

func (pm *PaymentManager) PaymentQuantity(c *data.Contract) int {
	params := c.Agreement.Params

	// Contract duration
	cDuration := float64(c.Duration.Months)

	// Period: Time between payments
	pPeriod := float64(params["payment_period"].Months)

	// Number of payments calculated, using contract duration and period
	pNumber := float64(cDuration / pPeriod)

	// Number of payments rounded to least higher interger
	pNumCeil := math.Ceil(pNumber)

	return int(pNumCeil)
}

func (pm *PaymentManager) lastPayment(c *data.Contract) float64 {
	params := c.Agreement.Params
	pPeriod := float64(params["payment_period"].Months)
	cDuration := float64(c.Duration.Months)
	cValue := float64(c.Price)

	r := math.Mod(c.Price, pPeriod)
	if r != 0 {
		return r
	}
	return (cValue / cDuration) * pPeriod
}

func (pm *PaymentManager) PaymentValue(c *data.Contract, last bool) float64 {

	params := c.Agreement.Params
	// Contract duration
	cDuration := float64(c.Duration.Months)
	cValue := float64(c.Price)

	// Period: Time between payments
	pPeriod := float64(params["payment_period"].Months)

	if last {
		return pm.lastPayment(c)
	}

	// Payment per period unit
	return (cValue / cDuration) * pPeriod
}
