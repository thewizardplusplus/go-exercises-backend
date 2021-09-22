package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/caarlos0/env"
	"github.com/go-log/log/print"
	middlewares "github.com/gorilla/handlers"
	"github.com/sethvargo/go-password/password"
	"github.com/thewizardplusplus/go-exercises-backend/gateways/handlers"
	"github.com/thewizardplusplus/go-exercises-backend/gateways/queues"
	"github.com/thewizardplusplus/go-exercises-backend/gateways/storages"
	"github.com/thewizardplusplus/go-exercises-backend/registers"
	httputils "github.com/thewizardplusplus/go-http-utils"
	rabbitmqutils "github.com/thewizardplusplus/go-rabbitmq-utils"
)

type options struct {
	Server struct {
		Address        string `env:"SERVER_ADDRESS" envDefault:":8080"`
		StaticFilePath string `env:"SERVER_STATIC_FILE_PATH" envDefault:"./static"`
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
	Authorization struct {
		TokenSigningKey string        `env:"AUTHORIZATION_TOKEN_SIGNING_KEY"`
		TokenTTL        time.Duration `env:"AUTHORIZATION_TOKEN_TTL" envDefault:"24h"`
	}
}

// nolint: gocyclo
func main() {
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lmicroseconds)

	var options options
	if err := env.Parse(&options); err != nil {
		logger.Fatalf("[error] unable to parse the options: %v", err)
	}

	if options.Authorization.TokenSigningKey == "" {
		tokenSigningKey, err := password.Generate(
			23,    // length
			0,     // number of digits
			0,     // number of symbols
			false, // no upper case
			true,  // allow repeat
		)
		if err != nil {
			logger.Fatalf("[error] unable to generate the token signing key: %v", err)
		}

		logger.Printf(
			"[info] token signing key has been generated: %s",
			tokenSigningKey,
		)

		options.Authorization.TokenSigningKey = tokenSigningKey
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

	messageBrokerClient, err := rabbitmqutils.NewClient(
		options.MessageBroker.Address,
		rabbitmqutils.WithMaximalQueueSize(options.SolutionRegister.BufferSize),
		rabbitmqutils.WithQueues([]string{
			queues.SolutionQueueName,
			queues.SolutionResultQueueName,
		}),
	)
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
			SolutionQueue: queues.SolutionQueue{
				SolutionQueueName: queues.SolutionQueueName,
				Client:            messageBrokerClient,
			},
			Logger: print.New(logger),
		},
	)
	go solutionRegister.StartConcurrently(options.SolutionRegister.Concurrency)
	defer solutionRegister.Stop()

	solutionResultConsumer, err := rabbitmqutils.NewMessageConsumer(
		messageBrokerClient,
		queues.SolutionResultQueueName,
		rabbitmqutils.Acknowledger{
			MessageHandling: rabbitmqutils.TwiceMessageHandling,
			MessageHandler: rabbitmqutils.JSONMessageHandler{
				MessageHandler: queues.SolutionResultHandler{
					SolutionResultRegister: registers.SolutionResultRegister{
						SolutionUpdater: storages.NewSolutionStorage(db),
					},
				},
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

	routerOptions := handlers.RouterOptions{
		StaticFilePath:  options.Server.StaticFilePath,
		TokenSigningKey: options.Authorization.TokenSigningKey,
		TokenTTL:        options.Authorization.TokenTTL,
	}
	router := handlers.NewRouter(routerOptions, handlers.RouterDependencies{
		TaskStorage:      storages.NewTaskStorage(db),
		SolutionStorage:  storages.NewSolutionStorage(db),
		SolutionRegister: solutionRegister,
		UserGetter:       storages.NewUserStorage(db, 0),
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
