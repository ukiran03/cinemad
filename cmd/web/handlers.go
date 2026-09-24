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

	// read posted data
	cardHolder := r.Form.Get("cardholder_name")
	email := r.Form.Get("cardholder_email")
	paymentIntent := r.Form.Get("payment_intent")
	paymentMethod := r.Form.Get("payment_method")
	paymentAmount := r.Form.Get("payment_amount")
	paymentCurrency := r.Form.Get("payment_currency")

	data := make(map[string]interface{})
	data["cardholder"] = cardHolder
	data["email"] = email
	data["pi"] = paymentIntent
	data["pm"] = paymentMethod
	data["pa"] = paymentAmount
	data["pc"] = paymentCurrency

	// render templates
	tData := app.NewTemplateData(r)
	tData.Data = data

	err = pages.PaymentSucceededPage(tData).Render(r.Context(), w)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Error: %v\n", err),
			http.StatusInternalServerError,
		)
		return
	}
}
