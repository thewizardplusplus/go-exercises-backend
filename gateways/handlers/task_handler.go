package handlers

import (
	"net/http"

	"github.com/go-log/log"
	"github.com/gorilla/schema"
	"github.com/pkg/errors"
	"github.com/thewizardplusplus/go-exercises-backend/entities"
	httputils "github.com/thewizardplusplus/go-http-utils"
)

// TaskUsecase ...
type TaskUsecase interface {
	entities.TaskGetter

	GetTasks(userID uint, pagination entities.Pagination) (
		entities.TaskGroup,
		error,
	)
	CreateTask(userID uint, task entities.Task) (entities.Task, error)
	UpdateTask(userID uint, taskID uint, task entities.Task) error
	DeleteTask(userID uint, taskID uint) error
}

// TaskHandler ...
type TaskHandler struct {
	TaskUsecase TaskUsecase
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
		httputils.LoggingError(handler.Logger, writer, err, http.StatusBadRequest)

		return
	}

	user := getUserFromRequest(request)
	taskGroup, err := handler.TaskUsecase.GetTasks(user.ID, pagination)
	if err != nil {
		err = errors.Wrap(err, "[error] unable to get the tasks")
		const statusCode = http.StatusInternalServerError
		httputils.LoggingError(handler.Logger, writer, err, statusCode)

		return
	}

	httputils.WriteJSON(writer, http.StatusOK, taskGroup)
}

// GetTask ...
func (handler TaskHandler) GetTask(
	writer http.ResponseWriter,
	request *http.Request,
) {
	var id uint
	if err := httputils.ParsePathParameter(request, "id", &id); err != nil {
		err = errors.Wrap(err, "[error] unable to decode the task ID")
		httputils.LoggingError(handler.Logger, writer, err, http.StatusBadRequest)

		return
	}

	user := getUserFromRequest(request)
	task, err := handler.TaskUsecase.GetTask(user.ID, uint(id))
	if err != nil {
		err = errors.Wrap(err, "[error] unable to get the task")
		statusCode := getStatusCodeFromError(err)
		httputils.LoggingError(handler.Logger, writer, err, statusCode)

		return
	}

	httputils.WriteJSON(writer, http.StatusOK, task)
}

// CreateTask ...
func (handler TaskHandler) CreateTask(
	writer http.ResponseWriter,
	request *http.Request,
) {
	var task entities.Task
	if err := httputils.ReadJSON(request.Body, &task); err != nil {
		err = errors.Wrap(err, "[error] unable to decode the task data")
		httputils.LoggingError(handler.Logger, writer, err, http.StatusBadRequest)

		return
	}

	user := getUserFromRequest(request)
	idAsModel, err := handler.TaskUsecase.CreateTask(user.ID, task)
	if err != nil {
		err = errors.Wrap(err, "[error] unable to create the task")
		statusCode := getStatusCodeFromError(err)
		httputils.LoggingError(handler.Logger, writer, err, statusCode)

		return
	}

	httputils.WriteJSON(writer, http.StatusOK, idAsModel)
}

// UpdateTask ...
func (handler TaskHandler) UpdateTask(
	writer http.ResponseWriter,
	request *http.Request,
) {
	var id uint
	if err := httputils.ParsePathParameter(request, "id", &id); err != nil {
		err = errors.Wrap(err, "[error] unable to decode the task ID")
		httputils.LoggingError(handler.Logger, writer, err, http.StatusBadRequest)

		return
	}

	var task entities.Task
	if err := httputils.ReadJSON(request.Body, &task); err != nil {
		err = errors.Wrap(err, "[error] unable to decode the task data")
		httputils.LoggingError(handler.Logger, writer, err, http.StatusBadRequest)

		return
	}

	user := getUserFromRequest(request)
	if err := handler.TaskUsecase.UpdateTask(user.ID, uint(id), task); err != nil {
		err = errors.Wrap(err, "[error] unable to update the task")
		statusCode := getStatusCodeFromError(err)
		httputils.LoggingError(handler.Logger, writer, err, statusCode)

		return
	}
}

// DeleteTask ...
func (handler TaskHandler) DeleteTask(
	writer http.ResponseWriter,
	request *http.Request,
) {
	var id uint
	if err := httputils.ParsePathParameter(request, "id", &id); err != nil {
		err = errors.Wrap(err, "[error] unable to decode the task ID")
		httputils.LoggingError(handler.Logger, writer, err, http.StatusBadRequest)

		return
	}

	user := getUserFromRequest(request)
	if err := handler.TaskUsecase.DeleteTask(user.ID, uint(id)); err != nil {
		err = errors.Wrap(err, "[error] unable to delete the task")
		statusCode := getStatusCodeFromError(err)
		httputils.LoggingError(handler.Logger, writer, err, statusCode)

		return
	}
}
