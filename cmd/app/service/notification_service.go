package service

import (
	"context"
	"paymentmc/infrastructure/log"
	"paymentmc/models"
)

func (s *PaymentService) Notification(ctx context.Context, param models.NotificationOrderRequest) error {

	_, err := s.NotificationClient.InsertNotificationOrder(ctx, param)
	if err != nil {
		// jangan return error, notif gagal tidak boleh ganggu flow utama
		log.Logger.Error("InsertNotificationOrder failed: ", err.Error())
		return err
	}

	return nil
}
