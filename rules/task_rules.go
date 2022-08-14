package rules

import (
	"context"
	"fmt"
	"time"

	db "github.com/danielmachado86/contracts/db"
)

// func roundDateToNextDay(t time.Time) time.Time {
// 	return time.Date(
// 		t.Year(),
// 		t.Month(),
// 		t.Day()+1,
// 		0, 0, 0, 0,
// 		t.Location(),
// 	)
// }

func paramsToMap(o_params []Parameter) map[string]interface{} {
	params := make(map[string]interface{})
	for _, param := range o_params {
		params[param.Name] = param.Value
	}
	return params
}

func checkInputs(inputs []RuleCategory, ruleRegistry map[string]Rule, result ResultContract) error {
	var err error
	for _, input := range inputs {
		rule := ruleRegistry[input.Name]
		result, err = rule.Run(ruleRegistry, result)
		if err != nil {
			return err
		}
	}
	return nil
}

type DatasourceFactory struct {
}

func (factory *DatasourceFactory) Create(store db.Store, spec *Spec) Rule {
	return &Datasource{
		Name:    spec.Name,
		Outputs: spec.Outputs,
		Store:   store,
	}
}

func NewDatasource(store db.Store, spec *Spec) Rule {
	factory := DatasourceFactory{}
	return factory.Create(store, spec)
}

type Datasource struct {
	Name    string
	Outputs []RuleCategory
	Store   db.Store
}

func (ds *Datasource) Run(ruleRegistry map[string]Rule, result ResultContract) (ResultContract, error) {

	id := result["id"].(string)

	resp, err := ds.Store.ListContractParties(context.Background(), id)
	if err != nil {
		return result, err
	}

	result[ds.Outputs[0].Name] = resp

	return result, nil
}

type EnoughPartiesFactory struct{}

func (factory *EnoughPartiesFactory) Create(spec *Spec) Rule {
	params := paramsToMap(spec.Parameters)
	return &Parties{
		Name:    spec.Name,
		Max:     int(params["max"].(float64)),
		Min:     int(params["min"].(float64)),
		Inputs:  spec.Inputs,
		Outputs: spec.Outputs,
	}
}

type Parties struct {
	Name    string
	Max     int
	Min     int
	Inputs  []RuleCategory
	Outputs []RuleCategory
}

func NewEnoughPartiesRule(spec *Spec) Rule {
	factory := EnoughPartiesFactory{}
	return factory.Create(spec)
}

func (rule *Parties) Run(ruleRegistry map[string]Rule, result ResultContract) (ResultContract, error) {
	err := checkInputs(rule.Inputs, ruleRegistry, result)
	if err != nil {
		return result, err
	}
	result["enough_parties"] = false
	numParties := len(result["party_list"].([]db.Party))
	if rule.Min <= numParties && numParties <= rule.Max {
		result["enough_parties"] = true
	}
	return result, nil
}

type EnoughSignaturesFactory struct{}

func (factory *EnoughSignaturesFactory) Create(spec *Spec) Rule {
	return &EnoughSignatures{
		Name:    spec.Name,
		Inputs:  spec.Inputs,
		Outputs: make(map[string]interface{}),
	}
}

type EnoughSignatures struct {
	Name       string
	Signatures []*db.Signature
	Inputs     []RuleCategory
	Outputs    map[string]interface{}
}

func NewEnoughSignaturesRule(spec *Spec) Rule {
	factory := EnoughSignaturesFactory{}
	return factory.Create(spec)
}

func (rule *EnoughSignatures) Run(ruleRegistry map[string]Rule, result ResultContract) (ResultContract, error) {

	err := checkInputs(rule.Inputs, ruleRegistry, result)
	if err != nil {
		return result, err
	}

	rule.Outputs[rule.Name] = rule.Signatures
	return result, nil
}

type IsSignedFactory struct{}

func (factory *IsSignedFactory) Create(spec *Spec) Rule {
	params := paramsToMap(spec.Parameters)
	return &IsSigned{
		Name:       spec.Name,
		Signatures: int(params["signatures"].(float64)),
		Inputs:     spec.Inputs,
		Outputs:    make(map[string]interface{}),
	}
}

type IsSigned struct {
	Name       string
	Signatures int
	Inputs     []RuleCategory
	Outputs    map[string]interface{}
}

func NewIsSignedRule(spec *Spec) Rule {
	factory := IsSignedFactory{}
	return factory.Create(spec)
}

func (rule *IsSigned) Run(ruleRegistry map[string]Rule, result ResultContract) (ResultContract, error) {

	err := checkInputs(rule.Inputs, ruleRegistry, result)
	if err != nil {
		return result, err
	}

	signatures := result["contract_signatures"].([]db.Signature)

	if len(signatures) == rule.Signatures {
		rule.Outputs[rule.Name] = true
	} else {
		rule.Outputs[rule.Name] = false
	}

	return result, nil
}

type SignatureDateFactory struct{}

func (factory *SignatureDateFactory) Create(spec *Spec) Rule {
	return &SignatureDate{
		Name:    spec.Name,
		Inputs:  spec.Inputs,
		Outputs: make(map[string]interface{}),
	}
}

type SignatureDate struct {
	Name    string
	Inputs  []RuleCategory
	Outputs map[string]interface{}
}

func NewSignatureDateRule(spec *Spec) Rule {
	factory := SignatureDateFactory{}
	return factory.Create(spec)
}

func (rule *SignatureDate) Run(ruleRegistry map[string]Rule, result ResultContract) (ResultContract, error) {

	err := checkInputs(rule.Inputs, ruleRegistry, result)
	if err != nil {
		return result, err
	}

	is_signed := result["contract_is_signed"].(bool)
	signatures := result["contract_signatures"].([]db.Signature)

	if !is_signed {
		return result, fmt.Errorf("contract not signed")
	}

	ld := time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC)
	for _, s := range signatures {
		d := s.CreatedAt
		if d.After(ld) {
			ld = d
		}
	}
	rule.Outputs[rule.Name] = ld
	return result, nil
}

// type StartDate struct {
// }

// func (r *StartDate) Run(inputs []Field, contract data.Contract) {

// 	exist := fieldExist(contract, CONTRACT_START_DATE)
// 	if exist {
// 		return
// 	}

// 	checkInputs(inputs, contract)

// 	sg := contract.Terms[CONTRACT_SIGNATURE_DATE].(time.Time)
// 	rounded := roundDateToNextDay(sg)
// 	contract.Terms[CONTRACT_SIGNATURE_DATE] = rounded
// 	return
// }

// type EndDate struct {
// }

// func (r *EndDate) Run(inputs []Field, contract data.Contract) {

// 	exist := fieldExist(contract, CONTRACT_END_DATE)
// 	if exist {
// 		return
// 	}

// 	checkInputs(inputs, contract)

// 	sd := contract.Terms[CONTRACT_START_DATE].(time.Time)

// 	td := sd.AddDate().Add(-time.Second * 1)

// 	contract.Terms[END_DATE_RULE_NAME] = td

// 	return
// }

// func (r *EndDate) Save() {
// 	r.task.Save()
// }

// type AdvanceNoticeDeadline struct {
// 	parties []*data.Party
// 	task    *data.Task
// }

// func (r *AdvanceNoticeDeadline) Compute() Rule {

// 	params := data.GetParams()

// 	// Period
// 	p := params.PeriodAN

// 	// End date rule
// 	er := &EndDate{parties: r.parties}
// 	er.Compute()
// 	// End date
// 	ed := er.task.GetDate()
// 	// Advance notice deadline
// 	nd := ed.AddDate(-p.Years, -p.Months, -p.Days)

// 	r.task = createTask("advance_notice_deadline", nd)
// 	return r
// }

// func (r *AdvanceNoticeDeadline) Save() {
// 	r.task.Save()
// }

// type NotificationDate struct {
// 	parties []*data.Party
// 	payment int
// 	task    *data.Task
// }

// func (r *NotificationDate) Compute() Rule {

// 	return r
// }

// func (r *NotificationDate) Save() {
// 	r.task.Save()
// }

// type PaymentDueDate struct {
// 	name string
// 	time time.Time
// 	task *data.Task
// }

// func (r *PaymentDueDate) Compute() Rule {

// 	//Payment closing date
// 	pd := r.time.AddDate(0, 0, 5)
// 	r.task = createTask(fmt.Sprintf("%s_due_date", r.name), pd)
// 	return r
// }

// func (r *PaymentDueDate) Save() {
// 	r.task.Save()
// }
