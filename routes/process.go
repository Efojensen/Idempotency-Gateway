package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Efojensen/Idempotency-Gateway/logging"
	"github.com/Efojensen/Idempotency-Gateway/types"
	"github.com/Efojensen/Idempotency-Gateway/utils"
)

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

	key := r.Context().Value("idempotencyKey").(string)

	p.mu.Lock()
	cachedResponse, exists := p.cache[key]

	if exists {
		p.mu.Unlock()
		if cachedResponse.Currency == paymentBody.Currency && cachedResponse.Amount == paymentBody.Amount {
			w.Header().Set("X-Cache-Hit", "true")
			utils.WriteResponse(w, p.cache[key].StatusCode, p.cache[key].Body)
			return
		} else {
			attackMsg := "WARN duplicate transaction request rejected: idempotency key already exists | idempotency_key="
			utils.WriteErrorResponse(w, http.StatusUnprocessableEntity,
				errors.New("Idempotency key already used for a different request body"),
			)
			logging.WriteAttackLogToFile(attackMsg, key)
			return
		}
	}

	time.Sleep(time.Second * 2)

	msg := fmt.Sprintf("Charged %d %s", paymentBody.Amount, paymentBody.Currency)

	p.cache[key] = types.CachedResponse{
		StatusCode:     201,
		Body:           msg,
		PaymentRequest: paymentBody,
	}
	p.mu.Unlock()

	err := logging.WriteLogToFile(msg, key)

	if err != nil {
		log.Println(err.Error())
	}

	utils.WriteResponse(w, http.StatusCreated, msg)
}
