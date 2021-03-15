package server

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/rs/zerolog/log"
)

type ServConfig struct {
	Port int    `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

type Config struct {
	ServConfig ServConfig `mapstructure:"servConfig"`
	DebugLevel string     `mapstructure:"debugLevel"`
}

func New() *http.Server {
	return &http.Server{}
}

func Configure(srv *http.Server, cnf *ServConfig, router *mux.Router) error {
	srv.Addr = cnf.Host + ":" + strconv.Itoa(cnf.Port)
	srv.Handler = router
	return nil
}

func Start(srv *http.Server) {
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Msg(err.Error())
		}
	}()
}
