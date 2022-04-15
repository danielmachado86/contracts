package data

import (
	"time"

	"github.com/danielmachado86/contracts/dashboard/utils"
)

type AgreementType int

const (
	Rental AgreementType = iota
)

// Defines agreement parameters
type Agreement struct {
	Name   string
	Type   AgreementType
	Params map[string]*utils.Date
}

func NewAgreement(n string, a AgreementType, p map[string]*utils.Date) *Agreement {
	return &Agreement{}
}

// Defines contract structure
type Contract struct {
	Duration  *utils.Date
	Price     float64
	Agreement *Agreement
}

type Contracts []*Contract

type Task struct {
	Name string
	Date time.Time
}

type Payment struct {
	Name   string
	Remain float64
	Value  float64
	Date   time.Time
}
