package models

import "time"

func (RejectedPayment) TableName() string {
	return "rejected_payment"
}

func (Refund) TableName() string {
	return "refunds"
}

func (RefundProof) TableName() string {
	return "refund_proofs"
}

type RequestRejectedPayment struct {
	PaymentID int64 `json:"payment_id" `
	UserID    int64 `json:"user_id"`
	Amount    int64 `json:"amount"`

	OrderCode   string `json:"order_code" binding:"required"`
	PaymentCode string `json:"payment_code"`
	OrderName   string `json:"order_name" binding:"required"`

	AdminNote *string `json:"admin_note,omitempty"`
}

type RejectedPayment struct {
	ID int64 `gorm:"primaryKey;autoIncrement" json:"id"`

	PaymentID int64 `gorm:"not null" json:"payment_id"`
	UserID    int64 `gorm:"not null" json:"user_id"`
	Amount    int64 `gorm:"not null" json:"amount"`

	OrderCode   string `gorm:"type:varchar(50);not null" json:"order_code"`
	PaymentCode string `gorm:"type:varchar(50);not null" json:"payment_code"`
	OrderName   string `gorm:"type:varchar(200);not null" json:"order_name"`

	AdminNote *string `gorm:"type:text" json:"admin_note,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	// Relation
	Refunds []Refund `gorm:"foreignKey:RejectedID;constraint:OnDelete:CASCADE" json:"refunds"`
}

type RequestRefund struct {
	BankName      string `gorm:"type:varchar(100);not null" binding:"required" json:"bank_name"`
	AccountNumber string `gorm:"type:varchar(100);not null" binding:"required" json:"account_number"`
	AccountName   string `gorm:"type:varchar(150);not null" binding:"required" json:"account_name"`

	Status string `gorm:"type:varchar(50);default:requested" json:"status"`

	TransferredAt *time.Time `json:"transferred_at"`
}

type Refund struct {
	ID int64 `gorm:"primaryKey;autoIncrement" json:"id"`

	RejectedID int64 `gorm:"not null" json:"rejected_id"`

	BankName      string `gorm:"type:varchar(100);not null" json:"bank_name"`
	AccountNumber string `gorm:"type:varchar(100);not null" json:"account_number"`
	AccountName   string `gorm:"type:varchar(150);not null" json:"account_name"`

	Status string `gorm:"type:varchar(50);default:requested" json:"status"`

	TransferredAt *time.Time `json:"transferred_at,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	// Relation
	// RejectedPayment RejectedPayment `gorm:"foreignKey:RejectedID;constraint:OnDelete:CASCADE" json:"rejected_payment,omitempty"`
	Proofs []RefundProof `gorm:"foreignKey:RefundID;constraint:OnDelete:CASCADE" json:"proofs,omitempty"`
}

type RefundProof struct {
	ID int64 `gorm:"primaryKey;autoIncrement" json:"id"`

	RefundID int64 `gorm:"not null" json:"refund_id"`

	FileURL string  `gorm:"type:text;not null" json:"file_url"`
	Note    *string `gorm:"type:text" json:"note,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	// Relation
	Refund Refund `gorm:"foreignKey:RefundID;constraint:OnDelete:CASCADE" json:"refund,omitempty"`
}

type RequestRefundProof struct {
	FileURL string  `gorm:"type:text;not null" form:"file_url" json:"file_url"`
	Note    *string `gorm:"type:text" form:"note" binding:"required" json:"note"`
}
