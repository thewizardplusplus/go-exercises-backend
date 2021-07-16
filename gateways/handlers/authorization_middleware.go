package handlers

import (
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-log/log"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/thewizardplusplus/go-exercises-backend/entities"
	httputils "github.com/thewizardplusplus/go-http-utils"
)

type userContextKey struct{}

// AuthorizationMiddleware ...
func AuthorizationMiddleware(
	tokenSigningKey string,
	logger log.Logger,
) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(
			writer http.ResponseWriter,
			request *http.Request,
		) {
			tokenAsStr :=
				strings.TrimPrefix(request.Header.Get("Authorization"), "Bearer ")
			token, err := jwt.ParseWithClaims(
				tokenAsStr,
				&entities.AccessTokenClaims{},
				func(token *jwt.Token) (interface{}, error) {
					return []byte(tokenSigningKey), nil
				},
			)
			if err != nil {
				err = errors.Wrap(err, "[error] failed token checking")
				httputils.LoggingError(logger, writer, err, http.StatusUnauthorized)

				return
			}

			tokenClaims := token.Claims.(*entities.AccessTokenClaims)
			request = request.WithContext(context.WithValue(
				request.Context(),
				userContextKey{},
				tokenClaims.User,
			))

			next.ServeHTTP(writer, request)
		})
	}
}
