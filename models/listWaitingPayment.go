package models

import "time"

type ListWaitingPayment struct {
	ID        int64 `json:"id"`
	PaymentID int64 `json:"payment_id"`
	OrderID   int64 `json:"order_id"`
	UserID    int64 `json:"user_id"`

	Amount            float64 `json:"amount"`
	OrderCode         string  `json:"order_code"`
	IconMethodPayment string  `json:"icon_method_payment"`
	NumberPayment     string  `json:"number_payment"`
	CodeQris          string  `json:"code_qris"`

	CheckoutAt time.Time `json:"checkout_at"`
	ExpiredAt  time.Time `json:"expired_at"`
	CreatedAt  time.Time `json:"created_at"`

	Payment Payment `json:"payment" gorm:"-"` // tambah ini
}
