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
	"gorm.io/gorm"
)

// SolutionStorage ...
type SolutionStorage interface {
	entities.SolutionGetter

	GetSolutions(userID uint, taskID uint, pagination entities.Pagination) (
		[]entities.Solution,
		error,
	)
	CreateSolution(taskID uint, solution entities.Solution) (id uint, err error)
}

// SolutionHandler ...
type SolutionHandler struct {
	SolutionStorage  SolutionStorage
	SolutionRegister entities.SolutionRegister
	Logger           log.Logger
}

// GetSolutions ...
func (handler SolutionHandler) GetSolutions(
	writer http.ResponseWriter,
	request *http.Request,
) {
	taskIDAsStr := mux.Vars(request)["taskID"]
	taskID, err := strconv.ParseUint(taskIDAsStr, 10, 64)
	if err != nil {
		err = errors.Wrap(err, "[error] unable to decode the task ID")
		handler.Logger.Log(err)
		http.Error(writer, err.Error(), http.StatusBadRequest)

		return
	}

	var pagination entities.Pagination
	err = schema.NewDecoder().Decode(&pagination, request.URL.Query())
	if err != nil {
		err = errors.Wrap(err, "[error] unable to decode the pagination parameters")
		handler.Logger.Log(err)
		http.Error(writer, err.Error(), http.StatusBadRequest)

		return
	}

	user := request.Context().Value(userContextKey{}).(entities.User)
	solutions, err :=
		handler.SolutionStorage.GetSolutions(user.ID, uint(taskID), pagination)
	if err != nil {
		err = errors.Wrap(err, "[error] unable to get the solutions")
		handler.Logger.Log(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return
	}

	for index := range solutions {
		solutions[index].User.PasswordHash = ""
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(solutions) // nolint: gosec, errcheck
}

// GetSolution ...
func (handler SolutionHandler) GetSolution(
	writer http.ResponseWriter,
	request *http.Request,
) {
	idAsStr := mux.Vars(request)["id"]
	id, err := strconv.ParseUint(idAsStr, 10, 64)
	if err != nil {
		err = errors.Wrap(err, "[error] unable to decode the solution ID")
		handler.Logger.Log(err)
		http.Error(writer, err.Error(), http.StatusBadRequest)

		return
	}

	if ok := handler.checkAccessToSolution(writer, request, uint(id)); !ok {
		return
	}

	solution, err := handler.SolutionStorage.GetSolution(uint(id))
	if err != nil {
		err = errors.Wrap(err, "[error] unable to get the solution")
		handler.Logger.Log(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return
	}

	solution.User.PasswordHash = ""

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(solution) // nolint: gosec, errcheck
}

// CreateSolution ...
func (handler SolutionHandler) CreateSolution(
	writer http.ResponseWriter,
	request *http.Request,
) {
	taskIDAsStr := mux.Vars(request)["taskID"]
	taskID, err := strconv.ParseUint(taskIDAsStr, 10, 64)
	if err != nil {
		err = errors.Wrap(err, "[error] unable to decode the task ID")
		handler.Logger.Log(err)
		http.Error(writer, err.Error(), http.StatusBadRequest)

		return
	}

	var solution entities.Solution
	if err := json.NewDecoder(request.Body).Decode(&solution); err != nil {
		err = errors.Wrap(err, "[error] unable to decode the solution data")
		handler.Logger.Log(err)
		http.Error(writer, err.Error(), http.StatusBadRequest)

		return
	}

	user := request.Context().Value(userContextKey{}).(entities.User)
	solution.UserID = user.ID

	id, err := handler.SolutionStorage.CreateSolution(uint(taskID), solution)
	if err != nil {
		err = errors.Wrap(err, "[error] unable to create the solution")
		handler.Logger.Log(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return
	}

	handler.SolutionRegister.RegisterSolution(id)

	idAsModel := entities.Solution{Model: gorm.Model{ID: id}}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(idAsModel) // nolint: gosec, errcheck
}

func (handler SolutionHandler) checkAccessToSolution(
	writer http.ResponseWriter,
	request *http.Request,
	id uint,
) bool {
	solution, err := handler.SolutionStorage.GetSolution(id)
	if err != nil {
		err = errors.Wrap(err, "[error] unable to get the solution")
		handler.Logger.Log(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return false
	}

	user := request.Context().Value(userContextKey{}).(entities.User)
	if user.ID != solution.UserID {
		const errMessage = "[error] access to the solution is denied"
		handler.Logger.Log(errMessage)
		http.Error(writer, errMessage, http.StatusForbidden)

		return false
	}

	return true
}
