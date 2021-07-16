package queues

import (
	"bytes"

	"github.com/go-log/log"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"github.com/thewizardplusplus/go-exercises-backend/entities"
	httputils "github.com/thewizardplusplus/go-http-utils"
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

// HandleMessage ...
func (handler SolutionResultHandler) HandleMessage(message amqp.Delivery) {
	solution, err := handler.performHandling(message)
	if err != nil {
		handler.Logger.Log(errors.Wrapf(
			err,
			"[error] unable to handle the result of solution #%d",
			solution.ID,
		))

		// requeue the message only once
		message.Reject(!message.Redelivered) // nolint: gosec, errcheck
		return
	}

	handler.Logger.
		Logf("[info] result of solution #%d has been handled", solution.ID)
	message.Ack(false) // nolint: gosec, errcheck
}

func (handler SolutionResultHandler) performHandling(
	message amqp.Delivery,
) (entities.Solution, error) {
	var solution entities.Solution
	reader := bytes.NewReader(message.Body)
	if err := httputils.ReadJSON(reader, &solution); err != nil {
		return solution, errors.Wrap(err, "unable to unmarshal the solution")
	}

	if err := handler.SolutionResultRegister.RegisterSolutionResult(
		solution,
	); err != nil {
		return solution, errors.Wrap(err, "unable to register the solution result")
	}

	return solution, nil
}
