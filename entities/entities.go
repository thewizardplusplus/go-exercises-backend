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
