package data

import "time"

type AgreementType int

const (
	Rental AgreementType = iota
)

// Defines agreement parameters
type AgreementManager struct {
	Name string
	Type AgreementType
}

// Defines contract structure
type Contract struct {
	ID        int
	Duration  time.Duration
	Price     float64
	Agreement AgreementManager
}

// Defines key dates
type Scheduler struct {
	Schedule []time.Time
}

type Task struct {
	Name string
}

// Defines activities and schedules
type TaskManager struct {
	ID       int
	Contract Contract
	Task     []Task
	Schedule Scheduler
}

func AddTaskManager(tm *TaskManager) {
	taskManagerList = append(taskManagerList, tm)
}

type Payment struct {
	Name  string
	Value float64
}

func AddPaymentManager(pm *PaymentManager) {
	paymentManagerList = append(paymentManagerList, pm)
}

// Defines payments and schedules
type PaymentManager struct {
	ID       int
	Contract Contract
	Payment  []Payment
	Schedule Scheduler
}

type Dashboard struct {
	Contract Contract
	Tasks    TaskManagers
	Payments PaymentManagers
}

func GetDashboard(c *Contract) (*Dashboard, error) {

	dashboard := &Dashboard{
		Contract: *c,
		Tasks:    taskManagerList,
		Payments: paymentManagerList,
	}

	return dashboard, nil
}

type Contracts []*Contract
type TaskManagers []*TaskManager
type PaymentManagers []*PaymentManager

var contractList = []*Contract{
	{
		ID:       1,
		Duration: 31536000000000000,
		Price:    24000000,
		Agreement: AgreementManager{
			Name: "Rental property",
			Type: Rental,
		},
	},
}

var taskManagerList = []*TaskManager{
	{
		ID:       1,
		Contract: *contractList[0],
		Task: []Task{
			{Name: "Contract signing"},
			{Name: "Advance notice"},
			{Name: "End"},
		},
		Schedule: Scheduler{
			Schedule: []time.Time{
				time.Now(),
				time.Now().AddDate(0, 8, 0),
				time.Now().AddDate(1, 0, 0),
			},
		},
	},
}

var paymentManagerList = []*PaymentManager{
	{
		ID:       1,
		Contract: *contractList[0],
		Payment: []Payment{
			{
				Name:  "Payment 1",
				Value: 2000000,
			},
			{
				Name:  "Payment 2",
				Value: 2000000,
			},
			{
				Name:  "Payment 3",
				Value: 2000000,
			},
			{
				Name:  "Payment 4",
				Value: 2000000,
			},
			{
				Name:  "Payment 5",
				Value: 2000000,
			},
			{
				Name:  "Payment 6",
				Value: 2000000,
			},
			{
				Name:  "Payment 7",
				Value: 2000000,
			},
			{
				Name:  "Payment 8",
				Value: 2000000,
			},
			{
				Name:  "Payment 9",
				Value: 2000000,
			},
			{
				Name:  "Payment 10",
				Value: 2000000,
			},
			{
				Name:  "Payment 11",
				Value: 2000000,
			},
			{
				Name:  "Payment 12",
				Value: 2000000,
			},
		},
		Schedule: Scheduler{
			Schedule: []time.Time{
				time.Now().AddDate(0, 0, 5),
				time.Now().AddDate(0, 1, 5),
				time.Now().AddDate(0, 2, 5),
				time.Now().AddDate(0, 3, 5),
				time.Now().AddDate(0, 4, 5),
				time.Now().AddDate(0, 5, 5),
				time.Now().AddDate(0, 6, 5),
				time.Now().AddDate(0, 7, 5),
				time.Now().AddDate(0, 8, 5),
				time.Now().AddDate(0, 9, 5),
				time.Now().AddDate(0, 10, 5),
				time.Now().AddDate(0, 11, 5),
			},
		},
	},
}
