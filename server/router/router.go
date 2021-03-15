package router

import (
	"go-exercise-europe/clients"
	"go-exercise-europe/server/handlers"
	repository "go-exercise-europe/server/repositories"
	"go-exercise-europe/server/services"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func New() *mux.Router {
	router := mux.NewRouter()
	return router
}

func SetHandlers(rt *mux.Router) error {
	handlers.QServ = &services.QServImpl{
		QRepo: &repository.JsonRepo{},
	}
	handlers.TServ = &services.GTranslateService{
		GoogleClient: clients.GoogleClient{},
	}
	log.Debug().Msg("Routing...")
	rt.HandleFunc("/ping", handlers.Ping).Methods("GET")
	rt.HandleFunc("/questions", handlers.GetQuestions).Methods("GET")
	rt.HandleFunc("/questions", handlers.CreateQuestion).Methods("POST")
	return nil
}
