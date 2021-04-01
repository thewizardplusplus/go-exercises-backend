package entities

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-gorm/datatypes"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Task ...
type Task struct {
	gorm.Model

	UserID          uint
	User            User
	Title           string
	Description     string
	BoilerplateCode string
	TestCases       datatypes.JSON
}

// Solution ...
type Solution struct {
	gorm.Model

	UserID    uint
	User      User
	TaskID    uint
	Task      Task
	Code      string
	IsCorrect bool
	Result    datatypes.JSON
}

// User ...
type User struct {
	gorm.Model

	Username     string `gorm:"unique"`
	Password     string `gorm:"-"`
	PasswordHash string
}

// CheckPassword ...
func (user User) CheckPassword(passwordHash string) error {
	if err := bcrypt.CompareHashAndPassword(
		[]byte(passwordHash),
		[]byte(user.Password),
	); err != nil {
		return errors.Wrap(err, "failed password checking")
	}

	return nil
}

// HashPassword ...
func (user *User) HashPassword(cost int) error {
	passwordHashBytes, err :=
		bcrypt.GenerateFromPassword([]byte(user.Password), cost)
	if err != nil {
		return errors.Wrap(err, "unable to hash the password")
	}

	user.PasswordHash = string(passwordHashBytes)
	return nil
}

// Credentials ...
type Credentials struct {
	AccessToken string
}

// AccessTokenClaims ...
type AccessTokenClaims struct {
	jwt.StandardClaims

	User User
}
