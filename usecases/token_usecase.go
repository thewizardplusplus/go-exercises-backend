package usecases

import (
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"github.com/thewizardplusplus/go-exercises-backend/entities"
)

// UserGetter ...
type UserGetter interface {
	GetUser(username string) (entities.User, error)
}

// Clock ...
type Clock func() time.Time

// TokenUsecase ...
type TokenUsecase struct {
	TokenSigningKey string
	TokenTTL        time.Duration
	UserGetter      UserGetter
	Clock           Clock
}

// CreateToken ...
func (usecase TokenUsecase) CreateToken(
	user entities.User,
) (
	entities.Credentials,
	error,
) {
	foundUser, err := usecase.UserGetter.GetUser(user.Username)
	if err != nil {
		return entities.Credentials{}, errors.Wrap(err, "unable to get the user")
	}

	if foundUser.IsDisabled != nil && *foundUser.IsDisabled {
		return entities.Credentials{}, entities.ErrUserIsDisabled
	}

	if err := user.CheckPassword(foundUser.PasswordHash); err != nil {
		return entities.Credentials{}, err
	}
	foundUser.PasswordHash = ""

	tokenExpirationTime := usecase.Clock().Add(usecase.TokenTTL).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, entities.AccessTokenClaims{
		StandardClaims: jwt.StandardClaims{ExpiresAt: tokenExpirationTime},
		User:           foundUser,
	})
	signedToken, err := token.SignedString([]byte(usecase.TokenSigningKey))
	if err != nil {
		return entities.Credentials{}, errors.Wrap(err, "unable to create the token")
	}

	credentials := entities.Credentials{AccessToken: signedToken}
	return credentials, nil
}

// ParseToken ...
func (usecase TokenUsecase) ParseToken(
	authorizationHeader string,
) (
	*entities.AccessTokenClaims,
	error,
) {
	tokenAsStr := strings.TrimPrefix(authorizationHeader, "Bearer ")
	token, err := jwt.ParseWithClaims(
		tokenAsStr,
		&entities.AccessTokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(usecase.TokenSigningKey), nil
		},
	)
	if err != nil {
		return nil, multierror.Append(err, entities.ErrFailedTokenChecking)
	}

	tokenClaims := token.Claims.(*entities.AccessTokenClaims)
	return tokenClaims, nil
}
