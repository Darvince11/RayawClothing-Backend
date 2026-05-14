package services

import (
	"database/sql"
	"rayaw-api/internal/interfaces"
	"rayaw-api/internal/models"
	"rayaw-api/internal/repositories"
)

type OrderService struct {
	or               *repositories.ImplOrderRepository
	pr               repositories.ProductsRepository
	db               *sql.DB
	paymentProcessor interfaces.PaymentProcessor
}

func NewOrderService(or *repositories.ImplOrderRepository, pr repositories.ProductsRepository, db *sql.DB, paymentProcessor interfaces.PaymentProcessor) *OrderService {
	return &OrderService{or: or, pr: pr, db: db, paymentProcessor: paymentProcessor}
}

func (os *OrderService) AddOrder(orderRequest *models.CreateOrderRequest) (string, error) {
	//Get the total amount from the product
	totalAmount := 0.0
	productsId := make([]int, len(orderRequest.Products))
	for _, product := range orderRequest.Products {
		productsId = append(productsId, product.ProdctID)
	}

	products, err := os.pr.GetProductsById(productsId)
	if err != nil {
		return "", err
	}

	for index, product := range *products {
		totalAmount += (product.Price * float64(orderRequest.Products[index].Quantity))
	}

	order := &models.Order{
		UserId:      orderRequest.UserId,
		TotalAmount: totalAmount,
		OrderStatus: models.OrderStatusPending,
	}

	orderId, err := os.or.AddOrder(order)
	if err != nil {
		return "", err
	}

	// add order items
	orderItems := []models.OrderItem{}
	for _, product := range orderRequest.Products {
		orderItem := models.OrderItem{
			OrderId:   orderId,
			ProductId: product.ProdctID,
			Quantity:  product.Quantity,
		}
		orderItems = append(orderItems, orderItem)
	}
	err = os.or.AddOrderItems(&orderItems)
	if err != nil {
		return "", err
	}

	// initialize payment
	authUrl, err := os.paymentProcessor.InitializePayment(orderRequest.Email, totalAmount)
	if err != nil {
		return "", err
	}

	return authUrl, nil
}

func (os *OrderService) GetOrdersByUserId(userId int) (*[]models.Order, error) {
	return os.or.GetOrdersByUserId(userId)
}
