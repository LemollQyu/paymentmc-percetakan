package usecase

import (
	"context"
	"paymentmc/models"
)

func (uc *PaymentUsecase) GetWaitingPayment(ctx context.Context, code string) (*models.ListWaitingPayment, error) {
	data, err := uc.PaymentService.GetWaitingPayment(ctx, code)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (uc *PaymentUsecase) GetListWaitingPayment(ctx context.Context) (*[]models.ListWaitingPayment, error) {
	data, err := uc.PaymentService.GetListWaitingPayment(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetMyPayment mengembalikan data checkout user yang login: payment status Pending, order Waiting_payment.
// Format data sama dengan list-waiting-payment, hanya difilter per user.
func (uc *PaymentUsecase) GetMyPayment(ctx context.Context, userID int64) (*[]models.ListWaitingPayment, error) {
	data, err := uc.PaymentService.GetListWaitingPaymentByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return data, nil
}
