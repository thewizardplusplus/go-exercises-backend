package entities

import (
	"github.com/pkg/errors"
)

// ...
var (
	ErrManagerialAccessIsDenied = errors.New("managerial access is denied")
	ErrUserIsDisabled           = errors.New("user is disabled")
	ErrFailedPasswordChecking   = errors.New("failed password checking")
	ErrFailedTokenChecking      = errors.New("failed token checking")
	ErrUnableToFormatCode       = errors.New("unable to format the code")
	ErrNotFound                 = errors.New("not found")
)
