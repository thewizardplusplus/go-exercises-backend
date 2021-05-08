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
func (storage TaskStorage) GetTasks(pagination entities.Pagination) (
	[]entities.Task,
	error,
) {
	query := storage.db.
		Select("tasks.*", "statuses.status").
		Joins("User").
		Joins(
			"JOIN (?) statuses ON statuses.tasks_id = tasks.id",
			storage.db.
				Model(&entities.Task{}).
				Select(
					"tasks.id AS tasks_id",
					`MAX(CASE
						WHEN is_correct THEN 2
						WHEN NOT is_correct AND result != '{}' THEN 1
						ELSE 0
					END) AS status`,
				).
				Joins("LEFT JOIN solutions ON solutions.task_id = tasks.id").
				Group("tasks.id"),
		).
		Order("created_at DESC")
	if !pagination.IsZero() {
		query = query.Offset(pagination.Offset()).Limit(pagination.PageSize)
	}

	var tasks []entities.Task
	if err := query.Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

// GetTask ...
func (storage TaskStorage) GetTask(id uint) (entities.Task, error) {
	var task entities.Task
	err := storage.db.
		Select("tasks.*", "statuses.status").
		Joins("User").
		Joins(
			"JOIN (?) statuses ON statuses.tasks_id = tasks.id",
			storage.db.
				Model(&entities.Task{}).
				Select(
					"tasks.id AS tasks_id",
					`MAX(CASE
						WHEN is_correct THEN 2
						WHEN NOT is_correct AND result != '{}' THEN 1
						ELSE 0
					END) AS status`,
				).
				Joins("LEFT JOIN solutions ON solutions.task_id = tasks.id").
				Group("tasks.id"),
		).
		First(&task, id).
		Error
	if err != nil {
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
