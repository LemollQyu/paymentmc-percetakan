package order

import (
	"context"
	"paymentmc/infrastructure/log"
	"paymentmc/proto/orderpb"
	"paymentmc/utils"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (g *OrderClient) GetOrderDetailByID(
	ctx context.Context,
	orderID int64,
) (*orderpb.GetOrderDetailByIDResponse, error) {

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	order, err := g.Client.GetOrderDetailByID(
		ctx,
		&orderpb.GetOrderDetailByIDRequest{
			OrderId: orderID,
		},
	)

	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			switch st.Code() {

			case codes.Unknown:
				return nil, utils.OrderExpiredOrNotFound

			default:
				return nil, err
			}
		}
		return nil, err
	}

	return order, nil
}

func (g *OrderClient) GetOrderDetailByCode(
	ctx context.Context,
	code string,
) (*orderpb.GetOrderDetailResponse, error) {

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	order, err := g.Client.GetOrderDetailByCode(
		ctx,
		&orderpb.GetOrderDetailByCodeRequest{
			OrderCode: code,
		},
	)

	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			switch st.Code() {

			case codes.Unknown:
				return nil, utils.OrderExpiredOrNotFound

			default:
				return nil, err
			}
		}
		return nil, err
	}

	return order, nil
}

// grpc unutk update status
// hanya update flow update status apa, di usecase
func (g *OrderClient) UpdateOrderStatus(
	ctx context.Context,
	code string,
	status string,
) (*orderpb.UpdateOrderStatusResponse, error) {

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var orderStatus orderpb.OrderStatus

	switch status {
	case "Created":
		orderStatus = orderpb.OrderStatus_Created
	case "Waiting_payment":
		orderStatus = orderpb.OrderStatus_Waiting_payment
	case "Paid":
		orderStatus = orderpb.OrderStatus_Paid
	case "Expired":
		orderStatus = orderpb.OrderStatus_Expired
	case "On_progress":
		orderStatus = orderpb.OrderStatus_On_progress
	case "Finished":
		orderStatus = orderpb.OrderStatus_Finished
	case "Completed":
		orderStatus = orderpb.OrderStatus_Completed
	case "Cancelled":
		orderStatus = orderpb.OrderStatus_Cancelled
	default:
		return nil, utils.ErrStatusNotCompatible
	}

	ok, err := g.Client.UpdateOrderStatus(
		ctx,
		&orderpb.UpdateOrderStatusRequest{
			OrderCode: code,
			Status:    orderStatus,
		},
	)

	if err != nil {
		log.Logger.Errorf("g.Client.UpdateOrderStatus, %v", err.Error())
		return nil, err
	}

	return ok, nil
}

func (g *OrderClient) ExpiredOrderStatus(ctx context.Context, ids []int64) (*orderpb.BulkUpdateOrderStatusResponse, error) {
	ok, err := g.Client.BulkUpdateOrderStatus(ctx, &orderpb.BulkUpdateOrderStatusRequest{
		OrderIds: ids,
		Status:   orderpb.OrderStatus_Expired,
	})

	if err != nil {
		log.Logger.Errorf("g.Client.BulkUpdateOrderStatus, %v", err.Error())

		return nil, err
	}

	return ok, err
}
