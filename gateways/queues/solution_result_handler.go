package queues

import (
	"bytes"

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
}

// HandleMessage ...
func (handler SolutionResultHandler) HandleMessage(
	message amqp.Delivery,
) error {
	var solution entities.Solution
	reader := bytes.NewReader(message.Body)
	if err := httputils.ReadJSON(reader, &solution); err != nil {
		return errors.Wrap(err, "unable to unmarshal the solution")
	}

	err := handler.SolutionResultRegister.RegisterSolutionResult(solution)
	if err != nil {
		return errors.Wrap(err, "unable to register the solution result")
	}

	return nil
}
