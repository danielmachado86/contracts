package rules

import (
	"fmt"
	"math"
	"time"

	"github.com/danielmachado86/contracts/dashboard/data"
)

func createPayment(n string, v float64, t time.Time) *data.Payment {
	tr := &PaymentDueDate{time: t, name: n}
	tt := tr.Run()
	return &data.Payment{Name: n, Value: v, Task: tt.Save()}
}

type Termination struct {
	time time.Time
}

func (r Termination) Run() *data.Payment {

	params := data.GetParams()

	n := "Penalty payment"
	pp := &PeriodicPayment{}
	v := pp.PeriodicPaymentValue(1, 1) * float64(params.Penalty.Months)

	return createPayment(n, v, r.time)
}

func (r Termination) Save() *data.Payment {
	return r.Run().Save()
}

type Payment struct {
	time  time.Time
	value float64
	name  string
}

func (r Payment) Run() *data.Payment {
	return createPayment(r.name, r.value, r.time)
}

func (r Payment) Save() *data.Payment {
	return r.Run().Save()
}

type PeriodicPayment struct {
}

func (r PeriodicPayment) Configure() *PaymentGroup {

	pq := r.PeriodicPaymentQuantity()
	pg := &PaymentGroup{}
	for i := 1; i <= pq; i++ {
		n := fmt.Sprintf("periodic_payment_%d", i)
		v := r.PeriodicPaymentValue(i, pq)
		pr := Payment{name: n, value: v}

		pg.PaymentRuleList = append(pg.PaymentRuleList, pr)
	}
	return pg
}

func (r PeriodicPayment) Execute() {
	rules := r.Configure()
	for _, rule := range rules.PaymentRuleList {
		rule.Run()
	}
}

func (r PeriodicPayment) PeriodicPaymentQuantity() int {

	params := data.GetParams()

	// Contract duration
	d := float64(params.Duration.Months)

	// Period: Time between payments
	p := float64(params.PaymentPeriod.Months)

	// Number of payments rounded to the least higher interger
	n := math.Ceil(float64(d / p))

	return int(n)
}

func (r PeriodicPayment) residualPayment() float64 {
	params := data.GetParams()

	// Period: Time between payments
	p := float64(params.PaymentPeriod.Months)
	// Contract value
	v := float64(params.Price)
	// Residual payment
	return math.Mod(v, p)

}

func (r PeriodicPayment) PeriodicPaymentValue(pn int, pq int) float64 {

	params := data.GetParams()

	// Period: Time between payments
	p := float64(params.PaymentPeriod.Months)
	// Contract duration
	d := float64(params.Duration.Months)
	// Contract value
	v := float64(params.Price)

	if pn == pq {
		residual := r.residualPayment()
		if residual > 0 {
			return residual
		}
	}

	// Payment value per period unit
	return (v / d) * p
}
