package main

import (
	"fmt"
	"net/http"

	"ukiran03.com/cinemad/cmd/web/ui/html/pages"
)

func (app *application) VirtualTerminal(
	w http.ResponseWriter,
	r *http.Request,
) {
	data := app.NewTemplateData(r)
	err := pages.VirtualTerminalPage(data).Render(r.Context(), w)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Error: %v\n", err),
			http.StatusInternalServerError,
		)
		return
	}

	app.logger.Info("Hit the handler")
}
