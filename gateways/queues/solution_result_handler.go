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
	if err := handler.performHandling(message); err != nil {
		err = errors.Wrap(err, "[error] unable to handle the solution result")
		handler.Logger.Log(err)

		message.Reject(true)
		return
	}

	handler.Logger.Log("[info] solution result has been handled")
	message.Ack(false)
}

// HandleSolutionResult ...
func (handler SolutionResultHandler) performHandling(
	message amqp.Delivery,
) error {
	var solution entities.Solution
	if err := json.Unmarshal(message.Body, &solution); err != nil {
		return errors.Wrap(err, "unable to unmarshal the solution")
	}

	if err := handler.SolutionResultRegister.RegisterSolutionResult(
		solution,
	); err != nil {
		return errors.Wrap(err, "unable to register the solution result")
	}

	return nil
}
