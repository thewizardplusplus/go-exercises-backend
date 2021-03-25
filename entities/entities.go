package entities

import (
	"gorm.io/gorm"
)

// Task ...
type Task struct {
	gorm.Model

	Title           string
	Description     string
	BoilerplateCode string
	TestCases       string
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
