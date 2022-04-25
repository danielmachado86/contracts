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

type SignatureDate struct {
	time time.Time
}

func (r SignatureDate) Run() *data.Task {
	return createTask("signature_date", r.time)
}

func (r SignatureDate) Save() *data.Task {
	return r.Run().Save()
}

type StartDate struct {
	time time.Time
}

func (r StartDate) Run() *data.Task {

	params := data.GetParams()

	sr := &SignatureDate{time: r.time}
	o := params.Offset

	rounded := roundDateToNextDay(sr.Run().GetDate())
	sd := rounded.AddDate(o.Years, o.Months, o.Days)

	return createTask("start_date", sd)
}

func (r StartDate) Save() *data.Task {
	return r.Run().Save()
}

type EndDate struct {
	time time.Time
}

func (r EndDate) Run() *data.Task {

	params := data.GetParams()

	d := params.Duration

	// Start rule
	sr := &StartDate{
		time: r.time,
	}
	// Termination date
	td := sr.Run().AddPeriod(d)

	return createTask("end_date", td)
}

func (r EndDate) Save() *data.Task {
	return r.Run().Save()
}

type AdvanceNoticeDeadline struct {
	time time.Time
}

func (r AdvanceNoticeDeadline) Run() *data.Task {

	params := data.GetParams()

	// Period
	p := params.PeriodAN

	// End date rule
	er := &EndDate{time: r.time}
	// End date
	ed := er.Run().GetDate()
	// Advance notice deadline
	nd := ed.AddDate(-p.Years, -p.Months, -p.Days)

	return createTask("advance_notice_deadline", nd)
}

func (r AdvanceNoticeDeadline) Save() *data.Task {
	return r.Run().Save()
}

type PeriodicPaymentClosingDate struct {
	time    time.Time
	payment int
}

func (r PeriodicPaymentClosingDate) Run() *data.Task {

	params := data.GetParams()

	sr := &StartDate{time: r.time}

	p := params.PaymentPeriod
	p.Months = p.Months * r.payment

	//Payment closing date
	pd := sr.Run().AddPeriod(p)

	return createTask(fmt.Sprintf("closing_date_%d", r.payment), pd)
}

func (r PeriodicPaymentClosingDate) Save() *data.Task {
	return r.Run().Save()
}

type PaymentDueDate struct {
	name string
	time time.Time
}

func (r PaymentDueDate) Run() *data.Task {

	//Payment closing date
	pd := r.time.AddDate(0, 0, 5)

	return createTask(fmt.Sprintf("%s_due_date", r.name), pd)
}

func (r PaymentDueDate) Save() *data.Task {
	return r.Run().Save()
}
