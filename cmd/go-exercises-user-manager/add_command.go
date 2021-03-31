package main

import (
	"github.com/pkg/errors"
	"github.com/thewizardplusplus/go-exercises-backend/entities"
	"github.com/thewizardplusplus/go-exercises-backend/gateways/storages"
	"golang.org/x/crypto/bcrypt"
)

type addCommand struct {
	Username    string `kong:"required,short='u',help='Username.'"`
	Password    string `kong:"short='p',help='User password.'"`
	HashingCost int    `kong:"short='c',name='cost',default='10',help='Cost of user password hashing (range: [4, 31]).'"` // nolint: lll
}

func (command addCommand) Validate() error {
	if command.HashingCost < bcrypt.MinCost {
		return errors.Errorf("cost is too low (minimum: %d)", bcrypt.MinCost)
	}
	if command.HashingCost > bcrypt.MaxCost {
		return errors.Errorf("cost is too high (maximum: %d)", bcrypt.MaxCost)
	}
	return nil
}

func (command addCommand) Run(ctx commandContext) error {
	userStorage := storages.NewUserStorage(ctx.DB, command.HashingCost)
	return userStorage.CreateUser(entities.User{
		Username: command.Username,
		Password: command.Password,
	})
}
