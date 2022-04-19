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

var ScheduleRules = map[string]ScheduleRule{
	"signature_date":          SignatureDate{},
	"start_date":              StartDate{},
	"end_date":                EndDate{},
	"advance_notice_deadline": AdvanceNoticeDeadline{},
}

var PaymentRules = map[string]PaymentRule{
	"periodic_payment": PeriodicPayment{},
	"termination":      Termination{},
}
