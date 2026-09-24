package models

type PaymentReceiptData struct {
	PaymentIntent string
	Cardholder    string
	Email         string
	PaymentMethod string
	Amount        string
	Currency      string
}
