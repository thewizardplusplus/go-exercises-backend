package queues

import (
	"github.com/pkg/errors"
	"github.com/thewizardplusplus/go-exercises-backend/entities"
	rabbitmqutils "github.com/thewizardplusplus/go-rabbitmq-utils"
)

// SolutionQueue ...
type SolutionQueue struct {
	client rabbitmqutils.Client
}

// NewSolutionQueue ...
func NewSolutionQueue(client rabbitmqutils.Client) SolutionQueue {
	return SolutionQueue{client: client}
}

// AddSolution ...
func (queue SolutionQueue) AddSolution(solution entities.Solution) error {
	err := queue.client.PublishMessage(SolutionQueueName, "", solution)
	if err != nil {
		return errors.Wrap(err, "unable to publish the solution")
	}

	return nil
}
