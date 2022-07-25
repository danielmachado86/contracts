package rules

import (
	"fmt"
)

type Value interface{}

type Parameter struct {
	Name  string `json:"name"`
	Value Value  `json:"value"`
}
type Attributes struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Parameters  []Parameter `json:"parameters"`
	Inputs      []string    `json:"inputs"`
	Outputs     []string    `json:"outputs"`
}

type TemplateSpecification struct {
	Category    string        `json:"category"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Attributes  []*Attributes `json:"attributes"`
}

type Rule interface {
	Run(map[string]Rule) error
	GetName() string
	GetOutput(string) interface{}
}

type RuleFactory interface {
	Create(*Attributes) Rule
}

func CreateRule(attributes *Attributes) (Rule, error) {
	if attributes.Name == "contract_parties" {
		return NewPartiesRule(attributes)
	}
	if attributes.Name == "contract_signatures" {
		return NewSignaturesRule(attributes)

	}
	if attributes.Name == "contract_is_signed" {
		return NewIsSignedRule(attributes)
	}
	if attributes.Name == "contract_signature_date" {
		return NewSignatureDateRule(attributes)
	}
	if attributes.Name == "contract_start_date" {
		return nil, nil
	}
	if attributes.Name == "contract_end_date" {
		return nil, nil
	}
	if attributes.Name == "contract_renew_deadline" {
		return nil, nil
	}
	if attributes.Name == "contract_installment_plan" {
		return nil, nil
	}

	err := fmt.Errorf("Rule %s doesn't exist", attributes.Name)

	return nil, err

}
