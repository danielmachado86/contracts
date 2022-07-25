package rules

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"time"

	"github.com/danielmachado86/contracts/data"
	"github.com/stretchr/testify/require"
)

func TestSignatureDateRule(t *testing.T) {

	file, err := ioutil.ReadFile("rental_contract.json")
	require.NoError(t, err)

	template := TemplateSpecification{}
	err = json.Unmarshal([]byte(file), &template)
	require.NoError(t, err)

	user1 := data.NewUser("Daniel M")
	user2 := data.NewUser("Jimena L")

	party1 := data.NewParty(user1)
	party2 := data.NewParty(user2)

	data.NewSignature(party1)
	data.NewSignature(party2)

	var parties []*data.Party
	parties = append(parties, party1)
	parties = append(parties, party2)

	ruleRegistry := make(map[string]Rule)

	for _, attribute := range template.Attributes {
		ruleRegistry[attribute.Name], err = CreateRule(attribute)
		require.NoError(t, err)
	}

	ruleRegistry["contract_parties"].(*Parties).Parties = parties

	rule := ruleRegistry["contract_signature_date"]
	err = rule.Run(ruleRegistry)
	require.NoError(t, err)

	got := rule.GetOutput("contract_signature_date").(time.Time)
	require.WithinDuration(t, got, time.Now(), time.Second)

}

// func TestStartDateRule(t *testing.T) {
// 	parties := []*data.Party{
// 		data.Party1,
// 		data.Party2,
// 	}

// 	r := StartDate{parties: parties}
// 	r.Compute()

// 	got := r.task.GetDate()

// 	want := time.Date(2022, 4, 28, 0, 0, 0, 0, time.UTC)

// 	if got != want {
// 		t.Error(fmt.Print("Got: ", got, "\n", "Want: ", want, "\n"))
// 	}
// }

// func TestEndDateRule(t *testing.T) {

// 	parties := []*data.Party{
// 		data.Party1,
// 		data.Party2,
// 	}

// 	r := EndDate{parties: parties}
// 	r.Compute()

// 	got := r.task.GetDate()

// 	want := time.Date(2023, 4, 27, 23, 59, 59, 0, time.UTC)

// 	if got != want {
// 		t.Error(fmt.Print("Got: ", got, "\n", "Want: ", want, "\n"))
// 	}
// }

// func TestAdvanceNoticeRule(t *testing.T) {
// 	parties := []*data.Party{
// 		data.Party1,
// 		data.Party2,
// 	}

// 	r := AdvanceNoticeDeadline{parties: parties}
// 	r.Compute()

// 	got := r.task.GetDate()

// 	want := time.Date(2023, 1, 27, 23, 59, 59, 0, time.UTC)

// 	if got != want {
// 		t.Error(fmt.Print("Got: ", got, "\n", "Want: ", want, "\n"))
// 	}
// }

// func TestPaymentQuantity(t *testing.T) {
// 	r := PeriodicPayment{}

// 	got := r.PeriodicPaymentQuantity()
// 	want := 3

// 	if got != want {
// 		t.Error(fmt.Printf("The quantity of payments: %d is differerent to: %d", got, want))
// 	}

// }

// func TestPeriodicPayment(t *testing.T) {
// 	pp := PeriodicPayment{}
// 	pp.Compute().Save()

// 	got := len(data.ContractInst.GetAttributes())
// 	want := 6

// 	if got != want {
// 		t.Error(fmt.Printf("The value of payment: %d is differerent to: %d", got, want))
// 	}
// }

// func TestTermination(t *testing.T) {
// 	r := Termination{}
// 	r.Compute().Save()

// 	attr := data.ContractInst.GetAttributes()
// 	got := attr["penalty_payment"].(*data.Payment).Value
// 	want := 6.0

// 	if got != want {
// 		t.Error(fmt.Printf("The value of payment: %f is differerent to: %f", got, want))
// 	}
// }
