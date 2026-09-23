package main

import (
	"encoding/json"
	"net/http"
)

type stripePayload struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type jsonResp struct {
	OK      bool   `json:"ok"`
	Msg     string `json:"msg"`
	Content string `json:"content"`
	ID      int    `json:"id"`
}

func (app *application) GetPaymentIntent(
	w http.ResponseWriter,
	r *http.Request,
) {
	j := jsonResp{
		OK: true,
	}

	out, err := json.MarshalIndent(j, "", "   ")
	if err != nil {
		app.logger.Error(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
