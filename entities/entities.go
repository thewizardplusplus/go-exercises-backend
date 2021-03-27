package entities

import (
	"github.com/go-gorm/datatypes"
	"gorm.io/gorm"
)

// Task ...
type Task struct {
	gorm.Model

	Title           string
	Description     string
	BoilerplateCode string
	TestCases       datatypes.JSON
}

// Solution ...
type Solution struct {
	gorm.Model

	TaskID    uint
	Task      Task
	Code      string
	IsCorrect bool
	Result    string `gorm:"type:json"`
}
