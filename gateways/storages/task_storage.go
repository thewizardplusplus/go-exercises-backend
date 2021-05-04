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
	var tasks []entities.Task
	if err := storage.db.Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

// GetTask ...
func (storage TaskStorage) GetTask(id uint) (entities.Task, error) {
	var task entities.Task
	if err := storage.db.Joins("User").First(&task, id).Error; err != nil {
		return entities.Task{}, err
	}

	return task, nil
}

// CreateTask ...
func (storage TaskStorage) CreateTask(task entities.Task) (id uint, err error) {
	task.Model = gorm.Model{} // reset the fields that are filled in automatically
	if err := storage.db.Create(&task).Error; err != nil {
		return 0, err
	}

	return task.ID, nil
}

// UpdateTask ...
func (storage TaskStorage) UpdateTask(id uint, task entities.Task) error {
	task.Model = gorm.Model{} // reset the fields that are filled in automatically
	return storage.db.
		Model(&entities.Task{Model: gorm.Model{ID: id}}).
		Updates(task).
		Error
}

// DeleteTask ...
func (storage TaskStorage) DeleteTask(id uint) error {
	return storage.db.Delete(&entities.Task{}, id).Error
}
