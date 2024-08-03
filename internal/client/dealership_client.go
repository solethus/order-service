package client

import (
	"context"
	pb "github.com/solethus/shared-proto/proto/dealership"
	"google.golang.org/grpc"
)

type DealershipClient interface {
	GetDealershipInfo(ctx context.Context, dealershipID string) (*pb.DealershipInfo, error)
}

type dealershipClient struct {
	client pb.DealershipServiceClient
}

func NewDealershipClient(conn *grpc.ClientConn) DealershipClient {
	return &dealershipClient{
		client: pb.NewDealershipServiceClient(conn),
	}
}

func (c *dealershipClient) GetDealershipInfo(ctx context.Context, dealershipID string) (*pb.DealershipInfo, error) {
	return c.client.GetDealerShipInfo(ctx, &pb.GetDealershipInfoRequest{DealershipId: dealershipID})
}
