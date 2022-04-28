package rules

import (
	"fmt"
	"testing"
	"time"

	"github.com/danielmachado86/contracts/dashboard/data"
)

func TestSignatureDateRule(t *testing.T) {
	tt := time.Now()

	r := SignatureDate{time: tt}
	r.Compute()

	if r.task.GetDate() != tt {
		t.Fail()
	}
}

func TestStartDateRule(t *testing.T) {
	tt := time.Now()
	sr := SignatureDate{time: tt}
	sr.Compute()
	sd := sr.task.GetDate()
	t_roundedToNextDay := time.Date(
		sd.Year(),
		sd.Month(),
		sd.Day()+1,
		0, 0, 0, 0,
		sd.Location(),
	)
	t_startDate := t_roundedToNextDay.AddDate(0, 0, 0)

	r := StartDate{time: tt}
	r.Compute()

	if r.task.GetDate() != t_startDate {
		t.Fail()
	}
}

func TestEndDateRule(t *testing.T) {
	tt := time.Now()

	sr := StartDate{time: tt}
	sr.Compute()
	t_startDate := sr.task.GetDate()
	t_endDate := t_startDate.AddDate(0, 12, 0)

	rule := EndDate{time: tt}
	rule.Compute()

	if rule.task.GetDate() != t_endDate {
		t.Fail()
	}
}

func TestAdvanceNoticeRule(t *testing.T) {
	tt := time.Now()

	er := EndDate{time: tt}
	er.Compute()
	t_endDate := er.task.GetDate()
	t_advanceNoticeDeadlineDate := t_endDate.AddDate(0, -3, 0)

	rule := AdvanceNoticeDeadline{time: tt}
	rule.Compute()

	if rule.task.GetDate() != t_advanceNoticeDeadlineDate {
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
	pp := PeriodicPayment{}
	pp.Compute().Save()

	if len(data.ContractInst.GetAttributes()) != 6 {
		t.Error(fmt.Printf("The value of payment: %d is differerent to: %d", len(data.ContractInst.GetAttributes()), 3))
	}
}

func TestTermination(t *testing.T) {
	pp := Termination{}
	pp.Compute().Save()

	attr := data.ContractInst.GetAttributes()

	if attr["penalty_payment"].(*data.Payment).Value != 6.0 {
		t.Error(fmt.Printf("The value of payment: %f is differerent to: %f", attr["penalty_payment"].(*data.Payment).Value, 12.0))
	}
}
