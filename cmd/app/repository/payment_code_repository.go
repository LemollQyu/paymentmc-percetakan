package repository

import (
	"context"
	"errors"
	"paymentmc/models"
	"paymentmc/utils/tables"

	"gorm.io/gorm"
)

func (r *PaymentRepository) InsertPaymentCode(ctx context.Context, param models.PaymentCode) (int64, error) {
	err := r.Database.WithContext(ctx).
		Table(tables.PaymentCodes).
		Create(&param).Error

	if err != nil {
		return 0, err
	}

	return param.ID, nil
}

func (r *PaymentRepository) GetPaymentCodeByID(ctx context.Context, id int64) (*models.PaymentCode, error) {
	var paymentCode models.PaymentCode
	err := r.Database.WithContext(ctx).
		Where("id = ?", id).
		First(&paymentCode).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &paymentCode, nil
}

func (r *PaymentRepository) GetFullPaymentCodeByCode(
	ctx context.Context,
	codePayment string,
) (*models.FullPaymentCode, error) {

	var fullPaymentCode models.FullPaymentCode

	err := r.Database.WithContext(ctx).
		Preload("Payment").
		Where("code = ?", codePayment).
		First(&fullPaymentCode).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &fullPaymentCode, nil
}
