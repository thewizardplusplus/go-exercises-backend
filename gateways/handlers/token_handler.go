package handlers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-log/log"
	"github.com/pkg/errors"
	"github.com/thewizardplusplus/go-exercises-backend/entities"
	httputils "github.com/thewizardplusplus/go-http-utils"
)

// UserGetter ...
type UserGetter interface {
	GetUser(username string) (entities.User, error)
}

// TokenHandler ...
type TokenHandler struct {
	TokenSigningKey string
	TokenTTL        time.Duration
	UserGetter      UserGetter
	Logger          log.Logger
}

// CreateToken ...
func (handler TokenHandler) CreateToken(
	writer http.ResponseWriter,
	request *http.Request,
) {
	var user entities.User
	if err := httputils.ReadJSON(request.Body, &user); err != nil {
		err = errors.Wrap(err, "[error] unable to decode the user data")
		handler.Logger.Log(err)
		http.Error(writer, err.Error(), http.StatusBadRequest)

		return
	}

	foundUser, err := handler.UserGetter.GetUser(user.Username)
	if err != nil {
		err = errors.Wrap(err, "[error] unable to get the user")
		handler.Logger.Log(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return
	}

	if err := user.CheckPassword(foundUser.PasswordHash); err != nil {
		handler.Logger.Log(err)
		http.Error(writer, err.Error(), http.StatusUnauthorized)

		return
	}

	foundUser.PasswordHash = ""

	tokenExpirationTime := time.Now().Add(handler.TokenTTL).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, entities.AccessTokenClaims{
		StandardClaims: jwt.StandardClaims{ExpiresAt: tokenExpirationTime},
		User:           foundUser,
	})
	signedToken, err := token.SignedString([]byte(handler.TokenSigningKey))
	if err != nil {
		err = errors.Wrap(err, "[error] unable to create the token")
		handler.Logger.Log(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return
	}

	credentials := entities.Credentials{AccessToken: signedToken}
	httputils.WriteJSON(writer, http.StatusOK, credentials)
}
