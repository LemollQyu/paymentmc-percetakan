package repository

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
