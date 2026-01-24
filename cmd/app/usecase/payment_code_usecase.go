package usecase

import (
	"context"
	"paymentmc/infrastructure/log"
	"paymentmc/models"
	"paymentmc/utils"
)

func (uc *PaymentUsecase) GetFullPaymentCodeByCode(
	ctx context.Context,
	codePayment string,
) (*models.FullPaymentCode, error) {
	dataFullPaymentCode, err := uc.PaymentService.GetFullPaymentCodeByCode(ctx, codePayment)

	if err != nil {
		log.Logger.Error("uc.PaymentUsecase.GetFullPaymentCodeByCode")
		return nil, err
	}

	if dataFullPaymentCode == nil {
		return nil, utils.ErrPaymentCodeNotFound
	}

	return dataFullPaymentCode, nil
}
