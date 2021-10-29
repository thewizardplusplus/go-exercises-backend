package handlers

import (
	"net/http"
	"time"

	"github.com/go-log/log"
	"github.com/gorilla/mux"
	"github.com/thewizardplusplus/go-exercises-backend/entities"
	"github.com/thewizardplusplus/go-exercises-backend/usecases"
	httputils "github.com/thewizardplusplus/go-http-utils"
)

// RouterOptions ...
type RouterOptions struct {
	StaticFilePath  string
	TokenSigningKey string
	TokenTTL        time.Duration
}

// RouterDependencies ...
type RouterDependencies struct {
	UserGetter       usecases.UserGetter
	TaskStorage      usecases.TaskStorage
	SolutionStorage  usecases.SolutionStorage
	SolutionRegister entities.SolutionRegister
	Clock            usecases.Clock
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

	tokenUsecase := usecases.TokenUsecase{
		TokenSigningKey: options.TokenSigningKey,
		TokenTTL:        options.TokenTTL,
		UserGetter:      dependencies.UserGetter,
		Clock:           dependencies.Clock,
	}
	authorizationMiddleware :=
		AuthorizationMiddleware(tokenUsecase, dependencies.Logger)
	apiRouterWithAuthorization.Use(authorizationMiddleware)

	staticFileHandler := httputils.StaticAssetHandler(
		http.Dir(options.StaticFilePath),
		dependencies.Logger,
	)
	rootRouter.
		PathPrefix("/").Handler(staticFileHandler).
		Methods(http.MethodGet)

	tokenHandler := TokenHandler{
		TokenCreator: tokenUsecase,
		Logger:       dependencies.Logger,
	}
	apiRouter.
		HandleFunc("/tokens/", tokenHandler.CreateToken).
		Methods(http.MethodPost)

	taskHandler := TaskHandler{
		TaskUsecase: usecases.TaskUsecase{
			TaskStorage: dependencies.TaskStorage,
		},
		Logger: dependencies.Logger,
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
		SolutionUsecase: usecases.SolutionUsecase{
			TaskGetter:       dependencies.TaskStorage,
			SolutionStorage:  dependencies.SolutionStorage,
			SolutionRegister: dependencies.SolutionRegister,
		},
		Logger: dependencies.Logger,
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
	apiRouterWithAuthorization.
		HandleFunc("/solutions/format", solutionHandler.FormatSolution).
		Methods(http.MethodPost)

	return rootRouter
}
