package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

const (
	StatusOrderWaitingPayment = "Waiting_payment"
	StatusOrderOnProgress     = "On_progress"
	StatusOrderExpired        = "Expired"
	StatusOrderCancelled      = "Cancelled"
	StatusOrderPaid           = "Paid"
	StatusOrderCreated        = "Created"
)

const (
	StatusPaymentPending   = "Pending"
	StatusPaymentExpired   = "Expired"
	StatusPaymentCancelled = "Cancelled"
	StatusPaymentSuccess   = "Success" // paid
)

func GeneratePaymentCode() (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 6

	result := make([]byte, length)
	for i := 0; i < length; i++ {
		// Mengambil index random dari charset menggunakan crypto/rand
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[num.Int64()]
	}

	return fmt.Sprintf("payment_%s", string(result)), nil
}

func Expire12Hour() time.Time {
	// Menambahkan durasi 2 jam ke waktu sekarang
	return time.Now().Add(2 * time.Hour)
}
