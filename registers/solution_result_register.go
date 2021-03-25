package registers

import (
	"github.com/pkg/errors"
	"github.com/thewizardplusplus/go-exercises-backend/entities"
)

// SolutionUpdater ...
type SolutionUpdater interface {
	UpdateSolution(id uint, solution entities.Solution) error
}

// SolutionResultRegister ...
type SolutionResultRegister struct {
	SolutionUpdater SolutionUpdater
}

// RegisterSolutionResult ...
func (register SolutionResultRegister) RegisterSolutionResult(
	solution entities.Solution,
) error {
	// update only these specific fields
	solutionUpdate := entities.Solution{
		IsCorrect: solution.IsCorrect,
		Result:    solution.Result,
	}
	if err := register.SolutionUpdater.UpdateSolution(
		solution.ID,
		solutionUpdate,
	); err != nil {
		return errors.Wrap(err, "unable to update the solution")
	}

	return nil
}
