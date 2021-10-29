package handlers

import (
	"net/http"

	"github.com/go-log/log"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/thewizardplusplus/go-exercises-backend/entities"
	httputils "github.com/thewizardplusplus/go-http-utils"
)

// TokenParser ...
type TokenParser interface {
	ParseToken(authorizationHeader string) (*entities.AccessTokenClaims, error)
}

// AuthorizationMiddleware ...
func AuthorizationMiddleware(
	tokenParser TokenParser,
	logger log.Logger,
) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(
			writer http.ResponseWriter,
			request *http.Request,
		) {
			authorizationHeader := request.Header.Get("Authorization")
			tokenClaims, err := tokenParser.ParseToken(authorizationHeader)
			if err != nil {
				statusCode := http.StatusInternalServerError
				if errors.Is(err, entities.ErrFailedTokenChecking) {
					statusCode = http.StatusUnauthorized
				}

				err = errors.Wrap(err, "[error] unable to parse the token")
				httputils.LoggingError(logger, writer, err, statusCode)

				return
			}

			request = setUserToRequest(request, tokenClaims.User)
			next.ServeHTTP(writer, request)
		})
	}
}
