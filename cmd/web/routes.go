package main

import (
	"io/fs"
	"net/http"

	"github.com/go-chi/chi/v5"
	"ukiran03.com/cinemad/cmd/web/ui"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	// A sub-file system for the embedded "static" directory
	staticFS, err := fs.Sub(ui.Files, "static")
	if err != nil {
		panic(err)
	}

	// Serve the embedded static files
	mux.Handle(
		"/static/*",
		http.StripPrefix("/static", http.FileServer(http.FS(staticFS))),
	)

	// Application routes
	mux.Get("/vt", app.VirtualTerminal)
	mux.Post("/payment-succeeded", app.PaymentSucceeded)

	return mux
}
