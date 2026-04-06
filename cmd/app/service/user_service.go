package service

import (
	"context"
	"paymentmc/models"
)

func (s *PaymentService) GetUser(
	ctx context.Context,
	userID int64,
) (*models.User, error) {

	res, err := s.UserClient.GetUserInfoByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:        res.Id,
		Name:      res.Name,
		Email:     res.Email,
		Phone:     &res.Phone,
		AvatarURL: &res.AvatarUrl,
	}

	return user, nil
}

func (s *PaymentService) GetUserIds(
	ctx context.Context,
	userIDs []int64,
) (map[int64]*models.User, error) {

	res, err := s.UserClient.GetUsersByIDs(ctx, userIDs)
	if err != nil {
		return nil, err
	}

	userMap := make(map[int64]*models.User)

	for _, u := range res.Users {
		userMap[u.Id] = &models.User{
			ID:        u.Id,
			Name:      u.Name,
			Email:     u.Email,
			Phone:     &u.Phone,
			AvatarURL: &u.AvatarUrl,
		}
	}

	return userMap, nil
}
