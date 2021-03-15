package main

import (
	"context"
	"go-exercise-europe/server"
	"go-exercise-europe/server/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	log "github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func main() {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Msg("Starting Server")

	setupConfigurator()
	conf := &server.Config{}
	if err := viper.Unmarshal(conf); err != nil {
		log.Fatal().Msg("Could not unmarshall config file")
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if conf.DebugLevel == "debug" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	rt := router.New()
	if err := router.SetHandlers(rt); err != nil {
		log.Fatal().Msg("Fatal error during router configuration")
	}
	srv := server.New()
	if err := server.Configure(srv, &conf.ServConfig, rt); err != nil {
		log.Fatal().Msg("Fatal error during server configuration")
	}
	ctx := context.Background()
	server.Start(srv)
	defer gracefullTerminate(ctx, srv)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Info().Str("Host", conf.ServConfig.Host).Int("Port", conf.ServConfig.Port).Msg("Successfully started")
	reason := <-done
	log.Warn().Str("Closing due", reason.String()).Msg("Server is about to be stopped")
}

func gracefullTerminate(ctx context.Context, srv *http.Server) {
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Msg(err.Error())
	}
	log.Warn().Msg("Server Shutdown Succeeded")
}

func setupConfigurator() {
	// Config
	viper.SetConfigName("config") // config file name without extension
	viper.SetConfigType("json")
	viper.AddConfigPath("../config/") // config file path
	viper.AutomaticEnv()              // read value ENV variable

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
}
