package utils

import "errors"

// error order
var (
	OrderNotFound          = errors.New("order tidak ditemukan")
	OrderExpiredOrNotFound = errors.New("code order sudah kadaluarsa atau tidak ada")

	ErrStatusNotCompatible        = errors.New("status tidak diperbolehkan")
	ErrStatusNotCreated           = errors.New("status order harus created")
	ErrPaymentCodeNotFound        = errors.New("payment code tidak ditemukan")
	ErrStatusPaymentShouldPending = errors.New("payment status harus pending")

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
