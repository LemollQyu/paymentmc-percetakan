package service

import (
	"context"
	"paymentmc/infrastructure/log"
	"paymentmc/models"
)

func (s *PaymentService) InsertListWaitingPayment(ctx context.Context, param *models.ListWaitingPayment) (int64, error) {
	id, err := s.PaymentRepository.CreateListWaitingPayment(ctx, *param)
	if err != nil {
		log.Logger.Error("CreateListWaitingPayment")
		return 0, err
	}

	return id, nil
}

func (s *PaymentService) GetWaitingPayment(
	ctx context.Context,
	code string,
) (*models.ListWaitingPayment, error) {
	data, err := s.PaymentRepository.GetWaitingPayment(ctx, code)
	if err != nil {
		log.Logger.Error(" GetWaitingPayment")
		return nil, err
	}

	return data, nil
}

func (s *PaymentService) GetListWaitingPayment(
	ctx context.Context,
) (*[]models.ListWaitingPayment, error) {
	data, err := s.PaymentRepository.GetListWaitingPayment(ctx)
	if err != nil {
		log.Logger.Error(" GetListWaitingPayment")
		return nil, err
	}

	return data, nil
}

func (s *PaymentService) GetListWaitingPaymentByUserID(
	ctx context.Context,
	userID int64,
) (*[]models.ListWaitingPayment, error) {
	data, err := s.PaymentRepository.GetListWaitingPaymentByUserID(ctx, userID)
	if err != nil {
		log.Logger.Error("GetListWaitingPaymentByUserID")
		return nil, err
	}
	return data, nil
}
