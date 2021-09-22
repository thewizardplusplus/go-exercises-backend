package main

import (
	"encoding/json"
	"os"

	"github.com/pkg/errors"
	"github.com/sethvargo/go-password/password"
	"github.com/thewizardplusplus/go-exercises-backend/entities"
	"github.com/thewizardplusplus/go-exercises-backend/gateways/storages"
	"golang.org/x/crypto/bcrypt"
)

const minimalPasswordLength = 6

type addUserCommand struct {
	Username       string `kong:"required,short='u',help='Username.'"`
	Password       string `kong:"short='p',help='User password.'"`
	HashingCost    int    `kong:"short='c',name='cost',default='10',help='Cost of the user password hashing (range: [4, 31]).'"`        // nolint: lll
	PasswordLength int    `kong:"short='l',name='length',default='6',help='Length of the user password to be generated (minimum: 6).'"` // nolint: lll
}

func (command addUserCommand) Validate() error {
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

func (command addUserCommand) Run(ctx commandContext) error {
	user := entities.User{Username: command.Username, Password: command.Password}
	if user.Password == "" {
		password, err := password.Generate(
			command.PasswordLength, // length
			0,                      // number of digits
			0,                      // number of symbols
			false,                  // no upper case
			true,                   // allow repeat
		)
		if err != nil {
			return errors.Wrap(err, "unable to generate the user password")
		}

		user.Password = password
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
