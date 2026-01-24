package models

import "time"

type Order struct {
	ID                  int64                 `json:"id"`
	UserID              int64                 `json:"user_id"`
	ServiceID           int64                 `json:"service_id"`
	ServiceNameSnapshot string                `json:"service_name_snapshot"`
	BasePriceSnapshot   int64                 `json:"base_price_snapshot"`
	TotalPriceSnapshot  int64                 `json:"total_price_snapshot"`
	UserNote            string                `json:"user_note"`
	Status              string                `json:"status"`
	Quantity            int32                 `json:"quantity"`
	User                *User                 `json:"user"`
	OrderCode           OrderCode             `json:"order_code"`
	OrderFile           *OrderFile            `json:"order_file"`
	OrderSpesifications []*OrderSpesification `json:"order_spesifications"`
}

type User struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Phone     *string `json:"phone"`
	AvatarUrl *string `json:"avatar_url"`
}

type OrderCode struct {
	ID      int64  `json:"id"`
	OrderID int64  `json:"order_id"`
	Code    string `json:"code"`
}

type OrderFile struct {
	ID       int64  `json:"id"`
	OrderID  int64  `json:"order_id"`
	FileUrl  string `json:"file_url"`
	Type     string `json:"type"`
	FileType string `json:"file_type"`
}

type OrderSpesification struct {
	ID                        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID                   int64     `json:"order_id"`
	SpesificationID           int64     `json:"spesification_id"`
	SpesificationNameSnapshot string    `json:"spesification_name_snapshot"`
	ValueSnapshot             string    `json:"value_snapshot"`
	AdditionalPriceSnapshot   int64     `json:"additional_price_snapshot"`
	CreatedAt                 time.Time `json:"created_at"`
}
