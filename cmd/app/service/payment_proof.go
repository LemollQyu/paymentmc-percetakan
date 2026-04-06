package service

import (
	"context"
	"paymentmc/infrastructure/log"
	"paymentmc/models"
)

func (s *PaymentService) GetPaymentProofByID(ctx context.Context, PaymentID int64) (*models.BuktiPembayaran, error) {
	data, err := s.PaymentRepository.GetProofPaymentByID(ctx, PaymentID)
	if err != nil {
		log.Logger.Error("s.PaymentRepository.GetProofPaymentByID")
		return nil, err
	}

	return data, nil
}
