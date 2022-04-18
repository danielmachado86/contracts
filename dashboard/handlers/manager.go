package handlers

import (
	"github.com/danielmachado86/contracts/dashboard/data"
)

type ScheduleRuleManager interface {
	Run() *data.Task
	Save() *data.Task
}

type PaymentRuleManager interface {
	Run() *data.Payment
	Save() *data.Payment
}

var ScheduleRules = map[string]ScheduleRuleManager{
	"signature_date":          SignatureDate{},
	"start_date":              StartDate{},
	"end_date":                EndDate{},
	"advance_notice_deadline": AdvanceNoticeDeadline{},
}

var PaymentRules = map[string]PaymentRuleManager{
	"periodic_payment": PeriodicPaymentValue{},
	"termination":      TerminationValue{},
}
