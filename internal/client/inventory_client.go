package client

import (
	"context"
	pb "github.com/solethus/shared-proto/proto/inventory"
	"google.golang.org/grpc"
)

type InventoryClient interface {
	GetCar(ctx context.Context, carID string) (*pb.Car, error)
	UpdateCarStock(ctx context.Context, carID string, change int32) error
}

type inventoryClient struct {
	client pb.InventoryServiceClient
}

func NewInventoryClient(conn *grpc.ClientConn) InventoryClient {
	return &inventoryClient{
		client: pb.NewInventoryServiceClient(conn),
	}
}

func (c *inventoryClient) GetCar(ctx context.Context, carID string) (*pb.Car, error) {
	return c.client.GetCar(ctx, &pb.GetCarRequest{Id: carID})
}

func (c *inventoryClient) UpdateCarStock(ctx context.Context, carID string, change int32) error {
	_, err := c.client.UpdateCarStock(ctx, &pb.UpdateCarStockRequest{
		Id:          carID,
		StockChange: change,
	})
	return err
}
