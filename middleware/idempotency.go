package middleware

import (
	"context"
	"net/http"
)

func CheckIdempotencyKey(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		idempotencyKey := r.Header.Get("Idempotency-Key")

		ctx := context.WithValue(r.Context(), "idempotencyKey", idempotencyKey)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}