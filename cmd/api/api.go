package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
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

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4001, "Server port to listen on")
	flag.StringVar(
		&cfg.env,
		"env", "devel",
		"Application environment {devel|prod|maintenance}",
	)
	flag.Parse()

	cfg.stripe.pubkey = os.Getenv("STRIPE_KEY")
	cfg.stripe.secret = os.Getenv("STRIPE_SECRET")

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

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

func (app *application) serve() error {
	srv := &http.Server{
		Addr:        fmt.Sprintf(":%d", app.config.port),
		Handler:     app.routes(),
		IdleTimeout: 30 * time.Second,
		ReadTimeout: 10 * time.Second,
	}
	app.logger.Info(
		"Starting Backend server",
		"port", app.config.port,
		"mode", app.config.env,
	)

	return srv.ListenAndServe()
}
