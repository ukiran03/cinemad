package main

import (
	"net/http"
	"time"

	"ukiran03.com/cinemad/internal/models"
)

func (app *application) NewTemplateData(r *http.Request) *models.TemplateData {
	return &models.TemplateData{
		StringMap:   make(map[string]string),
		IntMap:      make(map[string]int),
		FloatMap:    make(map[string]float32),
		Data:        make(map[string]interface{}),
		CurrentYear: time.Now().Year(),
		API:         app.config.API,
		CSSVersion:  cssVersion,
	}
}
