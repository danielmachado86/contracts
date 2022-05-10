package rules

import (
	"fmt"
	"math"
	"time"

	"github.com/danielmachado86/contracts/data"
)

func createPayment(n string, v float64, t time.Time) *data.Payment {
	tr := &PaymentDueDate{time: t, name: n}
	tr.Compute().Save()
	return &data.Payment{Name: n, Value: v, Task: tr.task}
}

type Termination struct {
	time    time.Time
	payment *data.Payment
}

func (r *Termination) Compute() Rule {

	params := data.GetParams()

	n := "penalty_payment"
	pp := &PeriodicPayment{}
	v := pp.PaymentValue(params.Price, params.Duration.Months) * float64(params.Penalty.Months)
	r.payment = createPayment(n, v, r.time)
	return r
}

func (r *Termination) Save() {
	r.payment.Save()
}

type Payment struct {
	time    time.Time
	value   float64
	name    string
	payment *data.Payment
}

func (r *Payment) Compute() Rule {
	r.payment = createPayment(r.name, r.value, r.time)
	return r
}

type PaymentGroup struct {
	PaymentRuleList []Rule
}

func NewPaymentGroup() *PaymentGroup {
	return &PaymentGroup{}
}

func (r *Payment) Save() {
	r.payment.Save()
}

type PeriodicPayment struct {
	group *PaymentGroup
}

func (r *PeriodicPayment) Compute() Rule {

	pq := r.PeriodicPaymentQuantity()
	pg := NewPaymentGroup()
	for i := 1; i <= pq; i++ {
		n := fmt.Sprintf("periodic_payment_%d", i)
		v := r.PeriodicPaymentValue(i, pq)
		pr := &Payment{name: n, value: v}

		pg.PaymentRuleList = append(pg.PaymentRuleList, pr)
	}
	r.group = pg
	return r
}

func (r *PeriodicPayment) Save() {
	rules := r.group
	for _, rule := range rules.PaymentRuleList {
		rule.Compute().Save()
	}
}

func (r *PeriodicPayment) PeriodicPaymentQuantity() int {

	params := data.GetParams()

	// Contract duration
	d := float64(params.Duration.Months)

	// Period: Time between payments
	p := float64(params.PaymentPeriod.Months)

	// Number of payments rounded to the least higher interger
	n := math.Ceil(float64(d / p))

	return int(n)
}

func (r *PeriodicPayment) residualPayment() float64 {
	params := data.GetParams()

	// Period: Time between payments
	p := float64(params.PaymentPeriod.Months)
	// Contract value
	v := float64(params.Price)
	// Residual payment
	return math.Mod(v, p)

}

func (r *PeriodicPayment) PaymentValue(price float64, duration int) float64 {
	return price / float64(duration)
}

func (r *PeriodicPayment) PeriodicPaymentValue(pn int, pq int) float64 {

	params := data.GetParams()

	// Period: Time between payments
	p := float64(params.PaymentPeriod.Months)
	// Contract duration
	d := params.Duration.Months
	// Contract value
	v := float64(params.Price)

	if pn == pq {
		residual := r.residualPayment()
		if residual > 0 {
			return residual
		}
	}

	// Payment value per period unit
	return r.PaymentValue(v, d) * p
}
