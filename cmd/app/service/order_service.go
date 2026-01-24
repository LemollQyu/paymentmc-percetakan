package service

import (
	"context"

	"paymentmc/infrastructure/log"
	"paymentmc/models"
)

func (s *PaymentService) GetOrderByID(
	ctx context.Context,
	orderID int64,
) (*models.Order, error) {

	res, err := s.OrderClient.GetOrderDetailByID(ctx, orderID)
	if err != nil {
		log.Logger.Error("OrderService:GetOrderDetailByID", err)
		return nil, err
	}

	user := &models.User{
		ID:        res.User.Id,
		Name:      res.User.Name,
		Email:     res.User.Email,
		Phone:     &res.User.Phone,
		AvatarUrl: &res.User.AvatarUrl,
	}

	orderCode := &models.OrderCode{
		ID:      res.OrderCode.Id,
		OrderID: res.OrderCode.OrderId,
		Code:    res.OrderCode.Code,
	}

	var orderFile *models.OrderFile
	if res.OrderFile != nil {
		orderFile = &models.OrderFile{
			ID:       res.OrderFile.Id,
			OrderID:  res.OrderFile.OrderId,
			FileUrl:  res.OrderFile.FileUrl,
			Type:     res.OrderFile.Type,
			FileType: res.OrderFile.FileType,
		}
	}

	specs := make([]*models.OrderSpesification, 0)
	for _, spec := range res.OrderSpecifications {
		specs = append(specs, &models.OrderSpesification{
			ID:                        spec.Id,
			OrderID:                   spec.OrderId,
			SpesificationID:           spec.SpecificationId,
			SpesificationNameSnapshot: spec.SpecificationNameSnapshot,
			ValueSnapshot:             spec.ValueSnapshot,
			AdditionalPriceSnapshot:   spec.AdditionalPriceSnapshot,
		})
	}

	return &models.Order{
		ID:                  res.Id,
		UserID:              res.UserId,
		ServiceID:           res.ServiceId,
		ServiceNameSnapshot: res.ServiceNameSnapshot,
		BasePriceSnapshot:   res.BasePriceSnapshot,
		TotalPriceSnapshot:  res.TotalPriceSnapshot,
		UserNote:            res.UserNote,
		Status:              res.Status.String(),
		Quantity:            res.Quantity,
		User:                user,
		OrderCode:           *orderCode,
		OrderFile:           orderFile,
		OrderSpesifications: specs,
	}, nil
}

func (s *PaymentService) GetOrderByCode(
	ctx context.Context,
	code string,
) (*models.Order, error) {

	res, err := s.OrderClient.GetOrderDetailByCode(ctx, code)
	if err != nil {
		log.Logger.Error("OrderService:GetOrderDetailByCode", err)
		return nil, err
	}

	user := &models.User{
		ID:        res.User.Id,
		Name:      res.User.Name,
		Email:     res.User.Email,
		Phone:     &res.User.Phone,
		AvatarUrl: &res.User.AvatarUrl,
	}

	orderCode := &models.OrderCode{
		ID:      res.OrderCode.Id,
		OrderID: res.OrderCode.OrderId,
		Code:    res.OrderCode.Code,
	}

	var orderFile *models.OrderFile
	if res.OrderFile != nil {
		orderFile = &models.OrderFile{
			ID:       res.OrderFile.Id,
			OrderID:  res.OrderFile.OrderId,
			FileUrl:  res.OrderFile.FileUrl,
			Type:     res.OrderFile.Type,
			FileType: res.OrderFile.FileType,
		}
	}

	specs := make([]*models.OrderSpesification, 0)
	for _, spec := range res.OrderSpecifications {
		specs = append(specs, &models.OrderSpesification{
			ID:                        spec.Id,
			OrderID:                   spec.OrderId,
			SpesificationID:           spec.SpecificationId,
			SpesificationNameSnapshot: spec.SpecificationNameSnapshot,
			ValueSnapshot:             spec.ValueSnapshot,
			AdditionalPriceSnapshot:   spec.AdditionalPriceSnapshot,
		})
	}

	return &models.Order{
		ID:                  res.Id,
		UserID:              res.UserId,
		ServiceID:           res.ServiceId,
		ServiceNameSnapshot: res.ServiceNameSnapshot,
		BasePriceSnapshot:   res.BasePriceSnapshot,
		TotalPriceSnapshot:  res.TotalPriceSnapshot,
		UserNote:            res.UserNote,
		Status:              res.Status.String(),
		Quantity:            res.Quantity,
		User:                user,
		OrderCode:           *orderCode,
		OrderFile:           orderFile,
		OrderSpesifications: specs,
	}, nil
}

func (s *PaymentService) SetStatusOrder(ctx context.Context, code string, status string) (bool, error) {
	res, err := s.OrderClient.UpdateOrderStatus(ctx, code, status)
	if err != nil {
		log.Logger.Error("s.OrderClient.UpdateOrderStatus")
		return false, err
	}

	if !res.Success {
		return false, err
	}

	return true, nil
}

func (s *PaymentService) UpdateExpiredOrder(ctx context.Context, ids []int64) (bool, int64, error) {
	res, err := s.OrderClient.ExpiredOrderStatus(ctx, ids)
	if err != nil {
		log.Logger.Error("s.OrderClient.ExpiredOrderStatus")
		return false, 0, err
	}

	return res.Success, res.TotalUpdated, nil
}
