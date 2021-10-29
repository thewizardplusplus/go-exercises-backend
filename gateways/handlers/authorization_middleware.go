package handlers

import (
	"net/http"

	"github.com/go-log/log"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/thewizardplusplus/go-exercises-backend/entities"
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
				err = errors.Wrap(err, "unable to parse the token")
				logError(logger, writer, err, autodetectedStatusCode)

				return
			}

			request = setUserToRequest(request, tokenClaims.User)
			next.ServeHTTP(writer, request)
		})
	}
}
