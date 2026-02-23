package repository

import (
	"context"
	"paymentmc/models"
	"paymentmc/utils"
	"paymentmc/utils/tables"
)

func (r *PaymentRepository) CreateListWaitingPayment(ctx context.Context, param models.ListWaitingPayment) (int64, error) {
	err := r.Database.WithContext(ctx).
		Table(tables.ListWaitingPayment).
		Create(&param).Error

	if err != nil {
		return 0, err
	}

	return param.ID, nil
}

func (r *PaymentRepository) GetListWaitingPayment(
	ctx context.Context,
) (*[]models.ListWaitingPayment, error) {

	var payments []models.ListWaitingPayment

	err := r.Database.WithContext(ctx).
		Table(tables.ListWaitingPayment).
		Find(&payments).Error

	if err != nil {
		return nil, err
	}

	return &payments, nil
}

// GetListWaitingPaymentByUserID mengambil data list payment waiting milik user,
// hanya payment dengan status Pending (order yang masih Waiting_payment).
func (r *PaymentRepository) GetListWaitingPaymentByUserID(
	ctx context.Context,
	userID int64,
) (*[]models.ListWaitingPayment, error) {

	var payments []models.ListWaitingPayment

	err := r.Database.WithContext(ctx).
		Table(tables.ListWaitingPayment+" lwp").
		Joins("JOIN "+tables.Payments+" p ON p.id = lwp.payment_id").
		Where("lwp.user_id = ? AND p.status = ?", userID, utils.StatusPaymentPending).
		Find(&payments).Error

	if err != nil {
		return nil, err
	}

	return &payments, nil
}
