package models

import "time"

type FullPayment struct {
	ID            int64   `json:"id"`
	OrderID       int64   `json:"order_id"`
	UserID        int64   `json:"user_id"`
	Amount        float64 `json:"amount"`
	PaymentMethod string  `json:"payment_method"`
	Status        string  `json:"status"`

	Code PaymentCode `json:"code"`

	PaidAt    *time.Time `json:"paid_at"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
