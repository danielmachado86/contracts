package handlers

import (
	"fmt"
	"testing"
	"time"

	"github.com/danielmachado86/contracts/dashboard/data"
	"github.com/danielmachado86/contracts/dashboard/utils"
)

func TestSignatureDateRule(t *testing.T) {
	var c = &data.Contract{}

	tt := time.Now()

	r := SignatureDate{time: tt}
	task := r.Run(c)

	if task.GetDate() != tt {
		t.Fail()
	}
}

func TestStartDateRule(t *testing.T) {
	var c = &data.Contract{}
	tt := time.Now()
	sr := SignatureDate{time: tt}
	sd := sr.Run(c).GetDate()
	t_roundedToNextDay := time.Date(
		sd.Year(),
		sd.Month(),
		sd.Day()+1,
		0, 0, 0, 0,
		sd.Location(),
	)
	t_startDate := t_roundedToNextDay.AddDate(0, 0, 0)

	r := StartDate{time: tt, offset: &utils.Period{Days: 0}}
	task := r.Run(c)

	if task.GetDate() != t_startDate {
		t.Fail()
	}
}

func TestEndDateRule(t *testing.T) {
	var c = &data.Contract{
		Duration: &utils.Period{Months: 12},
	}

	tt := time.Now()
	sr := StartDate{time: tt, offset: &utils.Period{Days: 0}}
	t_startDate := sr.Run(c).GetDate()
	t_endDate := t_startDate.AddDate(0, 12, 0)

	rule := EndDate{time: tt, offset: &utils.Period{Days: 0}}
	task := rule.Run(c)

	if task.GetDate() != t_endDate {
		t.Fail()
	}
}

func TestAdvanceNoticeRule(t *testing.T) {
	var c = &data.Contract{
		Duration: &utils.Period{Months: 12},
	}
	tt := time.Now()
	er := EndDate{time: tt, offset: &utils.Period{Days: 0}}
	t_endDate := er.Run(c).GetDate()
	t_advanceNoticeDeadlineDate := t_endDate.AddDate(0, -3, 0)

	rule := AdvanceNoticeDeadline{time: tt, offset: &utils.Period{Days: 0}, period: &utils.Period{Months: 3}}
	task := rule.Run(c)

	if task.GetDate() != t_advanceNoticeDeadlineDate {
		t.Fail()
	}
}

func TestPaymentQuantity(t *testing.T) {
	var a = &data.ContractTemplate{
		Params: map[string]*utils.Period{
			"payment_period": {Months: 2},
		},
	}

	var c = &data.Contract{
		Duration: &utils.Period{Months: 13},
		Price:    13,
		Template: a,
	}

	pr := &PaymentValue{period: &utils.Period{Months: 2}}
	qty := pr.PeriodicPaymentQuantity(c)
	tQty := 7

	if qty != tQty {
		t.Error(fmt.Printf("The quantity of payments: %d is differerent to: %d", qty, tQty))
	}

}

func TestPaymentValue(t *testing.T) {
	var a = &data.ContractTemplate{
		Params: map[string]*utils.Period{
			"payment_period": {Months: 5},
		},
	}

	var c = &data.Contract{
		Duration: &utils.Period{Months: 12},
		Price:    24,
		Template: a,
	}

	// Payment Quantity: 3
	// Last payment value: 4
	// Ordinary payment value: 10

	t.Run("Test ordinary payment value", func(t *testing.T) {

		r1 := &PaymentValue{last: false, period: &utils.Period{Months: 5}}
		p1 := r1.Run(c)

		if int(p1.GetValue()) != 10 {
			t.Error(fmt.Printf("The value of payment: %d is differerent to: %d", int(p1.GetValue()), 5))
		}

	})

	t.Run("Test last payment (residual) value", func(t *testing.T) {
		r2 := &PaymentValue{last: true, period: &utils.Period{Months: 5}}
		p2 := r2.Run(c)

		if int(p2.GetValue()) != 4 {
			t.Error(fmt.Printf("The value of payment: %d is differerent to: %d", int(p2.GetValue()), 8))
		}

	})
}
