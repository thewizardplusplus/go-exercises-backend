package usecases

import (
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
	CountSolutions(userID uint, taskID uint) (int64, error)
	CreateSolution(taskID uint, solution entities.Solution) (id uint, err error)
	UpdateSolution(id uint, solution entities.Solution) error
}

// SolutionQueue ...
type SolutionQueue interface {
	EnqueueSolution(solution entities.Solution) error
}

// SolutionUsecase ...
type SolutionUsecase struct {
	TaskGetter       entities.TaskGetter
	SolutionStorage  SolutionStorage
	SolutionRegister entities.SolutionRegister
	SolutionQueue    SolutionQueue
}

// GetSolutions ...
func (usecase SolutionUsecase) GetSolutions(
	userID uint,
	taskID uint,
	pagination entities.Pagination,
) (
	entities.SolutionGroup,
	error,
) {
	solutions, err :=
		usecase.SolutionStorage.GetSolutions(userID, taskID, pagination)
	if err != nil {
		return entities.SolutionGroup{},
			errors.Wrap(err, "unable to get the solutions")
	}

	solutionCount, err := usecase.SolutionStorage.CountSolutions(userID, taskID)
	if err != nil {
		return entities.SolutionGroup{},
			errors.Wrap(err, "unable to count the solutions")
	}

	solutionGroup :=
		entities.SolutionGroup{Solutions: solutions, TotalCount: solutionCount}
	for index := range solutionGroup.Solutions {
		solutionGroup.Solutions[index].User.PasswordHash = ""
	}

	return solutionGroup, nil
}

// GetSolution ...
func (usecase SolutionUsecase) GetSolution(
	userID uint,
	solutionID uint,
) (
	entities.Solution,
	error,
) {
	if err := usecase.checkAccessToSolution(userID, solutionID); err != nil {
		return entities.Solution{},
			errors.Wrap(err, "unable to check access to the solution")
	}

	solution, err := usecase.SolutionStorage.GetSolution(solutionID)
	if err != nil {
		return entities.Solution{}, errors.Wrap(err, "unable to get the solution")
	}

	solution.User.PasswordHash = ""
	return solution, nil
}

// CreateSolution ...
func (usecase SolutionUsecase) CreateSolution(
	userID uint,
	taskID uint,
	solution entities.Solution,
) (
	entities.Solution,
	error,
) {
	solution.UserID = userID
	if err := solution.FormatCode(); err != nil {
		return entities.Solution{}, errors.Wrap(err, "unable to format the code")
	}

	id, err := usecase.SolutionStorage.CreateSolution(taskID, solution)
	if err != nil {
		return entities.Solution{}, errors.Wrap(err, "unable to create the solution")
	}

	usecase.SolutionRegister.RegisterSolution(id)

	idAsModel := entities.Solution{Model: gorm.Model{ID: id}}
	return idAsModel, nil
}

// RegisterSolution ...
func (usecase SolutionUsecase) RegisterSolution(id uint) error {
	solution, err := usecase.SolutionStorage.GetSolution(id)
	if err != nil {
		return errors.Wrap(err, "unable to get the solution")
	}

	task, err := usecase.TaskGetter.GetTask(
		0, // user does not matter
		solution.TaskID,
	)
	if err != nil {
		return errors.Wrap(err, "unable to get the task")
	}

	solution.Task = task
	if err := usecase.SolutionQueue.EnqueueSolution(solution); err != nil {
		return errors.Wrap(err, "unable to enqueue the solution")
	}

	return nil
}

// RegisterSolutionResult ...
func (usecase SolutionUsecase) RegisterSolutionResult(
	solution entities.Solution,
) error {
	// update only these specific fields
	solutionUpdate := entities.Solution{
		IsCorrect: solution.IsCorrect,
		Result:    solution.Result,
	}
	err := usecase.SolutionStorage.UpdateSolution(solution.ID, solutionUpdate)
	if err != nil {
		return errors.Wrap(err, "unable to update the solution")
	}

	return nil
}

func (usecase SolutionUsecase) checkAccessToSolution(
	userID uint,
	solutionID uint,
) error {
	solution, err := usecase.SolutionStorage.GetSolution(solutionID)
	if err != nil {
		return errors.Wrap(err, "unable to get the solution")
	}

	task, err := usecase.TaskGetter.GetTask(userID, solution.TaskID)
	if err != nil {
		return errors.Wrap(err, "unable to get the task")
	}

	if userID != solution.UserID && userID != task.UserID {
		return entities.ErrManagerialAccessIsDenied
	}

	return nil
}
