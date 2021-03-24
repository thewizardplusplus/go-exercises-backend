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

// SolutionStorage ...
type SolutionStorage interface {
	GetSolutions(taskID uint) ([]entities.Solution, error)
	GetSolution(id uint) (entities.Solution, error)
	CreateSolution(taskID uint, solution entities.Solution) (id uint, err error)
}

// SolutionHandler ...
type SolutionHandler struct {
	SolutionStorage SolutionStorage
	Logger          log.Logger
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

	solutions, err := handler.SolutionStorage.GetSolutions(uint(taskID))
	if err != nil {
		err = errors.Wrap(err, "[error] unable to get the solutions")
		handler.Logger.Log(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return
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

	solution, err := handler.SolutionStorage.GetSolution(uint(id))
	if err != nil {
		err = errors.Wrap(err, "[error] unable to get the solution")
		handler.Logger.Log(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return
	}

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

	id, err := handler.SolutionStorage.CreateSolution(uint(taskID), solution)
	if err != nil {
		err = errors.Wrap(err, "[error] unable to create the solution")
		handler.Logger.Log(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return
	}

	idAsModel := entities.Solution{Model: gorm.Model{ID: id}}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(idAsModel) // nolint: gosec, errcheck
}
