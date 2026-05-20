package types

type CachedResponse struct {
	StatusCode int
	PaymentRequest
	Body       string
}
