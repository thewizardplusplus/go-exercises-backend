package main

import (
	"encoding/json"
	"os"

	"github.com/pkg/errors"
	"github.com/thewizardplusplus/go-exercises-backend/entities"
	"github.com/thewizardplusplus/go-exercises-backend/gateways/storages"
	"golang.org/x/crypto/bcrypt"
)

type updateUserCommand struct {
	Username         string `kong:"required,short='u',help='Username.'"`
	NewUsername      string `kong:"short='U',help='New username.'"`
	Password         string `kong:"short='p',help='User password.'"`
	HashingCost      int    `kong:"short='c',name='cost',default='10',help='Cost of the user password hashing (range: [4, 31]).'"`        // nolint: lll
	PasswordLength   int    `kong:"short='l',name='length',default='6',help='Length of the user password to be generated (minimum: 6).'"` // nolint: lll
	GeneratePassword bool   `kong:"short='g',name='generate',help='Generate the user password (only if it is empty).'"`                   // nolint: lll
	Disable          bool   `kong:"short='d',name='disable',help='Disable the user.'"`                                                    // nolint: lll
}

func (command updateUserCommand) Validate() error {
	if command.HashingCost < bcrypt.MinCost {
		return errors.Errorf("cost is too low (minimum: %d)", bcrypt.MinCost)
	}
	if command.HashingCost > bcrypt.MaxCost {
		return errors.Errorf("cost is too high (maximum: %d)", bcrypt.MaxCost)
	}

	if command.PasswordLength < minimalPasswordLength {
		return errors.Errorf(
			"length is too short (minimum: %d)",
			minimalPasswordLength,
		)
	}

	return nil
}

func (command updateUserCommand) Run(ctx commandContext) error {
	user := entities.User{
		Username:   command.NewUsername,
		Password:   command.Password,
		IsDisabled: command.Disable,
	}
	if user.Password == "" && command.GeneratePassword {
		if err := user.GeneratePassword(command.PasswordLength); err != nil {
			return errors.Wrap(err, "unable to generate the user password")
		}
	}

	userStorage := storages.NewUserStorage(ctx.DB, command.HashingCost)
	if err := userStorage.UpdateUser(command.Username, user); err != nil {
		return errors.Wrap(err, "unable to update the user")
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(user); err != nil {
		return errors.Wrap(err, "unable to marshal the user")
	}

	return nil
}
