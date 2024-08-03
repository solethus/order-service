package server

import (
	"context"

	"github.com/solethus/order-service/internal/service"
	pb "github.com/solethus/shared-proto/proto/order"
)

type OrderServer struct {
	pb.UnimplementedOrderServiceServer
	service service.OrderService
}

func NewOrderServer(svc service.OrderService) *OrderServer {
	return &OrderServer{
		service: svc,
	}
}

func (s *OrderServer) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.Order, error) {
	// Implement: convert request to model, call service, convert result to proto
	panic("implement me")
}
