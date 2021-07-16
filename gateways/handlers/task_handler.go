package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-log/log"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/pkg/errors"
	"github.com/thewizardplusplus/go-exercises-backend/entities"
	httputils "github.com/thewizardplusplus/go-http-utils"
	"gorm.io/gorm"
)

// TaskStorage ...
type TaskStorage interface {
	entities.TaskGetter

	GetTasks(userID uint, pagination entities.Pagination) ([]entities.Task, error)
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
	var pagination entities.Pagination
	err := schema.NewDecoder().Decode(&pagination, request.URL.Query())
	if err != nil {
		err = errors.Wrap(err, "[error] unable to decode the pagination parameters")
		handler.Logger.Log(err)
		http.Error(writer, err.Error(), http.StatusBadRequest)

		return
	}

	user := request.Context().Value(userContextKey{}).(entities.User)
	tasks, err := handler.TaskStorage.GetTasks(user.ID, pagination)
	if err != nil {
		err = errors.Wrap(err, "[error] unable to get the tasks")
		handler.Logger.Log(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return
	}

	for index := range tasks {
		tasks[index].User.PasswordHash = ""
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

	user := request.Context().Value(userContextKey{}).(entities.User)
	task, err := handler.TaskStorage.GetTask(user.ID, uint(id))
	if err != nil {
		err = errors.Wrap(err, "[error] unable to get the task")
		handler.Logger.Log(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return
	}

	task.User.PasswordHash = ""

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(task) // nolint: gosec, errcheck
}

// CreateTask ...
func (handler TaskHandler) CreateTask(
	writer http.ResponseWriter,
	request *http.Request,
) {
	var task entities.Task
	if err := httputils.ReadJSON(request.Body, &task); err != nil {
		err = errors.Wrap(err, "[error] unable to decode the task data")
		handler.Logger.Log(err)
		http.Error(writer, err.Error(), http.StatusBadRequest)

		return
	}

	user := request.Context().Value(userContextKey{}).(entities.User)
	task.UserID = user.ID
	if err := task.FormatBoilerplateCode(); err != nil {
		handler.Logger.Log(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return
	}

	id, err := handler.TaskStorage.CreateTask(task)
	if err != nil {
		err = errors.Wrap(err, "[error] unable to create the task")
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

	if ok := handler.checkAccessToTask(writer, request, uint(id)); !ok {
		return
	}

	var task entities.Task
	if err := httputils.ReadJSON(request.Body, &task); err != nil {
		err = errors.Wrap(err, "[error] unable to decode the task data")
		handler.Logger.Log(err)
		http.Error(writer, err.Error(), http.StatusBadRequest)

		return
	}

	if err := task.FormatBoilerplateCode(); err != nil {
		handler.Logger.Log(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return
	}

	if err := handler.TaskStorage.UpdateTask(uint(id), task); err != nil {
		err = errors.Wrap(err, "[error] unable to update the task")
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

	if ok := handler.checkAccessToTask(writer, request, uint(id)); !ok {
		return
	}

	if err := handler.TaskStorage.DeleteTask(uint(id)); err != nil {
		err = errors.Wrap(err, "[error] unable to delete the task")
		handler.Logger.Log(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return
	}
}

func (handler TaskHandler) checkAccessToTask(
	writer http.ResponseWriter,
	request *http.Request,
	id uint,
) bool {
	task, err := handler.TaskStorage.GetTask(0, id)
	if err != nil {
		err = errors.Wrap(err, "[error] unable to get the task")
		handler.Logger.Log(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return false
	}

	user := request.Context().Value(userContextKey{}).(entities.User)
	if user.ID != task.UserID {
		const errMessage = "[error] managerial access to the task is denied"
		handler.Logger.Log(errMessage)
		http.Error(writer, errMessage, http.StatusForbidden)

		return false
	}

	return true
}
