package router

import (
	"go-exercise-europe/server/handlers"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	router := mux.NewRouter()
	return router
}

func SetHandlers(rt *mux.Router) error {
	rt.HandleFunc("/ping", handlers.Ping)
	return nil
}
