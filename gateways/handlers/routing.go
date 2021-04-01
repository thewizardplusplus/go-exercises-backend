package handlers

import (
	"net/http"
	"time"

	"github.com/go-log/log"
	"github.com/gorilla/mux"
	"github.com/thewizardplusplus/go-exercises-backend/entities"
)

// RouterOptions ...
type RouterOptions struct {
	TokenSigningKey string
	TokenTTL        time.Duration
}

// RouterDependencies ...
type RouterDependencies struct {
	UserGetter       UserGetter
	TaskStorage      TaskStorage
	SolutionStorage  SolutionStorage
	SolutionRegister entities.SolutionRegister
	Logger           log.Logger
}

// NewRouter ...
func NewRouter(
	options RouterOptions,
	dependencies RouterDependencies,
) *mux.Router {
	rootRouter := mux.NewRouter()
	apiRouter := rootRouter.PathPrefix("/api/v1").Subrouter()
	apiRouterWithAuthorization := apiRouter.NewRoute().Subrouter()

	authorizationMiddleware :=
		AuthorizationMiddleware(options.TokenSigningKey, dependencies.Logger)
	apiRouterWithAuthorization.Use(authorizationMiddleware)

	tokenHandler := TokenHandler{
		TokenSigningKey: options.TokenSigningKey,
		TokenTTL:        options.TokenTTL,
		UserGetter:      dependencies.UserGetter,
		Logger:          dependencies.Logger,
	}
	apiRouter.
		HandleFunc("/tokens/", tokenHandler.CreateToken).
		Methods(http.MethodPost)

	taskHandler := TaskHandler{
		TaskStorage: dependencies.TaskStorage,
		Logger:      dependencies.Logger,
	}
	apiRouterWithAuthorization.
		HandleFunc("/tasks/{id}", taskHandler.GetTask).
		Methods(http.MethodGet)
	apiRouterWithAuthorization.
		HandleFunc("/tasks/{id}", taskHandler.UpdateTask).
		Methods(http.MethodPut)
	apiRouterWithAuthorization.
		HandleFunc("/tasks/{id}", taskHandler.DeleteTask).
		Methods(http.MethodDelete)
	apiRouterWithAuthorization.
		HandleFunc("/tasks/", taskHandler.GetTasks).
		Methods(http.MethodGet)
	apiRouterWithAuthorization.
		HandleFunc("/tasks/", taskHandler.CreateTask).
		Methods(http.MethodPost)

	solutionHandler := SolutionHandler{
		SolutionStorage:  dependencies.SolutionStorage,
		SolutionRegister: dependencies.SolutionRegister,
		Logger:           dependencies.Logger,
	}
	apiRouterWithAuthorization.
		HandleFunc("/tasks/{taskID}/solutions/", solutionHandler.GetSolutions).
		Methods(http.MethodGet)
	apiRouterWithAuthorization.
		HandleFunc("/tasks/{taskID}/solutions/", solutionHandler.CreateSolution).
		Methods(http.MethodPost)
	apiRouterWithAuthorization.
		HandleFunc("/solutions/{id}", solutionHandler.GetSolution).
		Methods(http.MethodGet)

	return rootRouter
}
