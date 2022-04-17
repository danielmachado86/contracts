package handlers

import (
	"fmt"
	"time"

	"github.com/danielmachado86/contracts/dashboard/data"
	"github.com/danielmachado86/contracts/dashboard/utils"
)

func roundDateToNextDay(s Scheduler) time.Time {
	t := s.GetDate()
	return time.Date(
		t.Year(),
		t.Month(),
		t.Day()+1,
		0, 0, 0, 0,
		t.Location(),
	)
}

type SignatureDate struct {
	time time.Time
}

func createTask(n string, d time.Time) *data.Task {
	return &data.Task{Name: n, Date: d}
}

func addDate(r ScheduleRuleManager, p *utils.Period, c *data.Contract) time.Time {
	// Start date
	d := r.Run(c).GetDate()
	// Termination date
	return d.AddDate(p.Years, p.Months, p.Days)
}

func (r SignatureDate) Run(c *data.Contract) Scheduler {
	return createTask("signature_date", r.time)
}

type StartDate struct {
	time   time.Time
	offset *utils.Period
}

func (r StartDate) Run(c *data.Contract) Scheduler {

	sr := &SignatureDate{time: r.time}
	o := r.offset

	rounded := roundDateToNextDay(sr.Run(c))
	sd := rounded.AddDate(o.Years, o.Months, o.Days)

	return createTask("start_date", sd)
}

type EndDate struct {
	time   time.Time
	offset *utils.Period
}

func (r EndDate) Run(c *data.Contract) Scheduler {

	d := c.Duration

	// Start rule
	sr := &StartDate{
		offset: r.offset,
		time:   r.time,
	}
	// Termination date
	td := addDate(sr, d, c)

	return createTask("end_date", td)
}

type AdvanceNoticeDeadline struct {
	time   time.Time
	offset *utils.Period
	period *utils.Period
}

func (r AdvanceNoticeDeadline) Run(c *data.Contract) Scheduler {

	// Period
	p := r.period

	// End date rule
	er := &EndDate{offset: r.offset, time: r.time}
	// End date
	ed := er.Run(c).GetDate()
	// Advance notice deadline
	nd := ed.AddDate(-p.Years, -p.Months, -p.Days)

	return createTask("advance_notice_deadline", nd)
}

type PaymentDeadline struct {
	offset  *utils.Period
	time    time.Time
	period  *utils.Period
	payment int
}

func (r PaymentDeadline) Run(c *data.Contract) Scheduler {

	sr := &StartDate{offset: r.offset, time: r.time}
	//Payment deadline date
	pd := addDate(sr, r.period, c)

	return createTask(fmt.Sprintf("payment %d", r.payment), pd)
}
