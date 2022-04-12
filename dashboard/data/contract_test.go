package data

import (
	"fmt"
	"testing"
	"time"
)

func TestStartDateRule(t *testing.T) {
	var agreement = &Agreement{
		params: map[string]Date{
			"start_date_offset": {days: 0},
		},
	}

	var contract = &Contract{
		Duration:  Date{months: 12},
		Agreement: *agreement,
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

	tm := &TaskManager{Contract: *contract}
	rule := StartDateRule{}
	task := rule.Calculate(tm)

	if task.Date != t_task.Date {
		t.Fail()
	}
}

func TestEndDateRule(t *testing.T) {
	var agreement = &Agreement{}

	var contract = &Contract{
		Duration:  Date{months: 12},
		Agreement: *agreement,
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

	tm := &TaskManager{Contract: *contract}
	rule := TerminationDateRule{}
	task := rule.Calculate(tm)

	if task.Date != t_task.Date {
		t.Fail()
	}
}

func TestAdvanceNoticeRule(t *testing.T) {
	var agreement = &Agreement{
		params: map[string]Date{
			"advance_notice_period": {months: 3},
		},
	}

	var contract = &Contract{
		Duration:  Date{months: 12},
		Agreement: *agreement,
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

	tm := &TaskManager{Contract: *contract}
	rule := AdvanceNoticeDeadlineRule{}
	task := rule.Calculate(tm)
	fmt.Printf("Date: %s", task.Date)

	if task.Date != t_task.Date {
		t.Fail()
	}
}

func TestPaymentValueRule(t *testing.T) {
	var agreement = &Agreement{
		params: map[string]Date{
			"payment_period": {months: 2},
		},
	}

	var contract = &Contract{
		Duration:  Date{months: 13},
		Price:     13,
		Agreement: *agreement,
	}

	pm := &PaymentManager{Contract: *contract}
	rule := PaymentValueRule{}
	payments := rule.Calculate(pm)

	if len(payments) != 7 {
		t.Fail()
	}

	if payments[0].Value != float64(2) {
		t.Fail()
	}

	if payments[len(payments)-1].Value != float64(1) {
		t.Fail()
	}
}

func TestSaveTaskRules(t *testing.T) {
	var agreement = &Agreement{
		params: map[string]Date{
			"start_date_offset":     {days: 0},
			"advance_notice_period": {months: 3},
			"payment_period":        {months: 2},
		},
	}

	var contract = &Contract{
		Duration:  Date{months: 12},
		Agreement: *agreement,
	}

	tm := &TaskManager{Contract: *contract}

	signDateRule := &SignatureDateRule{}
	signDateTask := signDateRule.Calculate(tm)
	signDateTask.saveTask(tm)

	startDateRule := &StartDateRule{}
	startDateTask := startDateRule.Calculate(tm)
	startDateTask.saveTask(tm)

	advanceNoticeDeadlineRule := &AdvanceNoticeDeadlineRule{}
	advanceNoticeTask := advanceNoticeDeadlineRule.Calculate(tm)
	advanceNoticeTask.saveTask(tm)

	if len(tm.Tasks) != 3 {
		t.Fail()
	}

	if tm.Tasks[len(tm.Tasks)-1].Name != "advance_notice_deadline" {
		t.Fail()
	}
}
