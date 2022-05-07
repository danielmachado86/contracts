package rules

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

func createTask(n string, d time.Time) *data.Task {
	t := &data.Task{Name: n, Date: d}
	return t
}

// Define las condiciones para calcular la fecha de firma de contrato
type SignatureDate struct {
	task    *data.Task
	parties []*data.Party
}

// Calcula fecha de firma de contrato
func (r *SignatureDate) Compute() Rule {
	t := time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC)
	for _, p := range r.parties {
		d := data.GetSignature(p).Date
		if d.After(t) {
			t = d
		}
	}
	r.task = createTask("signature_date", t)
	return r
}

func (r *SignatureDate) Save() {
	r.task.Save()
}

type StartDate struct {
	parties []*data.Party
	task    *data.Task
}

func (r *StartDate) Compute() Rule {

	params := data.GetParams()

	sr := &SignatureDate{parties: r.parties}
	o := params.Offset

	sr.Compute()
	rounded := roundDateToNextDay(sr.task.GetDate())
	sd := rounded.AddDate(o.Years, o.Months, o.Days)
	r.task = createTask("start_date", sd)
	return r
}

func (r *StartDate) Save() {
	r.task.Save()
}

type EndDate struct {
	parties []*data.Party
	task    *data.Task
}

func (r *EndDate) Compute() Rule {

	params := data.GetParams()

	d := params.Duration

	// Start rule
	sr := &StartDate{
		parties: r.parties,
	}
	// Termination date
	sr.Compute()
	td := sr.task.AddPeriod(d).Add(-time.Second * 1)

	r.task = createTask("end_date", td)
	return r
}

func (r *EndDate) Save() {
	r.task.Save()
}

type AdvanceNoticeDeadline struct {
	parties []*data.Party
	task    *data.Task
}

func (r *AdvanceNoticeDeadline) Compute() Rule {

	params := data.GetParams()

	// Period
	p := params.PeriodAN

	// End date rule
	er := &EndDate{parties: r.parties}
	er.Compute()
	// End date
	ed := er.task.GetDate()
	// Advance notice deadline
	nd := ed.AddDate(-p.Years, -p.Months, -p.Days)

	r.task = createTask("advance_notice_deadline", nd)
	return r
}

func (r *AdvanceNoticeDeadline) Save() {
	r.task.Save()
}

type NotificationDate struct {
	parties []*data.Party
	payment int
	task    *data.Task
}

func (r *NotificationDate) Compute() Rule {

	return r
}

func (r *NotificationDate) Save() {
	r.task.Save()
}

type PaymentDueDate struct {
	name string
	time time.Time
	task *data.Task
}

func (r *PaymentDueDate) Compute() Rule {

	//Payment closing date
	pd := r.time.AddDate(0, 0, 5)
	r.task = createTask(fmt.Sprintf("%s_due_date", r.name), pd)
	return r
}

func (r *PaymentDueDate) Save() {
	r.task.Save()
}
