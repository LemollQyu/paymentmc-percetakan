package models

import "time"

type BuktiPembayaran struct {
	ID         int64     `json:"id"`
	PaymentID  int64     `json:"payment_id"`
	ProofURL   string    `json:"proof_url"`
	Note       string    `json:"note"`
	UploadedAt time.Time `json:"uploaded_at"`
	VerifiedAt time.Time `json:"verified_at"`
	CreatedAt  string    `json:"created_at"`
}

type RequestBuktiPembayaran struct {
	PaymentID  int64  `json:"payment_id"`
	Note       string `json:"note"`
	ProofURL   string `json:"proof_url"`
	UploadedAt string `json:"uploaded_at"`
}

// response uploaded storage, bukti pembayaran
type ResponseUploadBuktiPembayaran struct {
	URL  string    `json:"url"`
	Time time.Time `json:"uploaded_at"`
}
