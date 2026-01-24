package service

import (
	"paymentmc/cmd/app/repository"
	"paymentmc/grpc/order"
)

type PaymentService struct {
	PaymentRepository repository.PaymentRepository
	OrderClient       *order.OrderClient
}

func NewPaymentService(paymentRepository repository.PaymentRepository, orderClient *order.OrderClient) *PaymentService {
	return &PaymentService{
		PaymentRepository: paymentRepository,
		OrderClient:       orderClient,
	}
}
