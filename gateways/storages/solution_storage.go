package storages

import (
	"errors"

	"github.com/go-gorm/datatypes"
	"github.com/thewizardplusplus/go-exercises-backend/entities"
	"gorm.io/gorm"
)

// SolutionStorage ...
type SolutionStorage struct {
	db *gorm.DB
}

// NewSolutionStorage ...
func NewSolutionStorage(db *gorm.DB) SolutionStorage {
	return SolutionStorage{db: db}
}

// GetSolutions ...
func (storage SolutionStorage) GetSolutions(
	userID uint,
	taskID uint,
	pagination entities.Pagination,
) ([]entities.Solution, error) {
	query := storage.
		makeSolutionQuery(userID, taskID).
		Joins("User").
		Order("created_at DESC")
	if !pagination.IsZero() {
		query = query.Offset(pagination.Offset()).Limit(pagination.PageSize)
	}

	var solutions []entities.Solution
	if err := query.Find(&solutions).Error; err != nil {
		return nil, err
	}

	return solutions, nil
}

// CountSolutions ...
func (storage SolutionStorage) CountSolutions(
	userID uint,
	taskID uint,
) (int64, error) {
	var count int64
	err := storage.makeSolutionQuery(userID, taskID).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

// GetSolution ...
func (storage SolutionStorage) GetSolution(id uint) (entities.Solution, error) {
	var solution entities.Solution
	err := storage.db.Joins("Task").Joins("User").First(&solution, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = entities.ErrNotFound
		}
		return entities.Solution{}, err
	}

	return solution, nil
}

// CreateSolution ...
func (storage SolutionStorage) CreateSolution(
	taskID uint,
	solution entities.Solution,
) (id uint, err error) {
	// reset the fields that are filled in automatically
	solution.Model = gorm.Model{}
	solution.TaskID = taskID
	solution.IsCorrect = false
	solution.Result = datatypes.JSON("{}") // empty JSON object

	if err := storage.db.Create(&solution).Error; err != nil {
		return 0, err
	}

	return solution.ID, nil
}

// UpdateSolution ...
func (storage SolutionStorage) UpdateSolution(
	id uint,
	solution entities.Solution,
) error {
	// reset the fields that are filled in automatically
	solution.Model = gorm.Model{}

	return storage.db.
		Model(&entities.Solution{Model: gorm.Model{ID: id}}).
		Updates(solution).
		Error
}

func (storage SolutionStorage) makeSolutionQuery(
	userID uint,
	taskID uint,
) *gorm.DB {
	return storage.db.
		Model(&entities.Solution{}).
		Joins("Task").
		Where(
			storage.db.
				Where(&entities.Solution{TaskID: taskID}).
				Where(
					storage.db.
						Where(&entities.Solution{UserID: userID}).
						Or(map[string]interface{}{"Task.user_id": userID}),
				),
		)
}
