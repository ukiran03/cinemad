package config

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type Config struct {
	WebPort int
	APIPort int
	Env     string
	API     string
	DB      struct {
		DSN string
	}
	Stripe struct {
		Key    string
		Secret string
	}
}

func Load() (*Config, error) {
	var cfg Config

	// custom FlagSet so errors can be returned instead of calling os.Exit
	fs := flag.NewFlagSet("config", flag.ContinueOnError)

	fs.IntVar(
		&cfg.WebPort,
		"web-port",
		4000,
		"Frotend Server port to listen on",
	)
	fs.IntVar(
		&cfg.APIPort,
		"api-port",
		4001,
		"Backend Server port to listen on",
	)
	fs.StringVar(
		&cfg.Env,
		"env",
		"devel",
		"Application environment {devel|prod}",
	)
	fs.StringVar(&cfg.API, "api", "http://localhost:4001", "URL to API")

	// parse flags and handle any parsing errors early
	if err := fs.Parse(os.Args[1:]); err != nil {
		return nil, fmt.Errorf("failed to parse flags: %w", err)
	}

	cfg.DB.DSN = os.Getenv("DB_DSN")
	cfg.Stripe.Key = os.Getenv("STRIPE_KEY")
	cfg.Stripe.Secret = os.Getenv("STRIPE_SECRET")

	// TODO:
	// if cfg.DB.DSN == "" {
	// 	return nil, errors.New("missing ENV variable: DB_DSN")
	// }

	if cfg.Env == "prod" && cfg.Stripe.Secret == "" {
		return nil, errors.New(
			"missing ENV variable: STRIPE_SECRET (required in production)",
		)
	}

	return &cfg, nil
}
