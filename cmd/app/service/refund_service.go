package service

import (
	"context"
	"paymentmc/infrastructure/log"
	"paymentmc/models"
)

func (s *PaymentService) CreateSubmitRefund(ctx context.Context, param models.RequestRejectedPayment) error {
	err := s.PaymentRepository.CreateSubmitRefund(ctx, param)
	if err != nil {
		log.Logger.Error("ServicePayment.CreateSubmitRefund")
		return err
	}

	return nil
}

func (s *PaymentService) GetMyRefund(
	ctx context.Context,
	userID int64,
) (*[]models.RejectedPayment, error) {

	data, err := s.PaymentRepository.GetMyRefund(ctx, userID)
	if err != nil {
		log.Logger.Error("GetMyRefund")
		return nil, err
	}

	return data, nil
}

func (s *PaymentService) GetAllRefunds(
	ctx context.Context,
) (*[]models.RejectedPayment, error) {

	data, err := s.PaymentRepository.GetAllRefund(ctx)
	if err != nil {
		log.Logger.Error("GetAllRefund")
		return nil, err
	}

	return data, nil
}

func (s *PaymentService) GetRejectPaymentByID(ctx context.Context, rejectID int64) (*models.RejectedPayment, error) {
	data, err := s.PaymentRepository.GetRejectPaymentByID(ctx, rejectID)
	if err != nil {
		log.Logger.Error("service.GetRejectPaymentByID")
		return nil, err
	}

	return data, nil
}

func (s *PaymentService) GetFullDetailRejectedPaymentByID(ctx context.Context, rejectID int64) (*models.RejectedPayment, error) {
	data, err := s.PaymentRepository.GetFullDetailRejectedPaymentByID(ctx, rejectID)
	if err != nil {
		log.Logger.Error("service.GetFullDetailRejectedPaymentByID")
		return nil, err
	}

	return data, nil

}

func (s *PaymentService) CreateRefund(ctx context.Context, refund *models.Refund) error {
	err := s.PaymentRepository.CreateRefund(ctx, refund)
	if err != nil {
		log.Logger.Error("service.CreateRefund")
		return err
	}

	return nil
}

func (s *PaymentService) GetRefundByRejectedID(ctx context.Context, rejectedID int64) (*models.Refund, error) {
	data, err := s.PaymentRepository.GetRefundByRejectedID(ctx, rejectedID)
	if err != nil {
		log.Logger.Error("service.GetRefund")
		return nil, err
	}

	return data, nil
}

func (s *PaymentService) GetRefundByID(ctx context.Context, refundID int64) (*models.Refund, error) {
	return s.PaymentRepository.GetRefundByID(ctx, refundID)
}

func (s *PaymentService) CreateRefundProof(ctx context.Context, proof *models.RefundProof) error {
	return s.PaymentRepository.CreateRefundProof(ctx, proof)
}

func (s *PaymentService) UpdateRefundStatus(ctx context.Context, refundID int64, status string) error {
	return s.PaymentRepository.UpdateRefundStatus(ctx, refundID, status)
}
