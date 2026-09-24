package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Movie is the type for all movies
type Movie struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Year           int       `json:"year"`
	Description    string    `json:"description"`
	InventoryLevel int       `json:"inventory_level"`
	Price          int       `json:"price"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
}

type MovieModel struct {
	DB *pgxpool.Pool
}

// Models is the wrapper for all the models
type Models struct {
	Movies MovieModel
}

func NewModels(db *pgxpool.Pool) Models {
	return Models{
		Movies: MovieModel{DB: db},
	}
}
