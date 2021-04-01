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
	TaskStorage      TaskStorage
	SolutionStorage  SolutionStorage
	SolutionRegister entities.SolutionRegister
	UserGetter       UserGetter
	Logger           log.Logger
}

// NewRouter ...
func NewRouter(
	options RouterOptions,
	dependencies RouterDependencies,
) *mux.Router {
	rootRouter := mux.NewRouter()
	apiRouter := rootRouter.PathPrefix("/api/v1").Subrouter()

	taskHandler := TaskHandler{
		TaskStorage: dependencies.TaskStorage,
		Logger:      dependencies.Logger,
	}
	apiRouter.
		HandleFunc("/tasks/{id}", taskHandler.GetTask).
		Methods(http.MethodGet)
	apiRouter.
		HandleFunc("/tasks/{id}", taskHandler.UpdateTask).
		Methods(http.MethodPut)
	apiRouter.
		HandleFunc("/tasks/{id}", taskHandler.DeleteTask).
		Methods(http.MethodDelete)
	apiRouter.
		HandleFunc("/tasks/", taskHandler.GetTasks).
		Methods(http.MethodGet)
	apiRouter.
		HandleFunc("/tasks/", taskHandler.CreateTask).
		Methods(http.MethodPost)

	solutionHandler := SolutionHandler{
		SolutionStorage:  dependencies.SolutionStorage,
		SolutionRegister: dependencies.SolutionRegister,
		Logger:           dependencies.Logger,
	}
	apiRouter.
		HandleFunc("/tasks/{taskID}/solutions/", solutionHandler.GetSolutions).
		Methods(http.MethodGet)
	apiRouter.
		HandleFunc("/tasks/{taskID}/solutions/", solutionHandler.CreateSolution).
		Methods(http.MethodPost)
	apiRouter.
		HandleFunc("/solutions/{id}", solutionHandler.GetSolution).
		Methods(http.MethodGet)

	tokenHandler := TokenHandler{
		TokenSigningKey: options.TokenSigningKey,
		TokenTTL:        options.TokenTTL,
		UserGetter:      dependencies.UserGetter,
		Logger:          dependencies.Logger,
	}
	apiRouter.
		HandleFunc("/tokens/", tokenHandler.CreateToken).
		Methods(http.MethodPost)

	return rootRouter
}
