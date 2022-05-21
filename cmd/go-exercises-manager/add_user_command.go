package main

import (
	"github.com/pkg/errors"
	"github.com/thewizardplusplus/go-exercises-backend/entities"
	"github.com/thewizardplusplus/go-exercises-backend/gateways/storages"
)

type addUserCommand struct {
	basicUserCommand
}

func (command addUserCommand) Validate() error {
	if err := command.basicUserCommand.Validate(); err != nil {
		return err
	}

	if command.Password == "" && !command.GeneratePassword {
		return errors.New("password should not be empty or should be generated")
	}

	return nil
}

func (command addUserCommand) Run(ctx commandContext) error {
	user := entities.User{
		Username:   command.Username,
		Password:   command.Password,
		IsDisabled: command.Disable,
	}
	if command.GeneratePassword {
		if err := user.GeneratePassword(command.PasswordLength); err != nil {
			return errors.Wrap(err, "unable to generate the user password")
		}
	}

	userStorage := storages.NewUserStorage(ctx.DB, command.HashingCost)
	if err := userStorage.CreateUser(user); err != nil {
		return errors.Wrap(err, "unable to create the user")
	}

	if err := outputUser(user); err != nil {
		return errors.Wrap(err, "unable to output the user")
	}

	return nil
}
