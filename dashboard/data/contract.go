package data

import (
	"time"

	"github.com/danielmachado86/contracts/dashboard/utils"
)

type TemplateType int

const (
	Rental TemplateType = iota
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

var scheduleRules = []string{
	"signature_date",
	"start_date",
	"end_date",
	"advance_notice_deadline",
}

var paymentRules = []string{
	"periodic_payment",
	"termination",
}

// Defines agreement parameters
type ContractTemplate struct {
	Name          string
	Type          TemplateType
	Params        *Params
	ScheduleRules []string
	PaymentRules  []string
}

func NewContractTemplate(n string, t TemplateType, p *Params, sr []string, pr []string) *ContractTemplate {
	return &ContractTemplate{
		Name:          n,
		Type:          t,
		Params:        p,
		ScheduleRules: scheduleRules,
		PaymentRules:  paymentRules,
	}
}

// Defines contract structure
type Contract struct {
	Attributes map[string]interface{}
}

type attributes map[string]interface{}

func GetAttributes() map[string]interface{} {
	return make(attributes)
}

type Task struct {
	Name string
	Date time.Time
}

func (t *Task) AddPeriod(p *utils.Period) time.Time {
	// Start date
	d := t.GetDate()
	// Termination date
	return d.AddDate(p.Years, p.Months, p.Days)
}

func (t *Task) GetDate() time.Time {
	return t.Date
}

func (t *Task) Save() *Task {
	attributes := GetAttributes()
	attributes[t.Name] = t
	return t
}

type Payment struct {
	Name  string
	Value float64
	Task  *Task
}

func (p *Payment) GetValue() float64 {
	return p.Value
}

func (p *Payment) Save() *Payment {
	attributes := GetAttributes()
	attributes[p.Name] = p
	return p
}
