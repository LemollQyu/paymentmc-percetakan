package repository

import (
	"context"
	"paymentmc/models"
	"paymentmc/utils/tables"
)

// // repo ambil order_id dan kumpulkan
// func (r *PaymentRepository) GetPaymentProofByPaymentID(
// 	ctx context.Context,
// 	paymentID int64,
// ) (*models.BuktiPembayaran, error) {

// 	err := r.Database.WithContext(ctx).
// 		Model(&models.Payment{}).
// 		Where("id IN ?", paymentIDs).
// 		Pluck("order_id", &orderIDs).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	return orderIDs, nil
// }

func (r *PaymentRepository) GetProofPaymentByID(
	ctx context.Context,
	id int64,
) (*models.BuktiPembayaran, error) {

	var proof models.BuktiPembayaran

	err := r.Database.WithContext(ctx).
		Table(tables.PaymentProofs).
		Where("payment_id = ?", id).
		First(&proof).Error
	if err != nil {
		return nil, err
	}

	return &proof, nil

}
