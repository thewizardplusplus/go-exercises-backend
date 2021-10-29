package handlers

import (
	"net/http"

	"github.com/go-log/log"
	"github.com/pkg/errors"
	"github.com/thewizardplusplus/go-exercises-backend/entities"
	httputils "github.com/thewizardplusplus/go-http-utils"
)

// TokenCreator ...
type TokenCreator interface {
	CreateToken(user entities.User) (entities.Credentials, error)
}

// TokenHandler ...
type TokenHandler struct {
	TokenCreator TokenCreator
	Logger       log.Logger
}

// CreateToken ...
func (handler TokenHandler) CreateToken(
	writer http.ResponseWriter,
	request *http.Request,
) {
	var user entities.User
	if err := httputils.ReadJSON(request.Body, &user); err != nil {
		err = errors.Wrap(err, "[error] unable to decode the user data")
		httputils.LoggingError(handler.Logger, writer, err, http.StatusBadRequest)

		return
	}

	credentials, err := handler.TokenCreator.CreateToken(user)
	if err != nil {
		err = errors.Wrap(err, "[error] unable to create the token")
		statusCode := getStatusCodeFromError(err)
		httputils.LoggingError(handler.Logger, writer, err, statusCode)

		return
	}

	httputils.WriteJSON(writer, http.StatusOK, credentials)
}
