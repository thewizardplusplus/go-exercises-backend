package queues

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"github.com/thewizardplusplus/go-exercises-backend/entities"
)

// SolutionQueue ...
type SolutionQueue struct {
	client Client
}

// NewSolutionQueue ...
func NewSolutionQueue(client Client) SolutionQueue {
	return SolutionQueue{client: client}
}

// AddSolution ...
func (queue SolutionQueue) AddSolution(solution entities.Solution) error {
	solutionAsJSON, err := json.Marshal(solution)
	if err != nil {
		return errors.Wrap(err, "unable to marshal the solution")
	}

	if err := queue.client.channel.Publish(
		"",                // exchange
		SolutionQueueName, // queue name
		false,             // mandatory
		false,             // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        solutionAsJSON,
		},
	); err != nil {
		return errors.Wrap(err, "unable to publish the solution")
	}

	return nil
}
