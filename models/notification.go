package models

type NotificationOrderRequest struct {
	UserID           int64  `json:"user_id"`
	Type             string `json:"type"`
	TypeNotification string `json:"type_notification"`
	Subject          string `json:"subject"`
	Body             string `json:"body"`
	Name             string `json:"name"`
	OrderCode        string `json:"order_code"`
	ExpiredAt        string `json:"expired_at"`
	Email            string `json:"email"` //email target
	Service          string `json:"service"`
	Amount           string `json:"amount"`
	PaymentCode      string `json:"payment_code"`
}
