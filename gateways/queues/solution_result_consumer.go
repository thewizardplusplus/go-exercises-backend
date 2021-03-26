package queues

import (
	"context"
	"sync"

	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

// SolutionResultConsumerName ...
const SolutionResultConsumerName = SolutionResultQueueName + "_consumer"

// MessageHandler ...
type MessageHandler interface {
	HandleMessage(message amqp.Delivery)
}

// SolutionResultConsumer ...
type SolutionResultConsumer struct {
	client            Client
	messages          <-chan amqp.Delivery
	stoppingCtx       context.Context
	stoppingCtxCancel context.CancelFunc
	messageHandler    MessageHandler
}

// NewSolutionResultConsumer ...
func NewSolutionResultConsumer(
	client Client,
	messageHandler MessageHandler,
) (SolutionResultConsumer, error) {
	messages, err := client.channel.Consume(
		SolutionResultQueueName,    // queue name
		SolutionResultConsumerName, // consumer name
		false,                      // auto-acknowledge
		false,                      // exclusive
		false,                      // no local
		false,                      // no wait
		nil,                        // arguments
	)
	if err != nil {
		return SolutionResultConsumer{},
			errors.Wrap(err, "unable to start the message consumption")
	}

	stoppingCtx, stoppingCtxCancel := context.WithCancel(context.Background())
	consumer := SolutionResultConsumer{
		client:            client,
		messages:          messages,
		stoppingCtx:       stoppingCtx,
		stoppingCtxCancel: stoppingCtxCancel,
		messageHandler:    messageHandler,
	}

	return consumer, nil
}

// Start ...
func (consumer SolutionResultConsumer) Start() {
	for message := range consumer.messages {
		consumer.messageHandler.HandleMessage(message)
	}
}

// StartConcurrently ...
func (consumer SolutionResultConsumer) StartConcurrently(concurrency int) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(concurrency)

	for i := 0; i < concurrency; i++ {
		go func() {
			defer waitGroup.Done()

			consumer.Start()
		}()
	}

	waitGroup.Wait()
	consumer.stoppingCtxCancel()
}

// Stop ...
func (consumer SolutionResultConsumer) Stop() error {
	if err := consumer.client.channel.Cancel(
		SolutionResultConsumerName, // consumer name
		false,                      // no wait
	); err != nil {
		return errors.Wrap(err, "unable to cancel the message consumption")
	}

	<-consumer.stoppingCtx.Done()
	return nil
}
