package repository

import (
	"context"
	"errors"
	"paymentmc/models"
	"paymentmc/utils/tables"
	"time"

	"gorm.io/gorm"
)

func (r *PaymentRepository) CreateSubmitRefund(
	ctx context.Context,
	param models.RequestRejectedPayment,
) error {

	req := &models.RejectedPayment{
		PaymentID:   param.PaymentID,
		UserID:      param.UserID,
		Amount:      param.Amount,
		OrderCode:   param.OrderCode,
		PaymentCode: param.PaymentCode,
		OrderName:   param.OrderName,
		AdminNote:   param.AdminNote,
	}

	err := r.Database.WithContext(ctx).
		Table(tables.RejectedPayment).
		Create(req).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *PaymentRepository) GetMyRefund(
	ctx context.Context,
	userID int64,
	status string,
	page int,
	limit int,
) (*[]models.RejectedPayment, error) {

	var refunds []models.RejectedPayment

	offset := (page - 1) * limit

	query := r.Database.WithContext(ctx).
		Model(&models.RejectedPayment{}).
		Where("user_id = ?", userID)

	if status != "" {
		query = query.Where("EXISTS (SELECT 1 FROM refunds WHERE refunds.rejected_id = rejected_payment.id AND refunds.status = ?)", status)
	}

	err := query.
		Preload("Refunds", func(db *gorm.DB) *gorm.DB {
			if status != "" {
				return db.Where("status = ?", status)
			}
			return db
		}).
		Preload("Refunds.Proofs").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&refunds).Error

	if err != nil {
		return nil, err
	}

	return &refunds, nil
}

func (r *PaymentRepository) GetAllRefund(
	ctx context.Context,
	status string,
	page int,
	limit int,
) (*[]models.RejectedPayment, error) {

	var refunds []models.RejectedPayment

	offset := (page - 1) * limit

	query := r.Database.WithContext(ctx).
		Model(&models.RejectedPayment{})

	if status != "" {
		query = query.Where("EXISTS (SELECT 1 FROM refunds WHERE refunds.rejected_id = rejected_payment.id AND refunds.status = ?)", status)
	}

	err := query.
		Preload("Refunds", func(db *gorm.DB) *gorm.DB {
			if status != "" {
				return db.Where("status = ?", status)
			}
			return db
		}).
		Preload("Refunds.Proofs").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&refunds).Error

	if err != nil {
		return nil, err
	}

	return &refunds, nil
}

func (r *PaymentRepository) GetRefundByID(ctx context.Context, refundID int64) (*models.Refund, error) {
	var refund models.Refund

	err := r.Database.WithContext(ctx).
		Table(tables.Refounds).
		Where("id = ?", refundID).
		First(&refund).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &refund, nil
}

func (r *PaymentRepository) GetRejectPaymentByID(
	ctx context.Context,
	rejectID int64,
) (*models.RejectedPayment, error) {

	var refund models.RejectedPayment

	err := r.Database.WithContext(ctx).
		Table(tables.RejectedPayment).
		Where("id = ?", rejectID).
		First(&refund).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &refund, nil
}

func (r *PaymentRepository) GetFullDetailRejectedPaymentByID(ctx context.Context, rejectID int64) (*models.RejectedPayment, error) {
	var rejectedPayment models.RejectedPayment

	err := r.Database.WithContext(ctx).
		Model(&models.RejectedPayment{}).
		Where("id = ?", rejectID).
		Preload("Refunds").
		Preload("Refunds.Proofs").
		First(&rejectedPayment).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &rejectedPayment, nil
}

func (r *PaymentRepository) CreateRefund(ctx context.Context, refund *models.Refund) error {
	err := r.Database.WithContext(ctx).
		Create(refund).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *PaymentRepository) GetRefundByRejectedID(ctx context.Context, rejectedID int64) (*models.Refund, error) {
	var refund models.Refund

	err := r.Database.WithContext(ctx).
		Model(&models.Refund{}).
		Where("rejected_id = ?", rejectedID).
		Preload("Proofs").
		First(&refund).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &refund, nil
}

func (r *PaymentRepository) CreateRefundProof(ctx context.Context, proof *models.RefundProof) error {
	err := r.Database.WithContext(ctx).
		Table(tables.RefoundProofs).
		Create(proof).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *PaymentRepository) UpdateRefundStatus(ctx context.Context, refundID int64, status string) error {
	err := r.Database.WithContext(ctx).
		Table(tables.Refounds).
		Where("id = ?", refundID).
		Updates(map[string]interface{}{
			"status":         status,
			"transferred_at": time.Now(),
		}).Error

	if err != nil {
		return err
	}

	return nil
}
