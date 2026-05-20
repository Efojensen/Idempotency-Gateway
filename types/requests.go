package types

type PaymentRequest struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}
