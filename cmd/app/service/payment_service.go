package service

import (
	"context"
	"paymentmc/infrastructure/log"
	"paymentmc/models"
)

func (s *PaymentService) CreatePayment(ctx context.Context, param *models.Payment) (int64, error) {
	paymentId, err := s.PaymentRepository.InsertPayment(ctx, param)

	if err != nil {
		log.Logger.Error("s.PaymentRepository.InsertPayment")
		return 0, err
	}

	return paymentId, nil
}

func (s *PaymentService) GetPaymentByID(ctx context.Context, id int64) (*models.Payment, error) {
	dataPayment, err := s.PaymentRepository.GetPaymentByID(ctx, id)
	if err != nil {
		log.Logger.Error("s.PaymentRepository.GetPaymentMethodByID")
		return nil, err
	}

	return dataPayment, nil
}

func (s *PaymentService) GetExpiredPendingPaymentCodes(
	ctx context.Context,
) ([]int64, int64, error) {
	data, total, err := s.PaymentRepository.GetExpiredPendingPaymentIDs(ctx)
	if err != nil {
		log.Logger.Error("s.PaymentRepository.GetExpiredPendingPaymentCodes")
		return nil, 0, err
	}

	return data, total, nil
}

func (s *PaymentService) GetOrderIDsByPaymentIDs(
	ctx context.Context,
	paymentIDs []int64,
) ([]int64, error) {
	data, err := s.PaymentRepository.GetOrderIDsByPaymentIDs(ctx, paymentIDs)
	if err != nil {
		log.Logger.Error("s.PaymentRepository.GetOrderIDsByPaymentIDs")
		return nil, err
	}

	return data, nil
}

func (s *PaymentService) UpdateExpiredPayment(
	ctx context.Context,
	paymentIDs []int64,
) error {
	err := s.PaymentRepository.UpdateExpiredPayment(ctx, paymentIDs)
	if err != nil {
		log.Logger.Error("s.PaymentRepository.UpdateExpiredPayment")
		return err
	}

	return nil
}

func (s *PaymentService) UpdatePaymentStatus(ctx context.Context, paymentID int64, newStatus string) error {
	err := s.PaymentRepository.UpdatePayment(ctx, paymentID, newStatus)

	if err != nil {
		log.Logger.Error("s.PaymentRepository.UpdatePayment")
		return err
	}

	return nil
}

func (s *PaymentService) UpdateApproved(ctx context.Context, paymentID int64) error {
	err := s.PaymentRepository.UpdateApproved(ctx, paymentID)
	if err != nil {
		log.Logger.Error("s.PaymentRepository.UpdateApproved")
		return err
	}

	return nil
}

func (s *PaymentService) UpdateRejected(ctx context.Context, paymentID int64) error {
	err := s.PaymentRepository.UpdateRejected(ctx, paymentID)
	if err != nil {
		log.Logger.Error("s.PaymentRepository.UpdateRejected")
		return err
	}

	return nil
}

func (s *PaymentService) PaymentProof(ctx context.Context, param models.RequestBuktiPembayaran) (*models.ResponseUploadBuktiPembayaran, error) {
	uploadAt, err := s.PaymentRepository.CreateProofPayment(ctx, param)

	if err != nil {
		log.Logger.Error("s.PaymentRepository.CreateProofPayment")
		return nil, err
	}

	return uploadAt, nil
}

func (s *PaymentService) UpdatedPaidAt(ctx context.Context, paymentID int64) error {
	err := s.PaymentRepository.UpdatedPaidAt(ctx, paymentID)
	if err != nil {
		log.Logger.Error("s.PaymentRepository..UpdatedPaidAt")
		return err
	}

	return nil
}

func (s *PaymentService) GetPayments(
	ctx context.Context,
	status string,
	limit int,
	offset int,
) ([]*models.Payment, error) {

	// safety default limit
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	data, err := s.PaymentRepository.GetPayments(ctx, status, limit, offset)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// service get payment by order id
func (s *PaymentService) GetPaymentByOrderID(ctx context.Context, orderID int64) (*models.Payment, error) {
	data, err := s.PaymentRepository.GetPaymentByOrderID(ctx, orderID)
	if err != nil {

		log.Logger.Error("GetPaymentByOrderID")
		return nil, err
	}
	return data, nil
}

// service delete payment by order id
func (s *PaymentService) DeletePaymentByOrderID(ctx context.Context, orderID int64) error {
	err := s.PaymentRepository.DeletePaymentByOrderID(ctx, orderID)
	if err != nil {
		log.Logger.Error("DeletePaymentByOrderID")
		return err
	}

	return nil
}
