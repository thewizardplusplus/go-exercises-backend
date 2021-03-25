package queues

import (
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

// Client ...
type Client struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

// NewClient ...
func NewClient(queueDSN string) (Client, error) {
	connection, err := amqp.Dial(queueDSN)
	if err != nil {
		return Client{}, errors.Wrap(err, "unable to dial the message broker")
	}

	channel, err := connection.Channel()
	if err != nil {
		return Client{}, errors.Wrap(err, "unable to open the channel")
	}

	if _, err := channel.QueueDeclare(
		"solution_queue", // queue name
		true,             // durable
		false,            // auto-delete
		false,            // exclusive
		false,            // no wait
		nil,              // arguments
	); err != nil {
		return Client{}, errors.Wrap(err, "unable to declare the queue")
	}

	client := Client{connection: connection, channel: channel}
	return client, nil
}

// Close ...
func (client Client) Close() error {
	if err := client.channel.Close(); err != nil {
		return errors.Wrap(err, "unable to close the channel")
	}

	if err := client.connection.Close(); err != nil {
		return errors.Wrap(err, "unable to close the connection")
	}

	return nil
}
