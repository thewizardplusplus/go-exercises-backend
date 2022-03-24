package handlers

import (
	"net/http"

	"github.com/go-log/log"
	"github.com/pkg/errors"
	"github.com/thewizardplusplus/go-exercises-backend/entities"
	httputils "github.com/thewizardplusplus/go-http-utils"
)

const autodetectedStatusCode = 0

func logError(
	logger log.Logger,
	writer http.ResponseWriter,
	err error,
	statusCode int,
) {
	err = errors.Wrap(err, "[error]")
	if statusCode == autodetectedStatusCode {
		statusCode = getStatusCodeFromError(err)
	}
	httputils.LoggingError(logger, writer, err, statusCode)
}

func getStatusCodeFromError(err error) int {
	statusCode := http.StatusInternalServerError
	if errors.Is(err, entities.ErrUnableToFormatCode) {
		statusCode = http.StatusBadRequest
	}
	if errors.Is(err, entities.ErrFailedPasswordChecking) ||
		errors.Is(err, entities.ErrFailedTokenChecking) {
		statusCode = http.StatusUnauthorized
	}
	if errors.Is(err, entities.ErrManagerialAccessIsDenied) ||
		errors.Is(err, entities.ErrUserIsDisabled) {
		statusCode = http.StatusForbidden
	}
	if errors.Is(err, entities.ErrNotFound) {
		statusCode = http.StatusNotFound
	}

	return statusCode
}
