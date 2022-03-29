package main

import (
	"encoding/json"
	"os"

	"github.com/pkg/errors"
	"github.com/thewizardplusplus/go-exercises-backend/entities"
	"golang.org/x/crypto/bcrypt"
)

const minimalPasswordLength = 6

type basicUserCommand struct {
	Username         string `kong:"required,short='u',help='Username.'"`
	Password         string `kong:"short='p',help='User password.'"`
	HashingCost      int    `kong:"short='c',name='cost',default='10',help='Cost of the user password hashing (range: [4, 31]).'"`        // nolint: lll
	GeneratePassword bool   `kong:"short='g',name='generate',help='Generate the user password.'"`                                         // nolint: lll
	PasswordLength   int    `kong:"short='l',name='length',default='6',help='Length of the user password to be generated (minimum: 6).'"` // nolint: lll
}

func (command basicUserCommand) Validate() error {
	if command.Password == "" && !command.GeneratePassword {
		return errors.New("password is required")
	}

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

func outputUser(user entities.User) error {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(user); err != nil {
		return errors.Wrap(err, "unable to marshal the user")
	}

	return nil
}
