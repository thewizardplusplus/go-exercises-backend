package handlers

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/thewizardplusplus/go-exercises-backend/entities"
)

func getStatusCodeFromError(err error) int {
	statusCode := http.StatusInternalServerError
	if errors.Is(err, entities.ErrUnableToFormatCode) {
		statusCode = http.StatusBadRequest
	}
	if errors.Is(err, entities.ErrFailedPasswordChecking) ||
		errors.Is(err, entities.ErrFailedTokenChecking) {
		statusCode = http.StatusUnauthorized
	}
	if errors.Is(err, entities.ErrManagerialAccessIsDenied) {
		statusCode = http.StatusForbidden
	}
	if errors.Is(err, entities.ErrNotFound) {
		statusCode = http.StatusNotFound
	}

	return statusCode
}
