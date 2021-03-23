package storages

import (
	"github.com/thewizardplusplus/go-exercises-backend/entities"
	"gorm.io/gorm"
)

// TaskStorage ...
type TaskStorage struct {
	db *gorm.DB
}

// NewTaskStorage ...
func NewTaskStorage(db *gorm.DB) TaskStorage {
	return TaskStorage{db: db}
}

// GetTasks ...
func (storage TaskStorage) GetTasks() ([]entities.Task, error) {
	panic("not yet implemented")
}

// GetTask ...
func (storage TaskStorage) GetTask(id uint) (entities.Task, error) {
	panic("not yet implemented")
}

// CreateTask ...
func (storage TaskStorage) CreateTask(task entities.Task) (id uint, err error) {
	panic("not yet implemented")
}

// UpdateTask ...
func (storage TaskStorage) UpdateTask(id uint, task entities.Task) error {
	panic("not yet implemented")
}

// DeleteTask ...
func (storage TaskStorage) DeleteTask(id uint) error {
	panic("not yet implemented")
}
