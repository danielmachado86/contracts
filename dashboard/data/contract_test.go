package data

import (
	"fmt"
	"testing"
	"time"
)

func TestStartDateRule(t *testing.T) {
	var a = &Agreement{
		params: map[string]Date{
			"start_date_offset": {days: 0},
		},
	}

	var c = &Contract{
		Duration:  Date{months: 12},
		Agreement: *a,
	}
	tt := time.Now()
	t_roundedToNextDay := time.Date(
		tt.Year(),
		tt.Month(),
		tt.Day()+1,
		0, 0, 0, 0,
		tt.Location(),
	)
	t_startDate := t_roundedToNextDay.AddDate(0, 0, 0)
	t_task := Task{Date: t_startDate}

	tm := NewTaskManager()
	r := StartDateRule{}
	task := r.Calculate(tm, c)

	if task.Date != t_task.Date {
		t.Fail()
	}
}

func TestEndDateRule(t *testing.T) {
	var a = &Agreement{}

	var c = &Contract{
		Duration:  Date{months: 12},
		Agreement: *a,
	}

	tt := time.Now()
	t_roundedToNextDay := time.Date(
		tt.Year(),
		tt.Month(),
		tt.Day()+1,
		0, 0, 0, 0,
		tt.Location(),
	)
	t_startDate := t_roundedToNextDay.AddDate(0, 0, 0)
	t_endDate := t_startDate.AddDate(0, 12, 0)
	t_task := Task{Date: t_endDate}

	tm := NewTaskManager()
	rule := TerminationDateRule{}
	task := rule.Calculate(tm, c)

	if task.Date != t_task.Date {
		t.Fail()
	}
}

func TestAdvanceNoticeRule(t *testing.T) {
	var a = &Agreement{
		params: map[string]Date{
			"advance_notice_period": {months: 3},
		},
	}

	var c = &Contract{
		Duration:  Date{months: 12},
		Agreement: *a,
	}
	tt := time.Now()
	t_roundedToNextDay := time.Date(
		tt.Year(),
		tt.Month(),
		tt.Day()+1,
		0, 0, 0, 0,
		tt.Location(),
	)
	t_startDate := t_roundedToNextDay.AddDate(0, 0, 0)
	t_endDate := t_startDate.AddDate(0, 12, 0)
	t_advanceNoticeDeadlineDate := t_endDate.AddDate(0, -3, 0)
	t_task := Task{Date: t_advanceNoticeDeadlineDate}
	fmt.Printf("Date: %s \n", t_advanceNoticeDeadlineDate)

	tm := NewTaskManager()
	rule := AdvanceNoticeDeadlineRule{}
	task := rule.Calculate(tm, c)
	fmt.Printf("Date: %s", task.Date)

	if task.Date != t_task.Date {
		t.Fail()
	}
}

func TestPaymentValueRule(t *testing.T) {
	var a = &Agreement{
		params: map[string]Date{
			"payment_period": {months: 2},
		},
	}

	var c = &Contract{
		Duration:  Date{months: 13},
		Price:     13,
		Agreement: *a,
	}

	pm := NewPaymentManager()
	rule := PaymentValueRule{}
	rule.Calculate(pm, c)

	if len(pm.Payments) != 7 {
		t.Fail()
	}

	if pm.Payments[0].Value != float64(2) {
		t.Fail()
	}

	if pm.Payments[len(pm.Payments)-1].Value != float64(1) {
		t.Fail()
	}
}

func TestSaveTask(t *testing.T) {
	var a = &Agreement{
		params: map[string]Date{
			"start_date_offset":     {days: 0},
			"advance_notice_period": {months: 3},
			"payment_period":        {months: 2},
		},
	}

	var c = &Contract{
		Duration:  Date{months: 12},
		Agreement: *a,
	}

	tm := NewTaskManager()

	signDate := &SignatureDateRule{}
	signDate.Calculate(tm, c)

	startDate := &StartDateRule{}
	startDate.Calculate(tm, c)

	advanceNoticeDeadlineRule := &AdvanceNoticeDeadlineRule{}
	advanceNoticeDeadlineRule.Calculate(tm, c)

	if len(tm.Tasks) != 3 {
		t.Fail()
	}

	if tm.Tasks[len(tm.Tasks)-1].Name != "advance_notice_deadline" {
		t.Fail()
	}
}

func TestSavePayment(t *testing.T) {
	var a = &Agreement{
		params: map[string]Date{
			"payment_period": {months: 2},
		},
	}

	var c = &Contract{
		Duration:  Date{months: 13},
		Price:     13,
		Agreement: *a,
	}

	pm := NewPaymentManager()

	p1 := &PaymentValueRule{last: false}
	p1.Calculate(pm, c)

	p2 := &PaymentValueRule{last: true}
	p2.Calculate(pm, c)

	if len(pm.Payments) != 2 {
		t.Error(fmt.Printf("The quantity of payments: %d is differerent to: %d", len(pm.Payments), 2))
	}

	if int(pm.Payments[0].Value) != 2 {
		t.Error(fmt.Printf("The value of payment: %d is differerent to: %d", int(pm.Payments[0].Value), 2))
	}

	if int(pm.Payments[1].Value) != 1 {
		t.Error(fmt.Printf("The value of payment: %d is differerent to: %d", int(pm.Payments[1].Value), 1))
	}
}
