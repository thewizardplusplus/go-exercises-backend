package handlers

import (
	"net/http"

	"github.com/go-log/log"
	"github.com/gorilla/mux"
)

// RouterDependencies ...
type RouterDependencies struct {
	TaskStorage      TaskStorage
	SolutionStorage  SolutionStorage
	SolutionRegister SolutionRegister
	Logger           log.Logger
}

// NewRouter ...
func NewRouter(dependencies RouterDependencies) *mux.Router {
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

	return rootRouter
}
