package routes

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Efojensen/Idempotency-Gateway/types"
	"github.com/Efojensen/Idempotency-Gateway/utils"
)

type PaymentHandler struct {
	idem_key string
}

func (p *PaymentHandler) RegisterPaymentRoutes (h *http.ServeMux) {
	h.HandleFunc("/process-payment", p.payment)
}

func (p *PaymentHandler) payment(w http.ResponseWriter, r *http.Request) {
	var paymentBody types.PaymentRequest

	if err := json.NewDecoder(r.Body).Decode(&paymentBody); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if paymentBody.Amount == 0 || paymentBody.Currency == "" {
		utils.WriteErrorResponse(w, http.StatusBadRequest, errors.New("missing amount or currency"))
		return
	}
}