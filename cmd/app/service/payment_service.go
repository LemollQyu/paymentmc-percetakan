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
