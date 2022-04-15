package handlers

// Defines activities and schedules
type TaskManager struct {
}

func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

var tmList []*TaskManager

type TaskManagers []*TaskManager
