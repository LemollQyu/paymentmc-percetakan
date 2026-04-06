package utils

import "errors"

// error order
var (
	OrderNotFound          = errors.New("order tidak ditemukan")
	OrderExpiredOrNotFound = errors.New("code order sudah kadaluarsa atau tidak ada")

	ErrStatusNotCompatible        = errors.New("status tidak diperbolehkan")
	ErrStatusNotCreated           = errors.New("status order harus created")
	ErrPaymentCodeNotFound        = errors.New("payment code tidak ditemukan")
	ErrStatusPaymentShouldPending = errors.New("payment status harus pending atau sudah expired")
	ErrPaymentPaid                = errors.New("payment sudah dibayar")
	ErrPaymentShouldSuccess       = errors.New("payment status harus success")
	ErrOrderShouldPaid            = errors.New("order status harus paid")
	ErrPaymentNotFound            = errors.New("payment tidak ditemukan")
	ErrNotDeletePayment           = errors.New("payment tidak bisa dihapus")

	ErrStatusOrderShouldWaitingPayment = errors.New("order status harus waiting_payment")
)

// error payment method
var (
	PaymentMethodExist    = errors.New("payment method sudah ada")
	PaymentMethodNotExist = errors.New("payment method tidak ada")
)

// error storage

var (
	FileMaxSize    = errors.New("ukuran file terlalu besar")
	FileExtInvalid = errors.New("ekstensi file tidak diperbolehkan")
	FileRequired   = errors.New("file wajib diisi")
	ErrDeleteFile  = errors.New("gagal menghapus file")
)

// error refunc
var (
	ErrRejectPaymentNotFound = errors.New("ID tidak ditemukan")
	ErrDataRefundIsFound     = errors.New("data refund sudah ada")
)
