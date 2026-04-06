package repository

import (
	"context"
	"errors"
	"paymentmc/models"
	"paymentmc/utils"
	"paymentmc/utils/tables"
	"time"

	"gorm.io/gorm"
)

func (r *PaymentRepository) InsertPayment(
	ctx context.Context,
	param *models.Payment,
) (int64, error) {

	payment := &models.Payment{
		OrderID:       param.OrderID,
		UserID:        param.UserID,
		Amount:        param.Amount,
		PaymentMethod: param.PaymentMethod,
		Status:        param.Status,
		PaidAt:        nil,
	}

	err := r.Database.WithContext(ctx).
		Table(tables.Payments).
		Create(payment).Error

	if err != nil {
		return 0, err
	}

	return payment.ID, nil
}

func (r *PaymentRepository) GetPaymentByID(ctx context.Context, id int64) (*models.Payment, error) {
	var payment models.Payment
	err := r.Database.WithContext(ctx).
		Where("id = ?", id).
		First(&payment).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &payment, nil
}

// get payment by orderid
func (r *PaymentRepository) GetPaymentByOrderID(ctx context.Context, orderID int64) (*models.Payment, error) {
	var payment models.Payment

	err := r.Database.WithContext(ctx).
		Preload("PaymentCodes").
		Where("order_id = ?", orderID).
		First(&payment).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &payment, nil
}

// delete payment by order id
func (r *PaymentRepository) DeletePaymentByOrderID(ctx context.Context, orderID int64) error {

	return r.Database.WithContext(ctx).
		Where("order_id = ?", orderID).
		Delete(&models.Payment{}).
		Error
}

// payment code yang sudah expired, dan status Payment Pending dan ambil datanya
func (r *PaymentRepository) GetExpiredPendingPaymentIDs(
	ctx context.Context,
) ([]int64, int64, error) {

	var (
		paymentIDs []int64
		total      int64
	)

	query := r.Database.WithContext(ctx).
		Table("payment_codes pc").
		Joins("JOIN payments p ON p.id = pc.payment_id").
		Where("pc.expired_at < NOW()").
		Where("p.status = ?", utils.StatusPaymentPending).
		Select("pc.payment_id")

	// hitung total payment_id expired
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// ambil kumpulan payment_id
	if err := query.Pluck("pc.payment_id", &paymentIDs).Error; err != nil {
		return nil, 0, err
	}

	return paymentIDs, total, nil
}

// repo ambil order_id dan kumpulkan
func (r *PaymentRepository) GetOrderIDsByPaymentIDs(
	ctx context.Context,
	paymentIDs []int64,
) ([]int64, error) {

	var orderIDs []int64

	if len(paymentIDs) == 0 {
		return orderIDs, nil
	}

	err := r.Database.WithContext(ctx).
		Model(&models.Payment{}).
		Where("id IN ?", paymentIDs).
		Pluck("order_id", &orderIDs).Error

	if err != nil {
		return nil, err
	}

	return orderIDs, nil
}

func (r *PaymentRepository) UpdateExpiredPayment(
	ctx context.Context,
	paymentIDs []int64,
) error {
	if len(paymentIDs) == 0 {
		return errors.New("payment ids kosong")
	}

	return r.Database.
		WithContext(ctx).
		Table(tables.Payments).
		Where("id IN ?", paymentIDs).
		Updates(map[string]interface{}{
			"status":     utils.StatusPaymentExpired,
			"updated_at": time.Now(),
		}).Error
}

func (r *PaymentRepository) UpdatePayment(
	ctx context.Context,
	paymentID int64,
	newStatus string,
) error {

	return r.Database.
		WithContext(ctx).
		Table(tables.Payments).
		Where("id = ?", paymentID).
		Updates(map[string]interface{}{
			"status":     newStatus,
			"updated_at": time.Now(),
		}).Error
}

// repo update approved at payment

func (r *PaymentRepository) UpdateApproved(ctx context.Context, paymentID int64) error {
	return r.Database.
		WithContext(ctx).
		Table(tables.Payments).
		Where("id = ?", paymentID).
		Updates(map[string]interface{}{
			"approved_at": time.Now(),
			"updated_at":  time.Now(),
		}).Error
}

// repo update rejected at payment

func (r *PaymentRepository) UpdateRejected(ctx context.Context, paymentID int64) error {
	return r.Database.
		WithContext(ctx).
		Table(tables.Payments).
		Where("id = ?", paymentID).
		Updates(map[string]interface{}{
			"status":     utils.StatusPaymentCancelled,
			"updated_at": time.Now(),
		}).Error
}

// create bukti pembayaran

func (r *PaymentRepository) CreateProofPayment(
	ctx context.Context,
	param models.RequestBuktiPembayaran,
) (*models.ResponseUploadBuktiPembayaran, error) {

	paymentProof := &models.BuktiPembayaran{
		PaymentID: param.PaymentID,
		ProofURL:  param.ProofURL,
		Note:      param.Note,
	}

	err := r.Database.WithContext(ctx).
		Table(tables.PaymentProofs).
		Create(paymentProof).Error
	if err != nil {
		return nil, err
	}

	return &models.ResponseUploadBuktiPembayaran{
		UploadAt: paymentProof.UploadedAt,
	}, nil

}

// repo update paid payment
func (r *PaymentRepository) UpdatedPaidAt(ctx context.Context, paymentID int64) error {
	return r.Database.
		WithContext(ctx).
		Table(tables.Payments).
		Where("id = ?", paymentID).
		Updates(map[string]interface{}{
			"status":     utils.StatusPaymentSuccess,
			"paid_at":    time.Now(),
			"updated_at": time.Now(),
		}).Error
}

// repo get all payments
func (r *PaymentRepository) GetPayments(
	ctx context.Context,
	status string,
	limit int,
	offset int,
) ([]*models.Payment, error) {

	var payments []*models.Payment

	db := r.Database.WithContext(ctx).
		Table(tables.Payments).
		Preload("PaymentCodes")

	// filter status jika ada
	if status != "" {
		db = db.Where("payments.status = ?", status)
	}

	err := db.
		Order("payments.created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&payments).Error

	if err != nil {
		return nil, err
	}

	return payments, nil
}
