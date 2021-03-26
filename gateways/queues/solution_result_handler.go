package queues

import (
	"encoding/json"

	"github.com/go-log/log"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"github.com/thewizardplusplus/go-exercises-backend/entities"
)

// SolutionResultRegister ...
type SolutionResultRegister interface {
	RegisterSolutionResult(solution entities.Solution) error
}

// SolutionResultHandler ...
type SolutionResultHandler struct {
	SolutionResultRegister SolutionResultRegister
	Logger                 log.Logger
}

// HandleSolutionResult ...
func (handler SolutionResultHandler) HandleSolutionResult(
	message amqp.Delivery,
) {
	solution, err := handler.performHandling(message)
	if err != nil {
		handler.Logger.Log(errors.Wrapf(
			err,
			"[error] unable to handle the result of solution #%d",
			solution.ID,
		))

		message.Reject(true)
		return
	}

	handler.Logger.
		Logf("[info] result of solution #%d has been handled", solution.ID)
	message.Ack(false)
}

// HandleSolutionResult ...
func (handler SolutionResultHandler) performHandling(
	message amqp.Delivery,
) (entities.Solution, error) {
	var solution entities.Solution
	if err := json.Unmarshal(message.Body, &solution); err != nil {
		return solution, errors.Wrap(err, "unable to unmarshal the solution")
	}

	if err := handler.SolutionResultRegister.RegisterSolutionResult(
		solution,
	); err != nil {
		return solution, errors.Wrap(err, "unable to register the solution result")
	}

	return solution, nil
}
