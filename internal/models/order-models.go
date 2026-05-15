package models

import "github.com/google/uuid"

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusPaid      OrderStatus = "paid"
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
	Products []OrderProduct `json:"products"`
	UserId   int            `json:"user_id"`
	Email    string         `json:"email"`
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

type GetOrderItemsResponse struct {
	ImageUrl    string  `json:"image_url"`
	ProductId   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Color       string  `json:"color"`
	Size        string  `json:"size"`
}

type OrderWithItems struct {
	OrderId               uuid.UUID               `json:"order_id"`
	Status                OrderStatus             `json:"status"`
	CreatedAt             string                  `json:"created_at"`
	GetOrderItemsResponse []GetOrderItemsResponse `json:"order_items"`
}

type AddOrderResponse struct {
	OrderId          uuid.UUID `json:"order_id"`
	AuthorizationUrl string    `json:"authorization_url"`
}
