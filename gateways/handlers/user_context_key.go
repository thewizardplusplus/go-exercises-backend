package handlers

import (
	"context"
	"net/http"

	"github.com/thewizardplusplus/go-exercises-backend/entities"
)

type userContextKey struct{}

func getUserFromRequest(request *http.Request) entities.User {
	return request.Context().Value(userContextKey{}).(entities.User)
}

func setUserToRequest(request *http.Request, user entities.User) *http.Request {
	return request.WithContext(context.WithValue(
		request.Context(),
		userContextKey{},
		user,
	))
}
