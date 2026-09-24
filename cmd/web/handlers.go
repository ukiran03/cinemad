package main

import (
	"fmt"
	"net/http"

	"ukiran03.com/cinemad/cmd/web/ui/html/pages"
	"ukiran03.com/cinemad/internal/models"
)

func (app *application) VirtualTerminal(
	w http.ResponseWriter,
	r *http.Request,
) {
	data := app.NewTemplateData(r)
	data.StringMap["publishable_key"] = app.config.Stripe.Key

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

func (app *application) PaymentSucceeded(
	w http.ResponseWriter,
	r *http.Request,
) {
	err := r.ParseForm()
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(
			w,
			fmt.Sprintf("Error: %v\n", err),
			http.StatusInternalServerError,
		)
		return
	}

	receipt := models.PaymentReceiptData{
		PaymentIntent: r.Form.Get("payment_intent"),
		Cardholder:    r.Form.Get("cardholder_name"),
		Email:         r.Form.Get("cardholder_email"),
		PaymentMethod: r.Form.Get("payment_method"),
		Amount:        r.Form.Get("payment_amount"),
		Currency:      r.Form.Get("payment_currency"),
	}

	data := app.NewTemplateData(r)

	err = pages.PaymentSucceededPage(receipt, data).Render(r.Context(), w)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(
			w,
			fmt.Sprintf("Error: %v\n", err),
			http.StatusInternalServerError,
		)
		return
	}
}
