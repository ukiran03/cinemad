package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"ukiran03.com/cinemad/internal/cards"
)

type stripePayload struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type jsonResp struct {
	OK      bool   `json:"ok"`
	Message string `json:"message,omitempty"`
	Content string `json:"content,omitempty"`
	ID      int    `json:"id,omitempty"`
}

func (app *application) GetPaymentIntent(
	w http.ResponseWriter,
	r *http.Request,
) {
	var payload stripePayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(jsonResp{
			OK:      false,
			Message: "Invalid request payload",
		})
		return
	}

	amount, err := strconv.Atoi(payload.Amount)
	if err != nil {
		app.logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(jsonResp{
			OK:      false,
			Message: "Invalid amount format",
		})
		return
	}

	card := cards.Card{
		Secret:   app.config.Stripe.Secret,
		Key:      app.config.Stripe.Key,
		Currency: payload.Currency,
	}

	ok := true // for gateway charges
	pi, msg, err := card.Charge(payload.Currency, amount)
	if err != nil {
		ok = false
	}

	if ok {
		out, err := json.MarshalIndent(pi, "", "    ")
		if err != nil {
			app.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(jsonResp{
				OK:      false,
				Message: "Internal server error",
			})
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(out)
	} else {
		j := jsonResp{
			OK:      false,
			Message: msg,
			Content: "",
		}

		out, err := json.MarshalIndent(j, "", "    ")
		if err != nil {
			app.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(jsonResp{
				OK:      false,
				Message: "Internal server error",
			})
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(out)
	}
}
