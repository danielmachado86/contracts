package rules

import (
	"fmt"
	"testing"
	"time"
)

func TestSignatureDateRule(t *testing.T) {
	tt := time.Now()

	r := SignatureDate{time: tt}
	task := r.Run()

	if task.GetDate() != tt {
		t.Fail()
	}
}

func TestStartDateRule(t *testing.T) {
	tt := time.Now()
	sr := SignatureDate{time: tt}
	sd := sr.Run().GetDate()
	t_roundedToNextDay := time.Date(
		sd.Year(),
		sd.Month(),
		sd.Day()+1,
		0, 0, 0, 0,
		sd.Location(),
	)
	t_startDate := t_roundedToNextDay.AddDate(0, 0, 0)

	r := StartDate{time: tt}
	task := r.Run()

	if task.GetDate() != t_startDate {
		t.Fail()
	}
}

func TestEndDateRule(t *testing.T) {
	tt := time.Now()

	sr := StartDate{time: tt}
	t_startDate := sr.Run().GetDate()
	t_endDate := t_startDate.AddDate(0, 12, 0)

	rule := EndDate{time: tt}
	task := rule.Run()

	if task.GetDate() != t_endDate {
		t.Fail()
	}
}

func TestAdvanceNoticeRule(t *testing.T) {
	tt := time.Now()

	er := EndDate{time: tt}
	t_endDate := er.Run().GetDate()
	t_advanceNoticeDeadlineDate := t_endDate.AddDate(0, -3, 0)

	rule := AdvanceNoticeDeadline{time: tt}
	task := rule.Run()

	if task.GetDate() != t_advanceNoticeDeadlineDate {
		t.Fail()
	}
}

func TestPaymentQuantity(t *testing.T) {
	r := PeriodicPayment{}

	qty := r.PeriodicPaymentQuantity()
	tQty := 3

	if qty != tQty {
		t.Error(fmt.Printf("The quantity of payments: %d is differerent to: %d", qty, tQty))
	}

}

func TestPaymentValue(t *testing.T) {
	tt := time.Now()

	// Payment Quantity: 3
	// Last payment value: 4
	// Ordinary payment value: 10

	t.Run("Test ordinary payment value", func(t *testing.T) {

		r1 := &PeriodicPayment{time: tt}
		p1 := r1.Run()

		if p1.Value != 24 {
			t.Error(fmt.Printf("The value of payment: %d is differerent to: %d", int(p1.GetValue()), 24))
		}

	})

	// t.Run("Test last payment (residual) value", func(t *testing.T) {
	// 	r2 := &PeriodicPayment{
	// 		time:          tt,
	// 		params:        params,
	// 	}
	// 	p2 := r2.Execute()

	// 	if int(p2.GetValue()) != 4 {
	// 		t.Error(fmt.Printf("The value of payment: %d is differerent to: %d", int(p2.GetValue()), 4))
	// 	}

	// })
}
