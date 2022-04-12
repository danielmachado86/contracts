package data

import (
	"fmt"
	"math"
	"time"
)

type AgreementType int

func roundDateToNextDay(t time.Time) time.Time {
	return time.Date(
		t.Year(),
		t.Month(),
		t.Day()+1,
		0, 0, 0, 0,
		t.Location(),
	)
}

const (
	Rental AgreementType = iota
)

type PaymentRule interface {
	Calculate(pm *PaymentManager) []Payment
}
type TaskRule interface {
	Calculate(tm *TaskManager) Task
}

type SignatureDateRule struct {
}

func (sd *SignatureDateRule) Calculate(tm *TaskManager) Task {

	tn := "contract_signature"

	signDate := time.Now()

	task := Task{Name: tn, Date: signDate}
	return task
}

type StartDateRule struct {
}

func (sd *StartDateRule) Calculate(tm *TaskManager) Task {

	tn := "contract_start"

	offset := tm.Contract.Agreement.params["start_date_offset"]

	rounded := roundDateToNextDay(time.Now())
	startDate := rounded.AddDate(offset.years, offset.months, offset.days)

	task := Task{Name: tn, Date: startDate}
	return task
}

type TerminationDateRule struct {
}

func (ed *TerminationDateRule) Calculate(tm *TaskManager) Task {

	tn := "contract_termination"
	// tm.log.Info("Creating task: ", tn)

	period := tm.Contract.Duration

	startDateRule := &StartDateRule{}
	startDateTask := startDateRule.Calculate(tm)
	startDate := startDateTask.Date

	terminationDate := startDate.AddDate(period.years, period.months, period.days)

	task := Task{Name: tn, Date: terminationDate}
	return task
}

type AdvanceNoticeDeadlineRule struct {
}

func (an *AdvanceNoticeDeadlineRule) Calculate(tm *TaskManager) Task {

	tn := "advance_notice_deadline"
	// tm.log.Info("Creating task: ", tn)

	period := tm.Contract.Agreement.params["advance_notice_period"]

	terminationDateRule := &TerminationDateRule{}
	endDateTask := terminationDateRule.Calculate(tm)
	endDate := endDateTask.Date

	advanceNoticeDeadline := endDate.AddDate(-period.years, -period.months, -period.days)

	task := Task{Name: tn, Date: advanceNoticeDeadline}
	return task
}

type PaymentDeadlineRule struct {
	payment int
}

func (sd *PaymentDeadlineRule) Calculate(tm *TaskManager) Task {

	tn := fmt.Sprintf("payment %d", sd.payment)
	// tm.log.Info("Creating task: ", tn)

	params := tm.Contract.Agreement.params
	offset := params["start_date_offset"]

	rounded := roundDateToNextDay(time.Now())
	startDate := rounded.AddDate(offset.years, offset.months, offset.days)
	paymentDate := startDate.AddDate(0, params["payment_period"].months*sd.payment, 0)

	task := Task{Name: tn, Date: paymentDate}
	return task
}

type PaymentValueRule struct {
}

func (an *PaymentValueRule) Calculate(pm *PaymentManager) []Payment {

	cValue := pm.Contract.Price
	params := pm.Contract.Agreement.params

	// Contract duration
	cDuration := float64(pm.Contract.Duration.months)

	// Period: Time between payments
	pPeriod := float64(params["payment_period"].months)

	// Payment per period unit
	pValue := cValue / cDuration

	// Number of payments calculated, using contract duration and period
	var pNumber float64 = cDuration / pPeriod

	// Number of payments rounded to lowest interger
	pNumFloor := math.Floor(pNumber)

	// Accumulated value of complete payments
	pValWOResidue := pValue * pNumFloor * pPeriod

	// Value of partial payment
	residue := cValue - pValWOResidue

	var paymentList []Payment
	tm := &TaskManager{Contract: pm.Contract}
	for i := 0; i <= int(pNumFloor)-1; i++ {
		rule := PaymentDeadlineRule{payment: i}
		task := rule.Calculate(tm)
		payment := Payment{
			Name:  "Payment",
			Value: pValue * pPeriod,
			Date:  task.Date,
		}
		paymentList = append(paymentList, payment)
	}
	if residue > 0 {
		rule := PaymentDeadlineRule{payment: int(pNumFloor)}
		task := rule.Calculate(tm)
		paymentList = append(paymentList, Payment{Name: "Payment", Value: residue, Date: task.Date})
	}

	return paymentList
}

// Defines agreement parameters
type Agreement struct {
	Name   string
	Type   AgreementType
	params map[string]Date
}

type Date struct {
	days   int
	months int
	years  int
}

// Defines contract structure
type Contract struct {
	Duration  Date
	Price     float64
	Agreement Agreement
}

type Task struct {
	Name string
	Date time.Time
}

func (t Task) saveTask(tm *TaskManager) {
	tm.Tasks = append(tm.Tasks, t)
}

// Defines activities and schedules
type TaskManager struct {
	Contract Contract
	Tasks    []Task
}

type Payment struct {
	Name   string
	Remain float64
	Value  float64
	Date   time.Time
}

func (p Payment) savePayment(pm *PaymentManager) {
	pm.Payments = append(pm.Payments, p)
}

// Defines payments and schedules
type PaymentManager struct {
	Contract Contract
	Payments []Payment
	Schedule []time.Time
}

type Dashboard struct {
	Contract Contract
}

type Contracts []*Contract
type TaskManagers []*TaskManager
type PaymentManagers []*PaymentManager
