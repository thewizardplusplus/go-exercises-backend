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
	"github.com/thewizardplusplus/go-exercises-backend/gateways/storages"
	httputils "github.com/thewizardplusplus/go-http-utils"
)

type options struct {
	Server struct {
		Address string `env:"SERVER_ADDRESS" envDefault:":8080"`
	}
	Storage struct {
		Address string `env:"STORAGE_ADDRESS" envDefault:"postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"` // nolint: lll
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

	router := handlers.NewRouter(handlers.RouterDependencies{
		TaskStorage:     storages.NewTaskStorage(db),
		SolutionStorage: storages.NewSolutionStorage(db),
		Logger:          print.New(logger),
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
