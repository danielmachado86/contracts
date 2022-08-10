package rules

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/danielmachado86/contracts/db"
)

type Value interface{}

type RuleCategory struct {
	Type string
	Name string
}

type Parameter struct {
	Name  string `json:"name"`
	Value Value  `json:"value"`
}
type Spec struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Type        string         `json:"type"`
	Parameters  []Parameter    `json:"parameters"`
	Inputs      []RuleCategory `json:"inputs"`
	Outputs     []RuleCategory `json:"outputs"`
}

type RuleEngine []*Spec
type ResultContract map[string]interface{}
type Metadata map[string]interface{}

type Rule interface {
	Run(map[string]Rule, ResultContract) (map[string]interface{}, error)
	Save(map[string]Rule, ResultContract) error
	GetName() string
}

type RuleFactory interface {
	Create(*Spec) Rule
}

func CreateRule(attributes *Spec) (Rule, error) {
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

func CalculateTerms(ctx context.Context, store db.Store, meta Metadata) error {
	file, err := ioutil.ReadFile("rental_contract.json")
	if err != nil {
		return err
	}
	template := make(RuleEngine, 0)
	err = json.Unmarshal([]byte(file), template)
	if err != nil {
		return err
	}

	result := make(ResultContract)
	result["metadata"] = meta

	ruleRegistry := make(map[string]Rule)

	for _, spec := range template {
		ruleRegistry[spec.Name], err = CreateRule(spec)
		if err != nil {
			return err
		}
	}

	for _, rule := range ruleRegistry {
		err = rule.Save(ruleRegistry, result)
		if err != nil {
			return err
		}
	}

	return nil
}
