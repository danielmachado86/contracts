package handlers

import "github.com/danielmachado86/contracts/dashboard/data"

// Defines activities and schedules
type TaskManager struct {
	Tasks []*data.Task
}

func NewTaskManager() *TaskManager {
	tm := &TaskManager{}
	tmList = append(tmList, tm)
	return tm
}

var tmList []*TaskManager

type TaskManagers []*TaskManager
