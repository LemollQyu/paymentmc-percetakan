package usecase

import (
	"context"
	"paymentmc/infrastructure/log"
	"paymentmc/models"
)

func (uc *PaymentUsecase) GetPaymentProofByID(ctx context.Context, PaymentID int64) (*models.BuktiPembayaran, error) {
	data, err := uc.PaymentService.GetPaymentProofByID(ctx, PaymentID)
	if err != nil {
		log.Logger.Error("s.PaymentService.GetProofPaymentByID")
		return nil, err
	}

	return data, nil
}
