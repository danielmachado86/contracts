package rules

import (
	"fmt"
	"testing"
	"time"

	"github.com/danielmachado86/contracts/dashboard/data"
)

func TestSignatureDateRule(t *testing.T) {

	parties := []*data.Party{
		data.Party1,
		data.Party2,
	}

	r := SignatureDate{parties: parties}
	r.Compute()

	got := r.task.GetDate()

	want := time.Date(2022, 4, 27, 15, 0, 0, 0, time.UTC)

	if got != want {
		t.Error(fmt.Print("Got: ", got, "\n", "Want: ", want, "\n"))
	}
}

func TestStartDateRule(t *testing.T) {
	parties := []*data.Party{
		data.Party1,
		data.Party2,
	}

	r := StartDate{parties: parties}
	r.Compute()

	got := r.task.GetDate()

	want := time.Date(2022, 4, 28, 0, 0, 0, 0, time.UTC)

	if got != want {
		t.Error(fmt.Print("Got: ", got, "\n", "Want: ", want, "\n"))
	}
}

func TestEndDateRule(t *testing.T) {

	parties := []*data.Party{
		data.Party1,
		data.Party2,
	}

	r := EndDate{parties: parties}
	r.Compute()

	got := r.task.GetDate()

	want := time.Date(2023, 4, 27, 23, 59, 59, 0, time.UTC)

	if got != want {
		t.Error(fmt.Print("Got: ", got, "\n", "Want: ", want, "\n"))
	}
}

func TestAdvanceNoticeRule(t *testing.T) {
	parties := []*data.Party{
		data.Party1,
		data.Party2,
	}

	r := AdvanceNoticeDeadline{parties: parties}
	r.Compute()

	got := r.task.GetDate()

	want := time.Date(2023, 1, 27, 23, 59, 59, 0, time.UTC)

	if got != want {
		t.Error(fmt.Print("Got: ", got, "\n", "Want: ", want, "\n"))
	}
}

func TestPaymentQuantity(t *testing.T) {
	r := PeriodicPayment{}

	got := r.PeriodicPaymentQuantity()
	want := 3

	if got != want {
		t.Error(fmt.Printf("The quantity of payments: %d is differerent to: %d", got, want))
	}

}

func TestPeriodicPayment(t *testing.T) {
	pp := PeriodicPayment{}
	pp.Compute().Save()

	got := len(data.ContractInst.GetAttributes())
	want := 6

	if got != want {
		t.Error(fmt.Printf("The value of payment: %d is differerent to: %d", got, want))
	}
}

func TestTermination(t *testing.T) {
	r := Termination{}
	r.Compute().Save()

	attr := data.ContractInst.GetAttributes()
	got := attr["penalty_payment"].(*data.Payment).Value
	want := 6.0

	if got != want {
		t.Error(fmt.Printf("The value of payment: %f is differerent to: %f", got, want))
	}
}
