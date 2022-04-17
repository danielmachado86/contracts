package handlers

import (
	"time"

	"github.com/danielmachado86/contracts/dashboard/data"
)

type ScheduleRuleManager interface {
	Run(*data.Contract) Scheduler
}
type PaymentRuleManager interface {
	Run(*data.Contract) Calculator
}

type Scheduler interface {
	GetDate() time.Time
}

type Calculator interface {
	GetValue() float64
}
