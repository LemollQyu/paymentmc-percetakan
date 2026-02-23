package models

import "time"

type BuktiPembayaran struct {
	ID        int64  `json:"id" gorm:"primaryKey"`
	PaymentID int64  `json:"payment_id"`
	ProofURL  string `json:"proof_url"`
	Note      string `json:"note"`

	UploadedAt time.Time  `json:"uploaded_at" gorm:"autoCreateTime"`
	VerifiedAt *time.Time `json:"verified_at"`
	CreatedAt  time.Time  `json:"created_at" gorm:"autoCreateTime"`
}

type RequestBuktiPembayaran struct {
	PaymentID int64  `form:"payment_id"`
	Note      string `form:"note" binding:"required"`
	ProofURL  string `form:"proof_url"`
}

// response uploaded storage, bukti pembayaran
type ResponseUploadBuktiPembayaran struct {
	UploadAt time.Time `json:"uploaded_at"`
}
