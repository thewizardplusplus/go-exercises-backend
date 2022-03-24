package main

import (
	"log"
	"os"

	"github.com/alecthomas/kong"
	"github.com/caarlos0/env"
	"github.com/thewizardplusplus/go-exercises-backend/gateways/storages"
	"gorm.io/gorm"
)

type options struct {
	Storage struct {
		Address string `env:"STORAGE_ADDRESS" envDefault:"postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"` // nolint: lll
	}
}

type commandContext struct {
	DB *gorm.DB
}

type cli struct {
	AddUser    addUserCommand    `kong:"cmd,help='Add the user.'"`
	UpdateUser updateUserCommand `kong:"cmd,help='Update the user.'"`
}

func main() {
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lmicroseconds)

	var options options
	if err := env.Parse(&options); err != nil {
		logger.Fatalf("[error] unable to parse the options: %v", err)
	}

	ctx, err := kong.Must(&cli{}).Parse(os.Args[1:])
	if err != nil {
		logger.Fatalf("[error] unable to parse the CLI: %v", err)
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

	if err := ctx.Run(commandContext{DB: db}); err != nil {
		logger.Fatalf("[error] unable to process the CLI: %v", err)
	}
}
