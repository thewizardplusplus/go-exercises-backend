package registers

import (
	"github.com/go-log/log"
	"github.com/pkg/errors"
	"github.com/thewizardplusplus/go-exercises-backend/entities"
)

// SolutionQueue ...
type SolutionQueue interface {
	AddSolution(solution entities.Solution) error
}

// SolutionRegister ...
type SolutionRegister struct {
	TaskGetter     entities.TaskGetter
	SolutionGetter entities.SolutionGetter
	SolutionQueue  SolutionQueue
	Logger         log.Logger
}

// RegisterSolution ...
func (register SolutionRegister) RegisterSolution(id uint) {
	if err := register.performRegistration(id); err != nil {
		err = errors.Wrapf(err, "[error] unable to register solution #%d", id)
		register.Logger.Log(err)

		return
	}

	register.Logger.Logf("[info] solution #%d has been registered", id)
}

func (register SolutionRegister) performRegistration(id uint) error {
	solution, err := register.SolutionGetter.GetSolution(id)
	if err != nil {
		return errors.Wrap(err, "unable to get the solution")
	}

	task, err := register.TaskGetter.GetTask(0, solution.TaskID)
	if err != nil {
		return errors.Wrap(err, "unable to get the task")
	}

	solution.Task = task
	if err := register.SolutionQueue.AddSolution(solution); err != nil {
		return errors.Wrap(err, "unable to enqueue the solution")
	}

	return nil
}
