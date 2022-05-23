package usecases

import (
	"github.com/pkg/errors"
	"github.com/thewizardplusplus/go-exercises-backend/entities"
)

// UserStorage ...
type UserStorage interface {
	CreateUser(user entities.User) error
	UpdateUser(username string, user entities.User) error
}

// UserUsecase ...
type UserUsecase struct {
	HashingCost int
	UserStorage UserStorage
}

// CreateUser ...
func (usecase UserUsecase) CreateUser(user entities.User) error {
	if err := user.HashPassword(usecase.HashingCost); err != nil {
		return err
	}

	if err := usecase.UserStorage.CreateUser(user); err != nil {
		return errors.Wrap(err, "unable to create the user")
	}

	return nil
}

// UpdateUser ...
func (usecase UserUsecase) UpdateUser(
	username string,
	user entities.User,
) error {
	if user.Password != "" {
		if err := user.HashPassword(usecase.HashingCost); err != nil {
			return err
		}
	} else {
		user.PasswordHash = ""
	}

	if err := usecase.UserStorage.UpdateUser(username, user); err != nil {
		return errors.Wrap(err, "unable to update the user")
	}

	return nil
}
