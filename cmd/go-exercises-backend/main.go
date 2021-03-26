package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/caarlos0/env"
	"github.com/go-log/log/print"
	middlewares "github.com/gorilla/handlers"
	"github.com/thewizardplusplus/go-exercises-backend/gateways/handlers"
	"github.com/thewizardplusplus/go-exercises-backend/gateways/queues"
	"github.com/thewizardplusplus/go-exercises-backend/gateways/storages"
	"github.com/thewizardplusplus/go-exercises-backend/registers"
	httputils "github.com/thewizardplusplus/go-http-utils"
)

type options struct {
	Server struct {
		Address string `env:"SERVER_ADDRESS" envDefault:":8080"`
	}
	Storage struct {
		Address string `env:"STORAGE_ADDRESS" envDefault:"postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"` // nolint: lll
	}
	MessageBroker struct {
		Address string `env:"MESSAGE_BROKER_ADDRESS" envDefault:"amqp://rabbitmq:rabbitmq@localhost:5672"` // nolint: lll
	}
	SolutionRegister struct {
		BufferSize  int `env:"SOLUTION_REGISTER_BUFFER_SIZE" envDefault:"1000"`
		Concurrency int `env:"SOLUTION_REGISTER_CONCURRENCY" envDefault:"1000"`
	}
}

func main() {
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lmicroseconds)

	var options options
	if err := env.Parse(&options); err != nil {
		logger.Fatalf("[error] unable to parse the options: %v", err)
	}

	db, err := storages.OpenDB(options.Storage.Address, logger)
	if err != nil {
		logger.Fatalf("[error] unable to create the storage: %v", err)
	}
	defer func() {
		if err := storages.CloseDB(db); err != nil {
			logger.Fatalf("[error] unable to close the storage: %v", err)
		}
	}()

	messageBrokerClient, err := queues.NewClient(options.MessageBroker.Address)
	if err != nil {
		logger.Fatalf("[error] unable to create the message broker client: %v", err)
	}
	defer func() {
		if err := messageBrokerClient.Close(); err != nil {
			logger.Fatalf("[error] unable to close the message broker client: %v", err)
		}
	}()

	solutionRegister := registers.NewConcurrentSolutionRegister(
		options.SolutionRegister.BufferSize,
		registers.SolutionRegister{
			TaskStorage:     storages.NewTaskStorage(db),
			SolutionStorage: storages.NewSolutionStorage(db),
			SolutionQueue:   queues.NewSolutionQueue(messageBrokerClient),
			Logger:          print.New(logger),
		},
	)
	go solutionRegister.StartConcurrently(options.SolutionRegister.Concurrency)
	defer solutionRegister.Stop()

	solutionResultConsumer, err := queues.NewSolutionResultConsumer(
		messageBrokerClient,
		queues.SolutionResultHandler{
			SolutionResultRegister: registers.SolutionResultRegister{
				SolutionUpdater: storages.NewSolutionStorage(db),
			},
			Logger: print.New(logger),
		},
	)
	if err != nil {
		logger.Fatalf(
			"[error] unable to create the solution result consumer: %v",
			err,
		)
	}
	go solutionResultConsumer.
		StartConcurrently(options.SolutionRegister.Concurrency)
	defer func() {
		if err := solutionResultConsumer.Stop(); err != nil {
			logger.Fatalf("[error] unable to stop the solution result consumer: %v", err)
		}
	}()

	router := handlers.NewRouter(handlers.RouterDependencies{
		TaskStorage:      storages.NewTaskStorage(db),
		SolutionStorage:  storages.NewSolutionStorage(db),
		SolutionRegister: solutionRegister,
		Logger:           print.New(logger),
	})
	router.Use(middlewares.RecoveryHandler(middlewares.RecoveryLogger(logger)))
	router.Use(func(next http.Handler) http.Handler {
		return middlewares.LoggingHandler(os.Stderr, next)
	})

	if ok := httputils.RunServer(
		context.Background(),
		&http.Server{
			Addr:    options.Server.Address,
			Handler: router,
		},
		print.New(logger),
		os.Interrupt,
	); !ok {
		os.Exit(1)
	}
}
