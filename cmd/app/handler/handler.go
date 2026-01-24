package handler

import "paymentmc/cmd/app/usecase"

type PaymentHandler struct {
	PaymentUsecase usecase.PaymentUsecase
}

func NewPaymentHandler(paymentUsecase usecase.PaymentUsecase) *PaymentHandler {
	return &PaymentHandler{
		PaymentUsecase: paymentUsecase,
	}
}
