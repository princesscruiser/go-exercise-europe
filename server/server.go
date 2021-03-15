package server

import (
	"context"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	log "github.com/rs/zerolog/log"
)

type ServConfig struct {
	Port int    `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

type Config struct {
	ServConfig ServConfig `mapstructure:"servConfig"`
}

func New() *http.Server {
	return &http.Server{}
}

func Configure(srv *http.Server, cnf *ServConfig, router *mux.Router) error {
	srv.Addr = cnf.Host + ":" + strconv.Itoa(cnf.Port)
	srv.Handler = router
	return nil
}

func Start(ctx context.Context, srv *http.Server) {
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Msg(err.Error())
		}
	}()
}

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}
