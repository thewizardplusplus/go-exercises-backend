package queues

import (
	"github.com/pkg/errors"
	"github.com/thewizardplusplus/go-exercises-backend/entities"
	rabbitmqutils "github.com/thewizardplusplus/go-rabbitmq-utils"
)

// SolutionQueue ...
type SolutionQueue struct {
	SolutionQueueName string
	Client            rabbitmqutils.Client
}

// AddSolution ...
func (queue SolutionQueue) AddSolution(solution entities.Solution) error {
	err := queue.Client.PublishMessage(queue.SolutionQueueName, "", solution)
	if err != nil {
		return errors.Wrap(err, "unable to publish the solution")
	}

	return nil
}
