package models

import "time"

type PaymentMethodRequest struct {
	PaymentMethod string `form:"payment_method" binding:"required"`
	UrlIcon       string `form:"url_icon"`
	NumberPayment string `form:"number_payment"`
	UrlCode       string `form:"url_code"`
}

type PaymentMethod struct {
	ID            int64     `json:"id"`
	PaymentMethod string    `json:"payment_method"`
	NumberPayment string    `json:"number_payment"`
	UrlCode       string    `json:"url_code"`
	UrlIcon       string    `json:"url_icon"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
