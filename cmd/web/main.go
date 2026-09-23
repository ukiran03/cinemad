package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"ukiran03.com/cinemad/internal/logger"
)

const (
	version    = "1.0.0"
	cssVersion = "1"
)

type config struct {
	port int
	env  string
	api  string
	db   struct {
		dsn string
	}
	stripe struct {
		secret string
		pubkey string
	}
}

type application struct {
	version string
	config  config
	logger  *slog.Logger
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:        fmt.Sprintf(":%d", app.config.port),
		Handler:     app.routes(),
		IdleTimeout: 30 * time.Second,
		ReadTimeout: 10 * time.Second,
	}
	app.logger.Info(
		"Starting HTTP server",
		"port", app.config.port,
		"mode", app.config.env,
	)

	return srv.ListenAndServe()
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(
		&cfg.env, "env", "devel", "Application environment {devel|prod}",
	)
	flag.StringVar(&cfg.api, "api", "http://localhost:4001", "URL to API")
	flag.Parse()

	cfg.stripe.pubkey = os.Getenv("STRIPE_KEY")
	cfg.stripe.secret = os.Getenv("STRIPE_SECRET")

	logger := logger.NewLogger()

	app := &application{
		config:  cfg,
		logger:  logger,
		version: version,
	}

	err := app.serve()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
