package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/Efojensen/Idempotency-Gateway/utils"
)

func CheckIdempotencyKey(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		idempotencyKey := r.Header.Get("Idempotency-Key")

		if idempotencyKey == "" {
			utils.WriteErrorResponse(w, http.StatusBadRequest, errors.New("missing idempotency key"))
			return
		}

		ctx := context.WithValue(r.Context(), "idempotencyKey", idempotencyKey)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}