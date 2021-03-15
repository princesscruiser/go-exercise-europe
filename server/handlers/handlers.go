package handlers

import (
	"encoding/json"
	"net/http"

	"go-exercise-europe/server/services"

	log "github.com/rs/zerolog/log"
)

var (
	QServ services.QService
	TServ services.TranslateService
)

func Ping(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Ping requested")
	w.WriteHeader(200)
	w.Write([]byte("Pong"))
}

func GetQuestions(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("GetQuestions requested")

	langs, ok := r.URL.Query()["lang"]

	if !ok || len(langs[0]) < 1 {
		log.Error().Msg("Url Param 'key' is missing")
		return
	}
	lang := langs[0]
	log.Info().Str("lang:", lang).Msg("Lang Requested")

	questions, err := QServ.Find("555")
	if err != nil {
		log.Error().Msg(err.Error())
	}

	res, err := json.Marshal(questions)
	if err != nil {
		log.Error().Msg("Url Param 'key' is missing")
	}
	w.WriteHeader(200)
	w.Write([]byte(res))
}

func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("CreateQuestions requested")
}
