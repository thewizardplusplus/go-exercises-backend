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
	apiRouter := rootRouter.PathPrefix("/api").Subrouter()
	apiV1Router := apiRouter.PathPrefix("/v1").Subrouter()

	taskHandler := TaskHandler{
		TaskStorage: dependencies.TaskStorage,
		Logger:      dependencies.Logger,
	}
	tasksAPIV1Router := apiV1Router.PathPrefix("/tasks").Subrouter()
	tasksAPIV1Router.
		HandleFunc("/{id}", taskHandler.GetTask).
		Methods(http.MethodGet)
	tasksAPIV1Router.
		HandleFunc("/{id}", taskHandler.UpdateTask).
		Methods(http.MethodPut)
	tasksAPIV1Router.
		HandleFunc("/{id}", taskHandler.DeleteTask).
		Methods(http.MethodDelete)
	tasksAPIV1Router.
		HandleFunc("/", taskHandler.GetTasks).
		Methods(http.MethodGet)
	tasksAPIV1Router.
		HandleFunc("/", taskHandler.CreateTask).
		Methods(http.MethodPost)

	return rootRouter
}
