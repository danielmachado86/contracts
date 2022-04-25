package rules

import (
	"github.com/danielmachado86/contracts/dashboard/data"
)

type ScheduleRule interface {
	Run() *data.Task
	Save() *data.Task
}

type PaymentRule interface {
	Run() *data.Payment
	Save() *data.Payment
}

type PaymentManager interface {
	Configure()
	Execute()
}

type PaymentGroup struct {
	PaymentRuleList []PaymentRule
}

func NewPaymentGroup() PaymentGroup {
	return PaymentGroup{}
}
