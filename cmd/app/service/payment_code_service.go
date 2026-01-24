package service

import (
	"context"
	"paymentmc/infrastructure/log"
	"paymentmc/models"
)

func (s *PaymentService) CreatePaymentCode(ctx context.Context, param models.PaymentCode) (int64, error) {
	id, err := s.PaymentRepository.InsertPaymentCode(ctx, param)
	if err != nil {
		log.Logger.Error("s.PaymentRepository.InsertPaymentCode")
		return 0, err
	}

	return id, nil
}

func (s *PaymentService) GetPaymentCodeByID(ctx context.Context, id int64) (*models.PaymentCode, error) {
	dataPaymentCode, err := s.PaymentRepository.GetPaymentCodeByID(ctx, id)
	if err != nil {
		log.Logger.Error("s.PaymentRepository.GetPaymentCodeByID")
		return nil, err
	}

	return dataPaymentCode, nil
}

func (s *PaymentService) GetFullPaymentCodeByCode(
	ctx context.Context,
	codePayment string,
) (*models.FullPaymentCode, error) {
	dataFullPaymentCode, err := s.PaymentRepository.GetFullPaymentCodeByCode(ctx, codePayment)
	if err != nil {
		log.Logger.Error("s.PaymentRepository.GetFullPaymentCodeByCode")
		return nil, err
	}

	return dataFullPaymentCode, nil
}
