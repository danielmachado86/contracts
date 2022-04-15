package handlers

import (
	"fmt"
	"testing"
	"time"

	"github.com/danielmachado86/contracts/dashboard/data"
	"github.com/danielmachado86/contracts/dashboard/utils"
)

func TestStartDateRule(t *testing.T) {
	var a = &data.Agreement{
		Params: map[string]*utils.Date{
			"start_date_offset": {Days: 0},
		},
	}

	var c = &data.Contract{
		Duration:  &utils.Date{Months: 12},
		Agreement: a,
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
	t_task := &data.Task{Date: t_startDate}

	tm := NewTaskManager()
	r := StartDateRule{}
	task := r.Calculate(tm, c)

	if task.Date != t_task.Date {
		t.Fail()
	}
}

func TestEndDateRule(t *testing.T) {
	var a = &data.Agreement{
		Params: map[string]*utils.Date{
			"start_date_offset": {Days: 0},
		},
	}

	var c = &data.Contract{
		Duration:  &utils.Date{Months: 12},
		Agreement: a,
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
	t_task := &data.Task{Date: t_endDate}

	tm := NewTaskManager()
	rule := TerminationDateRule{}
	task := rule.Calculate(tm, c)

	if task.Date != t_task.Date {
		t.Fail()
	}
}

func TestAdvanceNoticeRule(t *testing.T) {
	var a = &data.Agreement{
		Params: map[string]*utils.Date{
			"start_date_offset":     {Days: 0},
			"advance_notice_period": {Months: 3},
		},
	}

	var c = &data.Contract{
		Duration:  &utils.Date{Months: 12},
		Agreement: a,
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
	t_task := data.Task{Date: t_advanceNoticeDeadlineDate}
	fmt.Printf("Date: %s \n", t_advanceNoticeDeadlineDate)

	tm := NewTaskManager()
	rule := AdvanceNoticeDeadlineRule{}
	task := rule.Calculate(tm, c)
	fmt.Printf("Date: %s", task.Date)

	if task.Date != t_task.Date {
		t.Fail()
	}
}

func TestPaymentQuantity(t *testing.T) {
	var a = &data.Agreement{
		Params: map[string]*utils.Date{
			"payment_period": {Months: 2},
		},
	}

	var c = &data.Contract{
		Duration:  &utils.Date{Months: 13},
		Price:     13,
		Agreement: a,
	}

	pm := NewPaymentManager()
	qty := pm.PaymentQuantity(c)
	tQty := 7

	if qty != tQty {
		t.Error(fmt.Printf("The quantity of payments: %d is differerent to: %d", qty, tQty))
	}

}

func TestSavePayment(t *testing.T) {
	var a = &data.Agreement{
		Params: map[string]*utils.Date{
			"payment_period": {Months: 5},
		},
	}

	var c = &data.Contract{
		Duration:  &utils.Date{Months: 12},
		Price:     24,
		Agreement: a,
	}

	pm := NewPaymentManager()

	r1 := &PaymentValueRule{last: false}
	p1 := r1.Calculate(pm, c)

	if int(p1.Value) != 10 {
		t.Error(fmt.Printf("The value of payment: %d is differerent to: %d", int(p1.Value), 5))
	}

	r2 := &PaymentValueRule{last: true}
	p2 := r2.Calculate(pm, c)

	if int(p2.Value) != 4 {
		t.Error(fmt.Printf("The value of payment: %d is differerent to: %d", int(p2.Value), 8))
	}
}
