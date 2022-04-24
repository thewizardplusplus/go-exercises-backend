package main

import (
	"github.com/pkg/errors"
	"github.com/thewizardplusplus/go-exercises-backend/entities"
	"github.com/thewizardplusplus/go-exercises-backend/gateways/storages"
)

type updateUserCommand struct {
	basicUserCommand

	NewUsername string `kong:"short='U',help='New username.'"`
}

func (command updateUserCommand) Run(ctx commandContext) error {
	user := entities.User{
		Username:   command.NewUsername,
		Password:   command.Password,
		IsDisabled: command.Disable,
	}
	if command.GeneratePassword {
		if err := user.GeneratePassword(command.PasswordLength); err != nil {
			return errors.Wrap(err, "unable to generate the user password")
		}
	}

	userStorage := storages.NewUserStorage(ctx.DB, command.HashingCost)
	if err := userStorage.UpdateUser(command.Username, user); err != nil {
		return errors.Wrap(err, "unable to update the user")
	}

	if err := outputUser(user); err != nil {
		return errors.Wrap(err, "unable to output the user")
	}

	return nil
}
