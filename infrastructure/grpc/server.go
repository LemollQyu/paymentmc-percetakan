package grpc

import (
	"context"

	"paymentmc/cmd/app/usecase"
	"paymentmc/infrastructure/log"
	"paymentmc/proto/paymentpb"
	"paymentmc/utils"
)

type GRPCServer struct {
	paymentpb.UnimplementedPaymentServiceServer
	PaymentUsecase usecase.PaymentUsecase
}

func (g *GRPCServer) DeletePaymentByOrderID(
	ctx context.Context,
	req *paymentpb.RequestDeletePaymentByOrderID,
) (*paymentpb.ResponseDeletePaymentByOrderID, error) {

	// cek payment by order id
	dataPayment, err := g.PaymentUsecase.GetPaymentByOrderID(ctx, req.OrderId)
	if err != nil {
		log.Logger.Error("g.PaymentUsecase.GetPaymentByOrderID")
		return nil, err
	}

	// jika tidak ada payment -> tidak delete apa apa (idempotent)
	if dataPayment == nil {
		return &paymentpb.ResponseDeletePaymentByOrderID{
			Success: true,
			Message: "payment tidak ditemukan, tidak ada yang dihapus",
		}, nil
	}

	// jika masih pending tidak boleh dihapus
	if dataPayment.Status == utils.StatusPaymentPending {
		return nil, utils.ErrNotDeletePayment
	}

	// delete
	err = g.PaymentUsecase.DeletePaymentByOrderID(ctx, req.OrderId)
	if err != nil {
		log.Logger.Error("g.PaymentUsecase.DeletePaymentByOrderID")
		return nil, err
	}

	return &paymentpb.ResponseDeletePaymentByOrderID{
		Success: true,
		Message: "berhasil dihapus",
	}, nil
}
