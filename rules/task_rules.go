package rules

import (
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

func getParams(attributes *Attributes) map[string]interface{} {
	params := make(map[string]interface{})
	for _, param := range attributes.Parameters {
		params[param.Name] = param.Value
	}
	return params
}

func checkInputs(inputs []string, ruleRegistry map[string]Rule) error {
	var err error
	for _, input := range inputs {
		rule := ruleRegistry[input]
		err = rule.Run(ruleRegistry)
		if err != nil {
			return err
		}
	}
	return nil
}

type PartiesFactory struct{}

func (factory *PartiesFactory) Create(attributes *Attributes) Rule {
	params := getParams(attributes)
	return &Parties{
		Name:    attributes.Name,
		Max:     int(params["max"].(float64)),
		Min:     int(params["min"].(float64)),
		Inputs:  attributes.Inputs,
		Outputs: make(map[string]interface{}),
	}
}

type Parties struct {
	Name    string
	Max     int
	Min     int
	Parties []*db.Party
	Inputs  []string
	Outputs map[string]interface{}
}

func NewPartiesRule(attributes *Attributes) (Rule, error) {
	factory := PartiesFactory{}
	return factory.Create(attributes), nil
}

func (rule *Parties) Run(ruleRegistry map[string]Rule) error {

	err := checkInputs(rule.Inputs, ruleRegistry)
	if err != nil {
		return err
	}

	rule.Outputs[rule.Name] = rule.Parties
	return nil
}

func (rule *Parties) GetName() string {
	return rule.Name
}

func (rule *Parties) GetOutput(outputName string) interface{} {
	return rule.Outputs[outputName]
}

type SignaturesFactory struct{}

func (factory *SignaturesFactory) Create(attributes *Attributes) Rule {
	return &Signatures{
		Name:    attributes.Name,
		Inputs:  attributes.Inputs,
		Outputs: make(map[string]interface{}),
	}
}

type Signatures struct {
	Name       string
	Signatures []*db.Signature
	Inputs     []string
	Outputs    map[string]interface{}
}

func NewSignaturesRule(attributes *Attributes) (Rule, error) {
	factory := SignaturesFactory{}
	return factory.Create(attributes), nil
}

func (rule *Signatures) Run(ruleRegistry map[string]Rule) error {

	err := checkInputs(rule.Inputs, ruleRegistry)
	if err != nil {
		return err
	}

	rule.Outputs[rule.Name] = rule.Signatures
	return nil
}

func (rule *Signatures) GetName() string {
	return rule.Name
}

func (rule *Signatures) GetOutput(outputName string) interface{} {
	return rule.Outputs[outputName]
}

type IsSignedFactory struct{}

func (factory *IsSignedFactory) Create(attributes *Attributes) Rule {
	params := getParams(attributes)
	return &IsSigned{
		Name:       attributes.Name,
		Signatures: int(params["signatures"].(float64)),
		Inputs:     attributes.Inputs,
		Outputs:    make(map[string]interface{}),
	}
}

type IsSigned struct {
	Name       string
	Signatures int
	Inputs     []string
	Outputs    map[string]interface{}
}

func NewIsSignedRule(attributes *Attributes) (Rule, error) {
	factory := IsSignedFactory{}
	return factory.Create(attributes), nil
}

func (rule *IsSigned) Run(ruleRegistry map[string]Rule) error {

	err := checkInputs(rule.Inputs, ruleRegistry)
	if err != nil {
		return err
	}

	signatures := ruleRegistry["contract_signatures"].GetOutput("contract_signatures").([]*db.Signature)

	if len(signatures) == rule.Signatures {
		rule.Outputs[rule.Name] = true
	} else {
		rule.Outputs[rule.Name] = false
	}

	return nil
}

func (rule *IsSigned) GetName() string {
	return rule.Name
}

func (rule *IsSigned) GetOutput(outputName string) interface{} {
	return rule.Outputs[outputName]
}

type SignatureDateFactory struct{}

func (factory *SignatureDateFactory) Create(attributes *Attributes) Rule {
	return &SignatureDate{
		Name:    attributes.Name,
		Inputs:  attributes.Inputs,
		Outputs: make(map[string]interface{}),
	}
}

type SignatureDate struct {
	Name    string
	Inputs  []string
	Outputs map[string]interface{}
}

func NewSignatureDateRule(attributes *Attributes) (Rule, error) {
	factory := SignatureDateFactory{}
	return factory.Create(attributes), nil
}

func (rule *SignatureDate) Run(ruleRegistry map[string]Rule) error {

	err := checkInputs(rule.Inputs, ruleRegistry)
	if err != nil {
		return err
	}

	is_signed := ruleRegistry["contract_is_signed"].GetOutput("contract_is_signed").(bool)
	signatures := ruleRegistry["contract_signatures"].GetOutput("contract_signatures").([]*db.Signature)

	if !is_signed {
		return fmt.Errorf("contract not signed")
	}

	ld := time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC)
	for _, s := range signatures {
		d := s.CreatedAt
		if d.After(ld) {
			ld = d
		}
	}
	rule.Outputs[rule.Name] = ld
	return nil
}

func (rule *SignatureDate) GetName() string {
	return rule.Name
}

func (rule *SignatureDate) GetOutput(outputName string) interface{} {
	return rule.Outputs[outputName]
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
