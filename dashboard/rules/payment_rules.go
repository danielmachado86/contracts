package rules

import (
	"fmt"
	"math"
	"time"

	"github.com/danielmachado86/contracts/dashboard/data"
	"github.com/danielmachado86/contracts/dashboard/utils"
)

func createPayment(n string, v float64, t time.Time) *data.Payment {
	tr := &PaymentDueDate{time: t, name: n}
	return &data.Payment{Name: n, Value: v, Task: tr.Save()}
}

type Termination struct {
	time   time.Time
	params data.Params
}

func (r Termination) Run() *data.Payment {

	params := r.params

	n := fmt.Sprintf("Penalty payment")
	pp := &PeriodicPayment{params: r.params}
	v := pp.PeriodicPaymentValue(1) * float64(params.Penalty.Months)

	return createPayment(n, v, r.time)
}

func (r Termination) Save() *data.Payment {
	return r.Run().Save()
}

type Payment struct {
	time   time.Time
	value  float64
	name   string
	params data.Params
}

func (r Payment) Run() *data.Payment {

	return createPayment(r.name, r.value, r.time)
}

func (r Payment) Save() *data.Payment {
	return r.Run().Save()
}

type PeriodicPayment struct {
	time   time.Time
	params data.Params
}

func (r PeriodicPayment) Execute() {

	params := r.params

	for i := 1; i < r.PeriodicPaymentQuantity(); i++ {
		n := fmt.Sprintf("periodic_payment_%d", i)
		v := r.PeriodicPaymentValue(i)
		rcd := &PeriodicPaymentClosingDate{time: r.time, params: params, payment: i}
		cd := rcd.Run()
		p := &Payment{name: n, value: v}
		cd.Save()
		p.Save()
	}
}

func (r PeriodicPayment) PeriodicPaymentQuantity() int {

	params := r.params

	// Contract duration
	d := float64(params.Duration.Months)

	// Period: Time between payments
	p := float64(params.PaymentPeriod.Months)

	// Number of payments rounded to the least higher interger
	n := math.Ceil(float64(d / p))

	return int(n)
}

func residualPayment(price float64, period *utils.Period) float64 {
	// Period: Time between payments
	p := float64(period.Months)
	// Contract value
	v := float64(price)
	// Residual payment
	return math.Mod(v, p)

}

func (r PeriodicPayment) PeriodicPaymentValue(pn int) float64 {

	params := r.params

	// Period: Time between payments
	p := float64(params.PaymentPeriod.Months)
	// Contract duration
	d := float64(params.Duration.Months)
	// Contract value
	v := float64(params.Price)

	if pn == r.PeriodicPaymentQuantity() {
		residual := residualPayment(params.Price, params.PaymentPeriod)
		if residual > 0 {
			return residual
		}
	}

	// Payment value per period unit
	return (v / d) * p
}
