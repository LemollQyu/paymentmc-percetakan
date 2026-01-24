package repository

import (
	"context"
	"errors"
	"paymentmc/models"
	"paymentmc/utils/tables"

	"gorm.io/gorm"
)

// repo insert payment method
func (r *PaymentRepository) InsertPaymentMethods(ctx context.Context, param models.PaymentMethodRequest) error {
	err := r.Database.WithContext(ctx).
		Table(tables.PaymentMethods).
		Create(&param).Error

	if err != nil {
		return err
	}

	return nil
}

// repo get payment method by name method
func (r *PaymentRepository) GetPaymentMethod(ctx context.Context, method string) (*models.PaymentMethod, error) {
	var methodPayment models.PaymentMethod
	err := r.Database.WithContext(ctx).
		Where("payment_method = ?", method).
		First(&methodPayment).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &methodPayment, nil
}

// repo get payment method by id
func (r *PaymentRepository) GetPaymentMethodByID(ctx context.Context, id int64) (*models.PaymentMethod, error) {
	var methodPayment models.PaymentMethod
	err := r.Database.WithContext(ctx).
		Where("id = ?", id).
		First(&methodPayment).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &methodPayment, nil
}

// repo get all payment method
func (r *PaymentRepository) GetAllPaymentMethods(
	ctx context.Context,
) (*[]models.PaymentMethod, error) {

	var paymentMethods []models.PaymentMethod

	err := r.Database.WithContext(ctx).
		Find(&paymentMethods).
		Error

	if err != nil {
		return nil, err
	}

	return &paymentMethods, nil
}

// repo update payment method
func (r *PaymentRepository) UpdatePaymentMethod(
	ctx context.Context,
	id int64,
	param models.PaymentMethodRequest,
) error {

	err := r.Database.WithContext(ctx).
		Model(&models.PaymentMethod{}).
		Where("id = ?", id).
		Updates(param).
		Error

	return err
}

// repo delete payment method
func (r *PaymentRepository) DeletePaymentMethod(ctx context.Context, id int64) error {
	// Menggunakan GORM
	// Menghapus data berdasarkan ID
	result := r.Database.WithContext(ctx).
		Delete(&models.PaymentMethod{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
