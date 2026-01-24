package order

import (
	"log"
	"paymentmc/proto/orderpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OrderClient struct {
	Client   orderpb.OrderServiceClient
	HostGRPC string
}

func NewOrderClient(hostGRPC string) *OrderClient {
	conn, err := grpc.Dial(
		hostGRPC, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Gagal inisialisasi gRPC Client: %v", err)
	}

	client := orderpb.NewOrderServiceClient(conn)
	return &OrderClient{
		Client:   client,
		HostGRPC: hostGRPC,
	}
}
