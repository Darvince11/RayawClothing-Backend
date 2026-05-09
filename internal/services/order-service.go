package services

import (
	"rayaw-api/internal/models"
	"rayaw-api/internal/repositories"

	"github.com/google/uuid"
)

type OrderService struct {
	or *repositories.ImplOrderRepository
}

func NewOrderService(or *repositories.ImplOrderRepository) *OrderService {
	return &OrderService{or: or}
}

func (os *OrderService) AddOrder(order *models.Order) (uuid.UUID, error) {
	return os.or.AddOrder(order)
}

func (os *OrderService) GetOrdersByUserId(userId int) (*[]models.Order, error) {
	return os.or.GetOrdersByUserId(userId)
}

func (os *OrderService) AddOrderItems(orderItems *[]models.OrderItem) error {
	return os.or.AddOrderItems(orderItems)
}
