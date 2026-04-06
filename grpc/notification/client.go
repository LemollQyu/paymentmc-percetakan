package notification

import (
	"context"
	"log"
	"paymentmc/models"
	"paymentmc/proto/notificationpb"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type NotificationClient struct {
	Client notificationpb.NotificationServiceClient
}

func NewNotificationClient() *NotificationClient {
	conn, err := grpc.Dial(
		"localhost:50055", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed init grpc client: %v", err)
	}

	client := notificationpb.NewNotificationServiceClient(conn)
	return &NotificationClient{
		Client: client,
	}
}

func (nc *NotificationClient) InsertNotificationOrder(
	ctx context.Context,
	param models.NotificationOrderRequest,
) (int64, error) {

	res, err := nc.Client.InsertNotificationOrder(ctx, &notificationpb.InsertNotificationOrderRequest{
		UserId:           param.UserID,
		Type:             param.Type,
		TypeNotification: notificationpb.TypeNotification(notificationpb.TypeNotification_value[strings.ToUpper(param.TypeNotification)]),
		Name:             param.Name,
		OrderCode:        param.OrderCode,
		ExpiredAt:        param.ExpiredAt,
		Email:            param.Email,
		Service:          param.Service,
		Amount:           param.Amount,
		PaymentCode:      param.PaymentCode,
	})

	if err != nil {
		return 0, err
	}

	return res.Id, nil
}
