package handlers

import (
	"github.com/thewizardplusplus/go-exercises-backend/entities"
)

// TaskStorage ...
type TaskStorage interface {
	GetTasks() ([]entities.Task, error)
	GetTask(id uint) (entities.Task, error)
	CreateTask(task entities.Task) (id uint, err error)
	UpdateTask(id uint, task entities.Task) error
	DeleteTask(id uint) error
}
