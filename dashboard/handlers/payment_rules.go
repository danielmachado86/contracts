package handlers

import (
	"math"

	"github.com/danielmachado86/contracts/dashboard/data"
	"github.com/danielmachado86/contracts/dashboard/utils"
)

type TermminationValue struct {
	penalty float64
	period  *utils.Period
}

func (tv TermminationValue) Run(c *data.Contract) Calculator {
	p := &data.Payment{
		Name:  "Payment",
		Value: tv.penalty * PeriodicPaymentValue(c, tv.period, false),
	}
	return p
}

type PaymentValue struct {
	last   bool
	period *utils.Period
}

func (pr PaymentValue) Run(c *data.Contract) Calculator {

	pValue := PeriodicPaymentValue(c, pr.period, pr.last)

	p := &data.Payment{
		Name:  "Payment",
		Value: pValue,
	}

	return p
}

func (pr *PaymentValue) PeriodicPaymentQuantity(c *data.Contract) int {
	params := c.Template.Params

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

func lastPeriodicPayment(c *data.Contract) float64 {
	params := c.Template.Params
	pPeriod := float64(params["payment_period"].Months)
	cDuration := float64(c.Duration.Months)
	cValue := float64(c.Price)

	r := math.Mod(c.Price, pPeriod)
	if r != 0 {
		return r
	}
	return (cValue / cDuration) * pPeriod
}

func PeriodicPaymentValue(c *data.Contract, period *utils.Period, last bool) float64 {

	// Contract duration
	cDuration := float64(c.Duration.Months)
	cValue := float64(c.Price)

	// Period: Time between payments
	pPeriod := float64(period.Months)

	if last {
		return lastPeriodicPayment(c)
	}

	// Payment per period unit
	return (cValue / cDuration) * pPeriod
}
