package usecases

import (
	"github.com/pkg/errors"
	"github.com/thewizardplusplus/go-exercises-backend/entities"
	"gorm.io/gorm"
)

// TaskStorage ...
type TaskStorage interface {
	entities.TaskGetter

	GetTasks(userID uint, pagination entities.Pagination) ([]entities.Task, error)
	CountTasks() (int64, error)
	CreateTask(task entities.Task) (id uint, err error)
	UpdateTask(id uint, task entities.Task) error
	DeleteTask(id uint) error
}

// TaskUsecase ...
type TaskUsecase struct {
	TaskStorage TaskStorage
}

// GetTasks ...
func (usecase TaskUsecase) GetTasks(
	userID uint,
	pagination entities.Pagination,
) (
	entities.TaskGroup,
	error,
) {
	tasks, err := usecase.TaskStorage.GetTasks(userID, pagination)
	if err != nil {
		return entities.TaskGroup{}, errors.Wrap(err, "unable to get the tasks")
	}

	taskCount, err := usecase.TaskStorage.CountTasks()
	if err != nil {
		return entities.TaskGroup{}, errors.Wrap(err, "unable to count the tasks")
	}

	taskGroup := entities.TaskGroup{Tasks: tasks, TotalCount: taskCount}
	for index := range taskGroup.Tasks {
		taskGroup.Tasks[index].User.PasswordHash = ""
	}

	return taskGroup, nil
}

// GetTask ...
func (usecase TaskUsecase) GetTask(
	userID uint,
	taskID uint,
) (
	entities.Task,
	error,
) {
	task, err := usecase.TaskStorage.GetTask(userID, taskID)
	if err != nil {
		return entities.Task{}, errors.Wrap(err, "unable to get the task")
	}

	task.User.PasswordHash = ""
	return task, nil
}

// CreateTask ...
func (usecase TaskUsecase) CreateTask(
	userID uint,
	task entities.Task,
) (
	entities.Task,
	error,
) {
	task.UserID = userID
	if err := task.FormatBoilerplateCode(); err != nil {
		return entities.Task{},
			errors.Wrap(err, "unable to format the boilerplate code")
	}

	id, err := usecase.TaskStorage.CreateTask(task)
	if err != nil {
		return entities.Task{}, errors.Wrap(err, "unable to create the task")
	}

	idAsModel := entities.Task{Model: gorm.Model{ID: id}}
	return idAsModel, nil
}

// UpdateTask ...
func (usecase TaskUsecase) UpdateTask(
	userID uint,
	taskID uint,
	task entities.Task,
) error {
	if err := usecase.checkAccessToTask(userID, taskID); err != nil {
		return errors.Wrap(err, "unable to check access to the task")
	}

	if err := task.FormatBoilerplateCode(); err != nil {
		return errors.Wrap(err, "unable to format the boilerplate code")
	}

	if err := usecase.TaskStorage.UpdateTask(taskID, task); err != nil {
		return errors.Wrap(err, "unable to update the task")
	}

	return nil
}

// DeleteTask ...
func (usecase TaskUsecase) DeleteTask(userID uint, taskID uint) error {
	if err := usecase.checkAccessToTask(userID, taskID); err != nil {
		return errors.Wrap(err, "unable to check access to the task")
	}

	if err := usecase.TaskStorage.DeleteTask(taskID); err != nil {
		return errors.Wrap(err, "unable to delete the task")
	}

	return nil
}

func (usecase TaskUsecase) checkAccessToTask(
	userID uint,
	taskID uint,
) error {
	task, err := usecase.TaskStorage.GetTask(userID, taskID)
	if err != nil {
		return errors.Wrap(err, "unable to get the task")
	}

	if userID != task.UserID {
		return entities.ErrManagerialAccessIsDenied
	}

	return nil
}
