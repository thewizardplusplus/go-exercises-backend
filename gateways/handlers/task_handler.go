package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"github.com/thewizardplusplus/go-exercises-backend/entities"
	"gorm.io/gorm"
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
	var task entities.Task
	if err := json.NewDecoder(request.Body).Decode(&task); err != nil {
		err = errors.Wrap(err, "unable to decode the request body")
		http.Error(writer, err.Error(), http.StatusBadRequest)

		return
	}

	id, err := handler.TaskStorage.CreateTask(task)
	if err != nil {
		err = errors.Wrap(err, "unable to create a task")
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return
	}

	idAsModel := entities.Task{Model: gorm.Model{ID: id}}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(idAsModel)
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
