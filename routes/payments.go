package routes

import (
	"net/http"
	"sync"

	"github.com/Efojensen/Idempotency-Gateway/middleware"
	"github.com/Efojensen/Idempotency-Gateway/types"
)

type PaymentHandler struct {
	cache map[string]types.CachedResponse
	mu    sync.Mutex
}

func NewPaymentHandler(cache map[string]types.CachedResponse) *PaymentHandler {
	return &PaymentHandler{
		cache: cache,
	}
}

func (p *PaymentHandler) RegisterPaymentRoutes(h *http.ServeMux) {
	h.HandleFunc("/process-payment", middleware.CheckIdempotencyKey(p.payment))
}
