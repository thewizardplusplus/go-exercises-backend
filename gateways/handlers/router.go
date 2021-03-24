package handlers

import (
	"net/http"

	"github.com/go-log/log"
	"github.com/gorilla/mux"
)

// RouterDependencies ...
type RouterDependencies struct {
	TaskStorage TaskStorage
	Logger      log.Logger
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

	return rootRouter
}
