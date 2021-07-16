package queues

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/thewizardplusplus/go-exercises-backend/entities"
)

// SolutionResultRegister ...
type SolutionResultRegister interface {
	RegisterSolutionResult(solution entities.Solution) error
}

// SolutionResultHandler ...
type SolutionResultHandler struct {
	SolutionResultRegister SolutionResultRegister
}

// MessageType ...
func (handler SolutionResultHandler) MessageType() reflect.Type {
	return reflect.TypeOf(entities.Solution{})
}

// HandleMessage ...
func (handler SolutionResultHandler) HandleMessage(message interface{}) error {
	err := handler.SolutionResultRegister.
		RegisterSolutionResult(message.(entities.Solution))
	if err != nil {
		return errors.Wrap(err, "unable to register the solution result")
	}

	return nil
}
