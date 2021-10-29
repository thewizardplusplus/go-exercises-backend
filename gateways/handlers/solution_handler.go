package handlers

import (
	"net/http"

	"github.com/go-log/log"
	"github.com/gorilla/schema"
	"github.com/pkg/errors"
	"github.com/thewizardplusplus/go-exercises-backend/entities"
	httputils "github.com/thewizardplusplus/go-http-utils"
)

// SolutionUsecase ...
type SolutionUsecase interface {
	GetSolutions(userID uint, taskID uint, pagination entities.Pagination) (
		entities.SolutionGroup,
		error,
	)
	GetSolution(userID uint, solutionID uint) (entities.Solution, error)
	CreateSolution(userID uint, taskID uint, solution entities.Solution) (
		entities.Solution,
		error,
	)
	FormatSolution(solution entities.Solution) (entities.Solution, error)
}

// SolutionHandler ...
type SolutionHandler struct {
	SolutionUsecase SolutionUsecase
	Logger          log.Logger
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

	user := getUserFromRequest(request)
	solutionGroup, err :=
		handler.SolutionUsecase.GetSolutions(user.ID, uint(taskID), pagination)
	if err != nil {
		err = errors.Wrap(err, "[error] unable to get the solutions")
		const statusCode = http.StatusInternalServerError
		httputils.LoggingError(handler.Logger, writer, err, statusCode)

		return
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

	user := getUserFromRequest(request)
	solution, err := handler.SolutionUsecase.GetSolution(user.ID, uint(id))
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, entities.ErrManagerialAccessIsDenied) {
			statusCode = http.StatusForbidden
		}
		if errors.Is(err, entities.ErrNotFound) {
			statusCode = http.StatusNotFound
		}

		err = errors.Wrap(err, "[error] unable to get the solution")
		httputils.LoggingError(handler.Logger, writer, err, statusCode)

		return
	}

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

	user := getUserFromRequest(request)
	idAsModel, err :=
		handler.SolutionUsecase.CreateSolution(user.ID, uint(taskID), solution)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, entities.ErrUnableToFormatCode) {
			statusCode = http.StatusBadRequest
		}

		err = errors.Wrap(err, "[error] unable to create the solution")
		httputils.LoggingError(handler.Logger, writer, err, statusCode)

		return
	}

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

	solution, err := handler.SolutionUsecase.FormatSolution(solution)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, entities.ErrUnableToFormatCode) {
			statusCode = http.StatusBadRequest
		}

		err = errors.Wrap(err, "[error] unable to format the solution")
		httputils.LoggingError(handler.Logger, writer, err, statusCode)

		return
	}

	httputils.WriteJSON(writer, http.StatusOK, solution)
}
