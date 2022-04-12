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
	Calculate(pm *PaymentManager, c *Contract) []Payment
}
type TaskRule interface {
	Calculate(tm *TaskManager, c *Contract) Task
}

type SignatureDateRule struct {
}

func (tr *SignatureDateRule) Calculate(tm *TaskManager, c *Contract) Task {

	tn := "contract_signature"

	signDate := time.Now()

	task := Task{Name: tn, Date: signDate}
	task.save(tm)
	return task
}

type StartDateRule struct {
}

func (tr *StartDateRule) Calculate(tm *TaskManager, c *Contract) Task {

	tn := "contract_start"

	offset := c.Agreement.params["start_date_offset"]

	rounded := roundDateToNextDay(time.Now())
	startDate := rounded.AddDate(offset.years, offset.months, offset.days)

	task := Task{Name: tn, Date: startDate}
	task.save(tm)
	return task
}

type TerminationDateRule struct {
}

func (tr *TerminationDateRule) Calculate(tm *TaskManager, c *Contract) Task {

	tn := "contract_termination"
	// tm.log.Info("Creating task: ", tn)

	cDuraion := c.Duration

	startDateRule := &StartDateRule{}
	startDateTask := startDateRule.Calculate(tm, c)
	startDate := startDateTask.Date

	terminationDate := startDate.AddDate(cDuraion.years, cDuraion.months, cDuraion.days)

	task := Task{Name: tn, Date: terminationDate}
	task.save(tm)
	return task
}

type AdvanceNoticeDeadlineRule struct {
}

func (tr *AdvanceNoticeDeadlineRule) Calculate(tm *TaskManager, c *Contract) Task {

	tn := "advance_notice_deadline"
	// tm.log.Info("Creating task: ", tn)

	period := c.Agreement.params["advance_notice_period"]

	terminationDateRule := &TerminationDateRule{}
	endDateTask := terminationDateRule.Calculate(tm, c)
	endDate := endDateTask.Date

	advanceNoticeDeadline := endDate.AddDate(-period.years, -period.months, -period.days)

	task := Task{Name: tn, Date: advanceNoticeDeadline}
	task.save(tm)
	return task
}

type PaymentDeadlineRule struct {
	payment int
}

func (tr *PaymentDeadlineRule) Calculate(tm *TaskManager, c *Contract) Task {

	tn := fmt.Sprintf("payment %d", tr.payment)

	params := c.Agreement.params
	offset := params["start_date_offset"]

	rounded := roundDateToNextDay(time.Now())
	startDate := rounded.AddDate(offset.years, offset.months, offset.days)
	paymentDate := startDate.AddDate(0, params["payment_period"].months*tr.payment, 0)

	task := Task{Name: tn, Date: paymentDate}
	task.save(tm)
	return task
}

type PaymentValueRule struct {
	last bool
}

func (pm *PaymentManager) PaymentQuantity(c *Contract) int {
	params := c.Agreement.params

	// Contract duration
	cDuration := float64(c.Duration.months)

	// Period: Time between payments
	pPeriod := float64(params["payment_period"].months)

	// Number of payments calculated, using contract duration and period
	var pNumber float64 = cDuration / pPeriod

	// Number of payments rounded to least higher interger
	pNumCeil := math.Ceil(pNumber)

	return int(pNumCeil)
}

func (pm *PaymentManager) LastPayment(c *Contract) float64 {
	return float64(int(c.Price) % (pm.PaymentQuantity(c) - 1))
}

func (pr *PaymentValueRule) Calculate(pm *PaymentManager, c *Contract) Payment {

	params := c.Agreement.params
	// Contract duration
	cDuration := float64(c.Duration.months)
	cValue := float64(c.Price)

	// Period: Time between payments
	pPeriod := float64(params["payment_period"].months)

	// Payment per period unit
	pValue := cValue / cDuration

	if pr.last {
		pValue = pm.LastPayment(c)
	}

	p := Payment{
		Name:  "Payment",
		Value: pValue * pPeriod,
	}
	p.save(pm)

	return p
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

func (t *Task) save(tm *TaskManager) {
	tm.Tasks = append(tm.Tasks, t)
}

// Defines activities and schedules
type TaskManager struct {
	Tasks []*Task
}

func NewTaskManager() *TaskManager {
	tm := &TaskManager{}
	tmList = append(tmList, tm)
	return tm
}

type Payment struct {
	Name   string
	Remain float64
	Value  float64
	Date   time.Time
}

func (p *Payment) save(pm *PaymentManager) {
	pm.Payments = append(pm.Payments, p)
}

// Defines payments and schedules
type PaymentManager struct {
	QtyPayments int
	Payments    []*Payment
	Schedule    []time.Time
}

func NewPaymentManager() *PaymentManager {
	pm := &PaymentManager{}
	pmList = append(pmList, pm)
	return pm
}

func (pm *PaymentManager) ListPayments(c *Contract) []*Payment {

	// Create new TaskManager to save the created payment tasks
	tm := NewTaskManager()

	pQty := pm.PaymentQuantity(c)

	for i := 1; i <= pQty; i++ {
		// Payment deadline rule
		tr := PaymentDeadlineRule{i}
		tr.Calculate(tm, c)

		// Last payment
		last := false
		if i == pQty {
			last = true
		}
		pr := PaymentValueRule{last}
		pr.Calculate(pm, c)
	}

	return pm.Payments
}

type Dashboard struct {
	Contract        Contract
	TaskManagers    TaskManagers
	PaymentManagers PaymentManagers
}

type Contracts []*Contract
type TaskManagers []*TaskManager
type PaymentManagers []*PaymentManager

var tmList []*TaskManager
var pmList []*PaymentManager
