package storages

import (
	"errors"

	"github.com/thewizardplusplus/go-exercises-backend/entities"
	"gorm.io/gorm"
)

// UserStorage ...
type UserStorage struct {
	db          *gorm.DB
	hashingCost int
}

// NewUserStorage ...
func NewUserStorage(db *gorm.DB, hashingCost int) UserStorage {
	return UserStorage{db: db, hashingCost: hashingCost}
}

// GetUser ...
func (storage UserStorage) GetUser(username string) (entities.User, error) {
	var user entities.User
	err := storage.db.
		Where(&entities.User{Username: username}).
		First(&user).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = entities.ErrNotFound
		}
		return entities.User{}, err
	}

	return user, nil
}

// CreateUser ...
func (storage UserStorage) CreateUser(user entities.User) error {
	user.Model = gorm.Model{} // reset the fields that are filled in automatically
	if err := user.HashPassword(storage.hashingCost); err != nil {
		return err
	}

	if err := storage.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}
