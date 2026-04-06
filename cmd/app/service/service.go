package service

import (
	"paymentmc/cmd/app/repository"
	"paymentmc/grpc/notification"
	"paymentmc/grpc/order"
	"paymentmc/grpc/user"
)

type PaymentService struct {
	PaymentRepository  repository.PaymentRepository
	OrderClient        *order.OrderClient
	NotificationClient *notification.NotificationClient
	UserClient         *user.UserClient
}

func NewPaymentService(paymentRepository repository.PaymentRepository, orderClient *order.OrderClient, notificationClient *notification.NotificationClient, userClient *user.UserClient) *PaymentService {
	return &PaymentService{
		PaymentRepository:  paymentRepository,
		OrderClient:        orderClient,
		NotificationClient: notificationClient,
		UserClient:         userClient,
	}
}
