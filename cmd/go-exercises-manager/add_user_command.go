package main

import (
	"encoding/json"
	"os"

	"github.com/pkg/errors"
	"github.com/thewizardplusplus/go-exercises-backend/entities"
	"github.com/thewizardplusplus/go-exercises-backend/gateways/storages"
)

type addUserCommand struct {
	basicUserCommand
}

func (command addUserCommand) Run(ctx commandContext) error {
	user := entities.User{Username: command.Username, Password: command.Password}
	if user.Password == "" {
		if err := user.GeneratePassword(command.PasswordLength); err != nil {
			return errors.Wrap(err, "unable to generate the user password")
		}
	}

	userStorage := storages.NewUserStorage(ctx.DB, command.HashingCost)
	if err := userStorage.CreateUser(user); err != nil {
		return errors.Wrap(err, "unable to create the user")
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(user); err != nil {
		return errors.Wrap(err, "unable to marshal the user")
	}

	return nil
}
