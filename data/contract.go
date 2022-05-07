package data

import (
	"time"

	"github.com/danielmachado86/contracts/dashboard/utils"
)

type Params struct {
	PeriodAN      *utils.Period // Advanced notice period
	PaymentPeriod *utils.Period
	Offset        *utils.Period
	Duration      *utils.Period
	Price         float64
	Penalty       *utils.Period //Overdue payment penalty
}

var params = Params{
	PeriodAN:      &utils.Period{Months: 3},
	PaymentPeriod: &utils.Period{Months: 5},
	Offset:        &utils.Period{Days: 0},
	Duration:      &utils.Period{Months: 12},
	Price:         24,
	Penalty:       &utils.Period{Months: 3},
}

func GetParams() Params {
	return params
}

// Defines agreement parameters
type ContractTemplate struct {
	Name   string
	Params *Params
}

// Defines contract structure
type Contract struct {
	Attributes map[string]interface{}
}

func (c *Contract) GetAttributes() map[string]interface{} {
	return c.Attributes
}

func NewContract() *Contract {
	return &Contract{Attributes: make(map[string]interface{})}
}

var ContractInst = NewContract()

type Task struct {
	Name string
	Date time.Time
}

func (t Task) AddPeriod(p *utils.Period) time.Time {
	// Start date
	d := t.GetDate()
	// Termination date
	return d.AddDate(p.Years, p.Months, p.Days)
}

func (t Task) GetDate() time.Time {
	return t.Date
}

func (t Task) Save() {
	ContractInst.Attributes[t.Name] = t
}

type Payment struct {
	Name  string
	Value float64
	Task  *Task
}

func (p *Payment) GetValue() float64 {
	return p.Value
}

func (p *Payment) Save() {
	ContractInst.Attributes[p.Name] = p
}
