package usecase

import (
	"context"
	"paymentmc/infrastructure/log"
	"paymentmc/models"

	"github.com/sirupsen/logrus"
)

func (uc *PaymentUsecase) GetOrderByCode(ctx context.Context, code string) (*models.Order, error) {
	dataOrder, err := uc.PaymentService.GetOrderByCode(ctx, code)

	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"code":  code,
			"error": err.Error(),
		})

		return nil, err
	}

	return dataOrder, nil
}
