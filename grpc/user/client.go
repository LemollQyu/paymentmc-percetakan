package user

import (
	"context"
	"log"
	"paymentmc/proto/userpb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserClient struct {
	Client userpb.UserServiceClient
}

func NewUserClient() *UserClient {
	conn, err := grpc.Dial(
		"localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed init grpc client: %v", err)
	}

	client := userpb.NewUserServiceClient(conn)
	return &UserClient{
		Client: client,
	}
}

func (uc *UserClient) GetUserInfoByUserID(ctx context.Context, userID int64) (*userpb.GetUserInfoResult, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)

	defer cancel()

	userInfo, err := uc.Client.GetUserInfoByUserID(ctx, &userpb.GetUserInfoRequest{
		UserId: userID,
	})

	if err != nil {
		return nil, err
	}

	return userInfo, nil
}

func (uc *UserClient) GetUsersByIDs(
	ctx context.Context,
	userIDs []int64,
) (*userpb.GetUsersByIDsResult, error) {

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	res, err := uc.Client.GetUsersByIDs(ctx, &userpb.GetUsersByIDsRequest{
		UserIds: userIDs,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
