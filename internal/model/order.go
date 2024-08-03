package model

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID                    uuid.UUID     `json:"id"`
	CustomerID            uuid.UUID     `json:"customer_id"`
	DealershipID          uuid.UUID     `json:"dealership_id"`
	CarID                 uuid.UUID     `json:"car_id"`
	TotalPrice            float64       `json:"total_price"`
	Status                OrderStatus   `json:"status"`
	OrderDate             time.Time     `json:"order_date"`
	EstimatedDeliveryDate time.Time     `json:"estimated_delivery_date"`
	Items                 []OrderItem   `json:"items"`
	PaymentStatus         PaymentStatus `json:"payment_status"`
	ShippingAddress       Address       `json:"shipping_address"`
	Notes                 string        `json:"notes"`
}

type OrderItem struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Price    float64   `json:"price"`
	Quantity int       `json:"quantity"`
}

type OrderStatus string

const (
	OrderStatusPending          OrderStatus = "PENDING"
	OrderStatusConfirmed        OrderStatus = "CONFIRMED"
	OrderStatusInProduction     OrderStatus = "IN_PRODUCTION"
	OrderStatusReadyForDelivery OrderStatus = "READY_FOR_DELIVERY"
	OrderStatusDelivered        OrderStatus = "DELIVERED"
	OrderStatusCancelled        OrderStatus = "CANCELLED"
)

type PaymentStatus string

const (
	PaymentStatusPending  PaymentStatus = "PENDING"
	PaymentStatusPaid     PaymentStatus = "PAID"
	PaymentStatusRefunded PaymentStatus = "REFUNDED"
)

type Address struct {
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
}

type OrderFilter struct {
	CustomerID   uuid.UUID
	DealershipID uuid.UUID
	Status       OrderStatus
	StartDate    time.Time
	EndDate      time.Time
}
