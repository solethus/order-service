package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/solethus/order-service/internal/model"
	"github.com/solethus/order-service/internal/repository"
)

type OrderService interface {
	GetOrder(ctx context.Context, id uuid.UUID) (*model.Order, error)
	UpdateOrderStatus(ctx context.Context, id uuid.UUID, status model.OrderStatus)
}

type orderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{
		repo: repo,
	}
}

func (s *orderService) GetOrder(ctx context.Context, id uuid.UUID) (*model.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (s *orderService) UpdateOrderStatus(ctx context.Context, id uuid.UUID, status model.OrderStatus) {
	//TODO implement me
	panic("implement me")
}
