package handlers

import (
	"fmt"
	"math"
	"time"

	"github.com/danielmachado86/contracts/dashboard/data"
	"github.com/danielmachado86/contracts/dashboard/utils"
)

func createPayment(n string, v float64, r ScheduleRuleManager) *data.Payment {
	return &data.Payment{Name: n, Value: v, Task: r.Run()}
}

type TerminationValue struct {
	penalty  float64
	period   *utils.Period
	duration *utils.Period
	time     time.Time
	price    float64
	dueDate  *PaymentDueDate
}

func (r TerminationValue) Run() *data.Payment {
	n := fmt.Sprintf("Penalty payment")
	v := PaymentValue(r.duration, r.price, r.period, 1) * r.penalty
	rdd := &PaymentDueDate{time: r.time, name: "overdue_payment"}

	return createPayment(n, v, rdd)
}

func (r TerminationValue) Save() *data.Payment {
	r.dueDate.Save()
	return r.Run().Save()
}

type PeriodicPaymentValue struct {
	time          time.Time
	params        data.Params
	paymentNumber int
	closingDate   *PeriodicPaymentClosingDate
}

func (r PeriodicPaymentValue) Run() *data.Payment {

	params := r.params

	n := fmt.Sprintf("Recurring payment #%d", r.paymentNumber)
	v := PaymentValue(params.Duration, params.Price, params.PaymentPeriod, r.paymentNumber)

	rcd := &PeriodicPaymentClosingDate{time: r.time, params: params, payment: r.paymentNumber}
	r.closingDate = rcd

	return createPayment(n, v, rcd)
}

func (r PeriodicPaymentValue) Save() *data.Payment {
	r.closingDate.Save()
	return r.Run().Save()
}

func PeriodicPaymentQuantity(duration *utils.Period, period *utils.Period) int {
	// Contract duration
	d := float64(duration.Months)

	// Period: Time between payments
	p := float64(period.Months)

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

func PaymentValue(duration *utils.Period, price float64, period *utils.Period, payment int) float64 {
	// Period: Time between payments
	p := float64(period.Months)
	// Contract duration
	d := float64(duration.Months)
	// Contract value
	v := float64(price)

	if payment == PeriodicPaymentQuantity(duration, period) {
		residual := residualPayment(price, period)
		if residual > 0 {
			return residual
		}
	}

	// Payment value per period unit
	return (v / d) * p
}
