package main

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

const minimalPasswordLength = 6

type basicUserCommand struct {
	Username       string `kong:"required,short='u',help='Username.'"`
	Password       string `kong:"short='p',help='User password.'"`
	HashingCost    int    `kong:"short='c',name='cost',default='10',help='Cost of the user password hashing (range: [4, 31]).'"`        // nolint: lll
	PasswordLength int    `kong:"short='l',name='length',default='6',help='Length of the user password to be generated (minimum: 6).'"` // nolint: lll
}

func (command basicUserCommand) Validate() error {
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
