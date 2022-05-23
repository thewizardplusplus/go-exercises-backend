package queues

import (
	"github.com/pkg/errors"
	"github.com/thewizardplusplus/go-exercises-backend/entities"
)

// MessagePublisher ...
type MessagePublisher interface {
	PublishMessage(queue string, messageID string, messageData interface{}) error
}

// SolutionQueue ...
type SolutionQueue struct {
	SolutionQueueName string
	MessagePublisher  MessagePublisher
}

// EnqueueSolution ...
func (queue SolutionQueue) EnqueueSolution(solution entities.Solution) error {
	err := queue.MessagePublisher.
		PublishMessage(queue.SolutionQueueName, "", solution)
	if err != nil {
		return errors.Wrap(err, "unable to publish the solution")
	}

	return nil
}
