package models

import "time"

type PaymentCode struct {
	ID        int64     `json:"id"`
	Code      string    `json:"code"`
	PaymentID int64     `json:"payment_id"`
	ExpiredAt time.Time `json:"expired_at"`
	CreatedAt time.Time `json:"created_at"`
}

type FullPaymentCode struct {
	ID        int64     `json:"id"`
	Code      string    `json:"code"`
	PaymentID int64     `json:"payment_id"`
	Payment   Payment   `json:"payment" gorm:"foreignKey:PaymentID;references:ID"`
	ExpiredAt time.Time `json:"expired_at"`
	CreatedAt time.Time `json:"created_at"`
}

func (FullPaymentCode) TableName() string {
	return "payment_codes"
}
