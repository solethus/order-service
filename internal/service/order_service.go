package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/solethus/order-service/internal/client"
	"github.com/solethus/order-service/internal/model"
	"github.com/solethus/order-service/internal/repository"
)

type OrderService interface {
	GetOrder(ctx context.Context, id uuid.UUID) (*model.Order, error)
	UpdateOrderStatus(ctx context.Context, id uuid.UUID, status model.OrderStatus)
}

type orderService struct {
	repo         repository.OrderRepository
	inventoryCl  client.InventoryClient
	dealershipCl client.DealershipClient
}

func NewOrderService(repo repository.OrderRepository, inventoryClient client.InventoryClient, dealershipClient client.DealershipClient) OrderService {
	return &orderService{
		repo:         repo,
		inventoryCl:  inventoryClient,
		dealershipCl: dealershipClient,
	}
}

func (s *orderService) CreateOrder(ctx context.Context, order *model.Order) error {
	// Check if car is available
	car, err := s.inventoryCl.GetCar(ctx, order.CarID.String())
	if err != nil {
		return fmt.Errorf("failed to get car info: %w", err)
	}

	if car.Stock <= 0 {
		return errors.New("car not available")
	}

	// Check dealership
	_, err = s.dealershipCl.GetDealershipInfo(ctx, order.DealershipID.String())
	if err != nil {
		return fmt.Errorf("failed to get dealer ship info: %w", err)
	}

	// Create order
	err = s.repo.CreateOrder(ctx, order)
	if err != nil {
		return fmt.Errorf("failed to create order: %w", err)
	}

	// Update inventory
	err = s.inventoryCl.UpdateCarStock(ctx, order.CarID.String(), -1)
	if err != nil {
		return fmt.Errorf("failed to update inventory: %w", err)
	}

	return nil
}

func (s *orderService) GetOrder(ctx context.Context, id uuid.UUID) (*model.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (s *orderService) UpdateOrderStatus(ctx context.Context, id uuid.UUID, status model.OrderStatus) {
	//TODO implement me
	panic("implement me")
}
