package handlers

import (
	"net/http"

	"github.com/go-log/log"
	"github.com/gorilla/schema"
	"github.com/pkg/errors"
	"github.com/thewizardplusplus/go-exercises-backend/entities"
	httputils "github.com/thewizardplusplus/go-http-utils"
	"gorm.io/gorm"
)

// SolutionStorage ...
type SolutionStorage interface {
	entities.SolutionGetter

	GetSolutions(userID uint, taskID uint, pagination entities.Pagination) (
		[]entities.Solution,
		error,
	)
	CountSolutions(userID uint, taskID uint) (int64, error)
	CreateSolution(taskID uint, solution entities.Solution) (id uint, err error)
}

// SolutionHandler ...
type SolutionHandler struct {
	TaskStorage      TaskStorage
	SolutionStorage  SolutionStorage
	SolutionRegister entities.SolutionRegister
	Logger           log.Logger
}

// GetSolutions ...
func (handler SolutionHandler) GetSolutions(
	writer http.ResponseWriter,
	request *http.Request,
) {
	var taskID uint
	err := httputils.ParsePathParameter(request, "taskID", &taskID)
	if err != nil {
		err = errors.Wrap(err, "[error] unable to decode the task ID")
		httputils.LoggingError(handler.Logger, writer, err, http.StatusBadRequest)

		return
	}

	var pagination entities.Pagination
	err = schema.NewDecoder().Decode(&pagination, request.URL.Query())
	if err != nil {
		err = errors.Wrap(err, "[error] unable to decode the pagination parameters")
		httputils.LoggingError(handler.Logger, writer, err, http.StatusBadRequest)

		return
	}

	user := request.Context().Value(userContextKey{}).(entities.User)
	solutions, err :=
		handler.SolutionStorage.GetSolutions(user.ID, uint(taskID), pagination)
	if err != nil {
		err = errors.Wrap(err, "[error] unable to get the solutions")
		const statusCode = http.StatusInternalServerError
		httputils.LoggingError(handler.Logger, writer, err, statusCode)

		return
	}

	solutionCount, err :=
		handler.SolutionStorage.CountSolutions(user.ID, uint(taskID))
	if err != nil {
		err = errors.Wrap(err, "[error] unable to count the solutions")
		const statusCode = http.StatusInternalServerError
		httputils.LoggingError(handler.Logger, writer, err, statusCode)

		return
	}

	solutionGroup :=
		entities.SolutionGroup{Solutions: solutions, TotalCount: solutionCount}
	for index := range solutionGroup.Solutions {
		solutionGroup.Solutions[index].User.PasswordHash = ""
	}

	httputils.WriteJSON(writer, http.StatusOK, solutionGroup)
}

// GetSolution ...
func (handler SolutionHandler) GetSolution(
	writer http.ResponseWriter,
	request *http.Request,
) {
	var id uint
	if err := httputils.ParsePathParameter(request, "id", &id); err != nil {
		err = errors.Wrap(err, "[error] unable to decode the solution ID")
		httputils.LoggingError(handler.Logger, writer, err, http.StatusBadRequest)

		return
	}

	if ok := handler.checkAccessToSolution(writer, request, uint(id)); !ok {
		return
	}

	solution, err := handler.SolutionStorage.GetSolution(uint(id))
	if err != nil {
		err = errors.Wrap(err, "[error] unable to get the solution")
		const statusCode = http.StatusInternalServerError
		httputils.LoggingError(handler.Logger, writer, err, statusCode)

		return
	}

	solution.User.PasswordHash = ""

	httputils.WriteJSON(writer, http.StatusOK, solution)
}

// CreateSolution ...
func (handler SolutionHandler) CreateSolution(
	writer http.ResponseWriter,
	request *http.Request,
) {
	var taskID uint
	err := httputils.ParsePathParameter(request, "taskID", &taskID)
	if err != nil {
		err = errors.Wrap(err, "[error] unable to decode the task ID")
		httputils.LoggingError(handler.Logger, writer, err, http.StatusBadRequest)

		return
	}

	var solution entities.Solution
	if err := httputils.ReadJSON(request.Body, &solution); err != nil {
		err = errors.Wrap(err, "[error] unable to decode the solution data")
		httputils.LoggingError(handler.Logger, writer, err, http.StatusBadRequest)

		return
	}

	user := request.Context().Value(userContextKey{}).(entities.User)
	solution.UserID = user.ID
	if err := solution.FormatCode(); err != nil {
		const statusCode = http.StatusInternalServerError
		httputils.LoggingError(handler.Logger, writer, err, statusCode)

		return
	}

	id, err := handler.SolutionStorage.CreateSolution(uint(taskID), solution)
	if err != nil {
		err = errors.Wrap(err, "[error] unable to create the solution")
		const statusCode = http.StatusInternalServerError
		httputils.LoggingError(handler.Logger, writer, err, statusCode)

		return
	}

	handler.SolutionRegister.RegisterSolution(id)

	idAsModel := entities.Solution{Model: gorm.Model{ID: id}}
	httputils.WriteJSON(writer, http.StatusOK, idAsModel)
}

// FormatSolution ...
func (handler SolutionHandler) FormatSolution(
	writer http.ResponseWriter,
	request *http.Request,
) {
	var solution entities.Solution
	if err := httputils.ReadJSON(request.Body, &solution); err != nil {
		err = errors.Wrap(err, "[error] unable to decode the solution data")
		httputils.LoggingError(handler.Logger, writer, err, http.StatusBadRequest)

		return
	}

	if err := solution.FormatCode(); err != nil {
		const statusCode = http.StatusInternalServerError
		httputils.LoggingError(handler.Logger, writer, err, statusCode)

		return
	}

	httputils.WriteJSON(writer, http.StatusOK, solution)
}

func (handler SolutionHandler) checkAccessToSolution(
	writer http.ResponseWriter,
	request *http.Request,
	id uint,
) bool {
	solution, err := handler.SolutionStorage.GetSolution(id)
	if err != nil {
		err = errors.Wrap(err, "[error] unable to get the solution")
		const statusCode = http.StatusInternalServerError
		httputils.LoggingError(handler.Logger, writer, err, statusCode)

		return false
	}

	task, err := handler.TaskStorage.GetTask(0, solution.TaskID)
	if err != nil {
		err = errors.Wrap(err, "[error] unable to get the task")
		const statusCode = http.StatusInternalServerError
		httputils.LoggingError(handler.Logger, writer, err, statusCode)

		return false
	}

	user := request.Context().Value(userContextKey{}).(entities.User)
	if user.ID != solution.UserID && user.ID != task.UserID {
		const errMessage = "[error] access to the solution is denied"
		httputils.LoggingError(handler.Logger, writer, err, http.StatusForbidden)

		return false
	}

	return true
}
