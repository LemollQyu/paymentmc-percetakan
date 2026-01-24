package usecase

import (
	"paymentmc/cmd/app/service"
	"paymentmc/cmd/app/storage"
)

type PaymentUsecase struct {
	PaymentService service.PaymentService
	StorageService storage.Storage
}

func NewPaymentUsecase(
	paymentService service.PaymentService,
	storageService storage.Storage,
) *PaymentUsecase {
	return &PaymentUsecase{
		PaymentService: paymentService,
		StorageService: storageService,
	}
}
