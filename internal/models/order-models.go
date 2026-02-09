package models

import "github.com/google/uuid"

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusShipping  OrderStatus = "shipping"
	OrderStatusDelivery  OrderStatus = "delivery"
	OrderStatusCompleted OrderStatus = "completed"
	OrderStatusCancelled OrderStatus = "cancelled"
)

type OrderProduct struct {
	ProdctID int `json:"product_id"`
	Quantity int `json:"quantity"`
}

type CreateOrderRequest struct {
	Products      []OrderProduct `json:"products"`
	UserId        int            `json:"user_id"`
	PaymentMethod PaymentMethod  `json:"payment_method"`
	TotalAmount   float64        `json:"total_amount"`
	OrderStatus   OrderStatus    `json:"order_status"`
}

type Order struct {
	Id          uuid.UUID   `json:"id"`
	UserId      int         `json:"user_id"`
	TotalAmount float64     `json:"total_amount"`
	OrderStatus OrderStatus `json:"order_status"`
	OrderDate   string      `json:"order_date"`
}

type OrderItem struct {
	Id        int       `json:"id"`
	OrderId   uuid.UUID `json:"order_id"`
	ProductId int       `json:"product_id"`
	Quantity  int       `json:"quantity"`
}
