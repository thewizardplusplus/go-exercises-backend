package registers

import (
	"github.com/go-log/log"
	"github.com/pkg/errors"
)

// FailableSolutionRegister ...
type FailableSolutionRegister interface {
	RegisterSolution(id uint) error
}

// SolutionRegister ...
type SolutionRegister struct {
	FailableSolutionRegister FailableSolutionRegister
	Logger                   log.Logger
}

// RegisterSolution ...
func (register SolutionRegister) RegisterSolution(id uint) {
	if err := register.FailableSolutionRegister.RegisterSolution(id); err != nil {
		err = errors.Wrapf(err, "[error] unable to register solution #%d", id)
		register.Logger.Log(err)

		return
	}

	register.Logger.Logf("[info] solution #%d has been registered", id)
}
