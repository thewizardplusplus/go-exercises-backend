package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-log/log"
	"github.com/gorilla/mux"
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
	Logger      log.Logger
}

// GetTasks ...
func (handler TaskHandler) GetTasks(
	writer http.ResponseWriter,
	request *http.Request,
) {
	tasks, err := handler.TaskStorage.GetTasks()
	if err != nil {
		err = errors.Wrap(err, "[error] unable to get the tasks")
		handler.Logger.Log(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(tasks) // nolint: gosec, errcheck
}

// GetTask ...
func (handler TaskHandler) GetTask(
	writer http.ResponseWriter,
	request *http.Request,
) {
	idAsStr := mux.Vars(request)["id"]
	id, err := strconv.ParseUint(idAsStr, 10, 64)
	if err != nil {
		err = errors.Wrap(err, "[error] unable to decode the task ID")
		handler.Logger.Log(err)
		http.Error(writer, err.Error(), http.StatusBadRequest)

		return
	}

	task, err := handler.TaskStorage.GetTask(uint(id))
	if err != nil {
		err = errors.Wrap(err, "[error] unable to get the task")
		handler.Logger.Log(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(task) // nolint: gosec, errcheck
}

// CreateTask ...
func (handler TaskHandler) CreateTask(
	writer http.ResponseWriter,
	request *http.Request,
) {
	var task entities.Task
	if err := json.NewDecoder(request.Body).Decode(&task); err != nil {
		err = errors.Wrap(err, "[error] unable to decode the request body")
		handler.Logger.Log(err)
		http.Error(writer, err.Error(), http.StatusBadRequest)

		return
	}

	id, err := handler.TaskStorage.CreateTask(task)
	if err != nil {
		err = errors.Wrap(err, "[error] unable to create a task")
		handler.Logger.Log(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return
	}

	idAsModel := entities.Task{Model: gorm.Model{ID: id}}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(idAsModel) // nolint: gosec, errcheck
}

// UpdateTask ...
func (handler TaskHandler) UpdateTask(
	writer http.ResponseWriter,
	request *http.Request,
) {
	idAsStr := mux.Vars(request)["id"]
	id, err := strconv.ParseUint(idAsStr, 10, 64)
	if err != nil {
		err = errors.Wrap(err, "[error] unable to decode the task ID")
		handler.Logger.Log(err)
		http.Error(writer, err.Error(), http.StatusBadRequest)

		return
	}

	var task entities.Task
	if err := json.NewDecoder(request.Body).Decode(&task); err != nil {
		err = errors.Wrap(err, "[error] unable to decode the request body")
		handler.Logger.Log(err)
		http.Error(writer, err.Error(), http.StatusBadRequest)

		return
	}

	if err := handler.TaskStorage.UpdateTask(uint(id), task); err != nil {
		err = errors.Wrap(err, "[error] unable to update a task")
		handler.Logger.Log(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return
	}
}

// DeleteTask ...
func (handler TaskHandler) DeleteTask(
	writer http.ResponseWriter,
	request *http.Request,
) {
	idAsStr := mux.Vars(request)["id"]
	id, err := strconv.ParseUint(idAsStr, 10, 64)
	if err != nil {
		err = errors.Wrap(err, "[error] unable to decode the task ID")
		handler.Logger.Log(err)
		http.Error(writer, err.Error(), http.StatusBadRequest)

		return
	}

	if err := handler.TaskStorage.DeleteTask(uint(id)); err != nil {
		err = errors.Wrap(err, "[error] unable to delete a task")
		handler.Logger.Log(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return
	}
}
