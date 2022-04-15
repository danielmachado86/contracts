package handlers

import (
	"fmt"
	"time"

	"github.com/danielmachado86/contracts/dashboard/data"
)

func roundDateToNextDay(t time.Time) time.Time {
	return time.Date(
		t.Year(),
		t.Month(),
		t.Day()+1,
		0, 0, 0, 0,
		t.Location(),
	)
}

type TaskRule interface {
	Calculate(tm *TaskManager, c *data.Contract) data.Task
}

type SignatureDateRule struct {
}

func (tr *SignatureDateRule) Calculate(tm *TaskManager, c *data.Contract) *data.Task {

	tn := "contract_signature"

	signDate := time.Now()

	task := &data.Task{Name: tn, Date: signDate}
	return task
}

type StartDateRule struct {
}

func (tr *StartDateRule) Calculate(tm *TaskManager, c *data.Contract) *data.Task {

	offset := c.Agreement.Params["start_date_offset"]

	rounded := roundDateToNextDay(time.Now())
	startDate := rounded.AddDate(offset.Years, offset.Months, offset.Days)

	tn := "contract_start"
	task := &data.Task{Name: tn, Date: startDate}
	return task
}

type TerminationDateRule struct {
}

func (tr *TerminationDateRule) Calculate(tm *TaskManager, c *data.Contract) *data.Task {

	tn := "contract_termination"

	cDuraion := c.Duration

	startDateRule := &StartDateRule{}
	startDateTask := startDateRule.Calculate(tm, c)
	startDate := startDateTask.Date

	terminationDate := startDate.AddDate(cDuraion.Years, cDuraion.Months, cDuraion.Days)

	task := &data.Task{Name: tn, Date: terminationDate}
	return task
}

type AdvanceNoticeDeadlineRule struct {
}

func (tr *AdvanceNoticeDeadlineRule) Calculate(tm *TaskManager, c *data.Contract) *data.Task {

	tn := "advance_notice_deadline"

	period := c.Agreement.Params["advance_notice_period"]

	terminationDateRule := &TerminationDateRule{}
	endDateTask := terminationDateRule.Calculate(tm, c)
	endDate := endDateTask.Date

	advanceNoticeDeadline := endDate.AddDate(-period.Years, -period.Months, -period.Days)

	task := &data.Task{Name: tn, Date: advanceNoticeDeadline}
	return task
}

type PaymentDeadlineRule struct {
	payment int
}

func (tr *PaymentDeadlineRule) Calculate(tm *TaskManager, c *data.Contract) *data.Task {

	tn := fmt.Sprintf("payment %d", tr.payment)

	params := c.Agreement.Params
	offset := params["start_date_offset"]

	rounded := roundDateToNextDay(time.Now())
	startDate := rounded.AddDate(offset.Years, offset.Months, offset.Days)
	paymentDate := startDate.AddDate(0, params["payment_period"].Months*tr.payment, 0)

	task := &data.Task{Name: tn, Date: paymentDate}
	return task
}
