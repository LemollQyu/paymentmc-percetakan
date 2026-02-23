package models

import "time"

// note semetara waktu saja

type RequestPayment struct {
	OrderID       int64   `json:"order_id"`
	OrderCode     string  `json:"order_code"`
	UserID        int64   `json:"user_id"`
	Amount        float64 `json:"amount"`
	PaymentMethod string  `json:"payment_method" binding:"required"`
}

type Payment struct {
	ID            int64       `json:"id"`
	OrderID       int64       `json:"order_id"`
	UserID        int64       `json:"user_id"`
	Amount        float64     `json:"amount"`
	PaymentMethod string      `json:"payment_method"`
	Status        string      `json:"status"`
	PaidAt        *time.Time  `json:"paid_at"`
	ApprovedAt    *time.Time  `json:"approved_at"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
	PaymentCodes  PaymentCode `gorm:"foreignKey:PaymentID" json:"payment_code"`
}

type ResponsePayment struct {
	OrderID       int64     `json:"order_id"`
	UserID        int64     `json:"user_id"`
	Amount        float64   `json:"amount"`
	OrderCode     string    `json:"order_code"`
	NumberPayment string    `json:"number_payment"`
	CodeQris      string    `json:"code_qris"`
	ExpiredAt     time.Time `json:"expired_at"`
}

type RequestProofPayment struct {
	UrlProofPayment string `form:"url_proof_payment"`
	Note            string `form:"note" binding:"required"`
	OrderID         int64  `form:"order_id"`
	UserID          int64  `form:"user_id"`
	PaymentID       int64  `form:"payment_id"`
}

type ResponseListWaitingPayment struct {
}
