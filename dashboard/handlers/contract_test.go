package handlers

import (
	"fmt"
	"testing"
	"time"

	"github.com/danielmachado86/contracts/dashboard/data/"
	"github.com/danielmachado86/contracts/dashboard/utils"
)

func TestStartDateRule(t *testing.T) {
	var a = &data.Agreement{
		Params: map[string]*utils.Date{
			"start_date_offset": &utils.Date{Days: 0},
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
	var a = &data.Agreement{}

	var c = &data.Contract{
		Duration:  utils.Date{Months: 12},
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
		Params: map[string]utils.Date{
			"advance_notice_period": {Months: 3},
		},
	}

	var c = &data.Contract{
		Duration:  utils.Date{Months: 12},
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

func TestSaveTask(t *testing.T) {
	var a = &data.Agreement{
		Params: map[string]utils.Date{
			"start_date_offset":     {Days: 0},
			"advance_notice_period": {Months: 3},
			"payment_period":        {Months: 2},
		},
	}

	var c = &data.Contract{
		Duration:  utils.Date{Months: 12},
		Agreement: *a,
	}

	tm := NewTaskManager()

	r1 := &SignatureDateRule{}
	t1 := r1.Calculate(tm, c)
	t1.save(tm)

	r2 := &StartDateRule{}
	t2 := r2.Calculate(tm, c)
	t2.save(tm)

	r3 := &AdvanceNoticeDeadlineRule{}
	t3 := r3.Calculate(tm, c)
	t3.save(tm)

	if len(tm.Tasks) != 3 {
		t.Error(fmt.Printf("Number of tasks is not 3, result: %d", len(tm.Tasks)))
	}

	if tm.Tasks[len(tm.Tasks)-1].Name != "advance_notice_deadline" {
		t.Error(fmt.Printf("Name of task is not advance_notice_deadline, result: %s", tm.Tasks[len(tm.Tasks)-1].Name))
	}
}

func TestPaymentQuantity(t *testing.T) {
	var a = &data.Agreement{
		Params: map[string]utils.Date{
			"payment_period": {Months: 2},
		},
	}

	var c = &data.Contract{
		Duration:  utils.Date{Months: 13},
		Price:     13,
		Agreement: *a,
	}

	pm := NewPaymentManager()
	qty := pm.PaymentQuantity(c)
	tQty := 7

	if qty != tQty {
		t.Error(fmt.Printf("The quantity of payments: %d is differerent to: %d", qty, tQty))
	}

}

func TestSavePayment(t *testing.T) {
	var a = &Agreement{
		Params: map[string]Date{
			"payment_period": {Months: 5},
		},
	}

	var c = &Contract{
		Duration:  Date{Months: 12},
		Price:     24,
		Agreement: *a,
	}

	pm := NewPaymentManager()

	r1 := &PaymentValueRule{last: false}
	p1 := r1.Calculate(pm, c)
	p1.save(pm)

	if int(p1.Value) != 10 {
		t.Error(fmt.Printf("The value of payment: %d is differerent to: %d", int(pm.Payments[0].Value), 5))
	}

	r2 := &PaymentValueRule{last: true}
	p2 := r2.Calculate(pm, c)
	p2.save(pm)

	if int(p2.Value) != 4 {
		t.Error(fmt.Printf("The value of payment: %d is differerent to: %d", int(pm.Payments[1].Value), 8))
	}
}
