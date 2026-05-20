package middleware

import "net/http"

func CheckIdempotencyKey(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		r.Header.Get("Idempotency-Key")

		next.ServeHTTP(w, r)
	})
}