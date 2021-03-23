package handlers

import (
	"net/http"

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

// TaskHandler ...
type TaskHandler struct {
	TaskStorage TaskStorage
}

// GetTasks ...
func (handler TaskHandler) GetTasks(
	writer http.ResponseWriter,
	request *http.Request,
) {
}

// GetTask ...
func (handler TaskHandler) GetTask(
	writer http.ResponseWriter,
	request *http.Request,
) {
}

// CreateTask ...
func (handler TaskHandler) CreateTask(
	writer http.ResponseWriter,
	request *http.Request,
) {
}

// UpdateTask ...
func (handler TaskHandler) UpdateTask(
	writer http.ResponseWriter,
	request *http.Request,
) {
}

// DeleteTask ...
func (handler TaskHandler) DeleteTask(
	writer http.ResponseWriter,
	request *http.Request,
) {
}
