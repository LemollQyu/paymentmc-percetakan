package service

import (
	"context"
	"paymentmc/infrastructure/log"
	"paymentmc/models"
)

// service create payment method
func (s *PaymentService) CreatePaymentMethod(ctx context.Context, param models.PaymentMethodRequest) error {
	err := s.PaymentRepository.InsertPaymentMethods(ctx, param)
	if err != nil {
		log.Logger.Error("s.PaymentRepository.InsertPaymentMethods")
		return err
	}

	return nil
}

// service get payment method by method
func (s *PaymentService) GetPaymentMethod(ctx context.Context, method string) (*models.PaymentMethod, error) {
	dataPaymentMethod, err := s.PaymentRepository.GetPaymentMethod(ctx, method)
	if err != nil {
		log.Logger.Error(" s.PaymentRepository.GetPaymentMethod")
		return nil, err
	}

	return dataPaymentMethod, nil
}

// service get payment method bg id
func (s *PaymentService) GetPaymentMethodByID(ctx context.Context, id int64) (*models.PaymentMethod, error) {
	dataPaymentMethod, err := s.PaymentRepository.GetPaymentMethodByID(ctx, id)
	if err != nil {
		log.Logger.Error(" s.PaymentRepository.GetPaymentMethod")
		return nil, err
	}

	return dataPaymentMethod, nil
}

// service get all payment method by method
func (s *PaymentService) GetAllPaymentMethod(ctx context.Context) (*[]models.PaymentMethod, error) {
	dataPaymentMethod, err := s.PaymentRepository.GetAllPaymentMethods(ctx)
	if err != nil {
		log.Logger.Error(" s.PaymentRepository.GetAllPaymentMethods")
		return nil, err
	}

	return dataPaymentMethod, nil
}

// service update payment method
func (s *PaymentService) UpdatePaymentMethod(ctx context.Context, id int64, param models.PaymentMethodRequest) error {
	err := s.PaymentRepository.UpdatePaymentMethod(ctx, id, param)
	if err != nil {
		log.Logger.Error("s.PaymentRepository.UpdatePaymentMethod")
		return err
	}

	return nil
}

// service delete payment method
func (s *PaymentService) DeletePaymentMethod(ctx context.Context, id int64) error {
	err := s.PaymentRepository.DeletePaymentMethod(ctx, id)
	if err != nil {
		log.Logger.Error("s.PaymentRepository.DeletePaymentMethod")
		return err
	}

	return nil
}
