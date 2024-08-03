package repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/solethus/order-service/internal/model"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *model.Order) error
	GetOrder(ctx context.Context, id uuid.UUID) (*model.Order, error)
	UpdateOrder(ctx context.Context, order *model.Order) error
	ListOrders(ctx context.Context, filter model.OrderFilter) ([]*model.Order, int, error)
	DeleteOrder(ctx context.Context, id uuid.UUID) error
}

type orderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (o orderRepository) CreateOrder(ctx context.Context, order *model.Order) error {
	//TODO implement me
	panic("implement me")
}

func (o orderRepository) GetOrder(ctx context.Context, id uuid.UUID) (*model.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o orderRepository) UpdateOrder(ctx context.Context, order *model.Order) error {
	//TODO implement me
	panic("implement me")
}

func (o orderRepository) ListOrders(ctx context.Context, filter model.OrderFilter) ([]*model.Order, int, error) {
	//TODO implement me
	panic("implement me")
}

func (o orderRepository) DeleteOrder(ctx context.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
