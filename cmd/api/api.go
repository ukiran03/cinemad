package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"ukiran03.com/cinemad/internal/config"
	"ukiran03.com/cinemad/internal/driver"
	"ukiran03.com/cinemad/internal/logger"
)

const version = "1.0.0"

type application struct {
	version string
	config  *config.Config
	logger  *slog.Logger
}

func main() {
	logger := logger.NewLogger()
	cfg, err := config.Load()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	conn, err := driver.OpenDB(cfg)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	app := &application{
		config:  cfg,
		logger:  logger,
		version: version,
	}

	err = app.serve()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:        fmt.Sprintf(":%d", app.config.APIPort),
		Handler:     app.routes(),
		IdleTimeout: 30 * time.Second,
		ReadTimeout: 10 * time.Second,
	}
	app.logger.Info(
		"Starting Backend server",
		"port", app.config.APIPort,
		"mode", app.config.Env,
	)

	return srv.ListenAndServe()
}
