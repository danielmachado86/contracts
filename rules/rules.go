package rules

import (
	"context"
	"encoding/json"
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

type Specs []*Spec
type ResultContract map[string]interface{}
type Metadata map[string]interface{}

type ExternalData interface {
	Run(db.Store, ResultContract) (ResultContract, error)
}

type Rule interface {
	Run(map[string]Rule, ResultContract) (ResultContract, error)
}

type RuleFactory interface {
	Create(*Spec) Rule
}

func (spec *Spec) RegisterDataSource(ruleRegistry map[string]Rule, store db.Store) (map[string]Rule, error) {
	ruleRegistry[spec.Name] = NewDatasource(store, spec)
	return ruleRegistry, nil
}

func (spec *Spec) RegisterRule(ruleRegistry map[string]Rule) (map[string]Rule, error) {
	var rule Rule
	switch spec.Name {
	case "contract_parties":
		rule = NewEnoughPartiesRule(spec)
	case "contract_signatures":
		rule = NewEnoughSignaturesRule(spec)
	case "contract_is_signed":
		rule = NewIsSignedRule(spec)
	case "contract_signature_date":
		rule = NewSignatureDateRule(spec)
	default:
		rule = nil
	}
	ruleRegistry[spec.Name] = rule
	return ruleRegistry, nil

}

func CalculateTerms(ctx context.Context, store db.Store, meta Metadata) error {
	file, err := ioutil.ReadFile("rental_contract.json")
	if err != nil {
		return err
	}
	specs := make(Specs, 0)
	err = json.Unmarshal([]byte(file), &specs)
	if err != nil {
		return err
	}

	result := make(ResultContract)
	result["metadata"] = meta

	ruleRegistry := make(map[string]Rule)

	for _, spec := range specs {
		switch spec.Type {
		case "datasource":
			ruleRegistry, err = spec.RegisterDataSource(ruleRegistry, store)
		case "rule":
			ruleRegistry, err = spec.RegisterRule(ruleRegistry)
		}
		if err != nil {
			return err
		}
	}

	for _, rule := range ruleRegistry {
		result, err = rule.Run(ruleRegistry, result)
		if err != nil {
			return err
		}
	}

	return nil
}
