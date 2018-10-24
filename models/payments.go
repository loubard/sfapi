package models

// FetchResponse holds one payment
type FetchResponse struct {
	Data *Payment `json:"data"`
}

// Payment represent a payment resource
type Payment struct {
	PaymentID string `json:"payment_id"`
}
