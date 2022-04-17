package data

import (
	"time"

	"github.com/danielmachado86/contracts/dashboard/utils"
)

type TemplateType int

const (
	Rental TemplateType = iota
)

// Defines agreement parameters
type ContractTemplate struct {
	Name   string
	Type   TemplateType
	Params map[string]*utils.Period
}

// Defines contract structure
type Contract struct {
	Duration *utils.Period
	Price    float64
	Template *ContractTemplate
}

type Task struct {
	Name string
	Date time.Time
}

func (t *Task) GetDate() time.Time {
	return t.Date
}

type Payment struct {
	Name  string
	Value float64
}

func (p *Payment) GetValue() float64 {
	return p.Value
}
