package usecase

import (
	"context"
	"paymentmc/infrastructure/log"
	"paymentmc/models"
	"time"
)

func (uc *PaymentUsecase) SendNotificationOrder(ctx context.Context, param models.NotificationOrderRequest) error {

	// pakai context baru supaya tidak ikut timeout HTTP request
	notifCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := uc.PaymentService.Notification(notifCtx, param)
	if err != nil {
		log.Logger.Error("PaymentUsecase: SendNotificationOrder", err.Error())
		return err
	}

	return nil
}
