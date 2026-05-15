package services

import (
	"database/sql"
	"rayaw-api/internal/interfaces"
	"rayaw-api/internal/models"
	"rayaw-api/internal/repositories"

	"github.com/google/uuid"
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

func (os *OrderService) AddOrder(orderRequest *models.CreateOrderRequest) (*models.AddOrderResponse, error) {
	//Get the total amount from the product
	totalAmount := 0.0
	productsId := make([]int, len(orderRequest.Products))
	for _, product := range orderRequest.Products {
		productsId = append(productsId, product.ProdctID)
	}

	products, err := os.pr.GetProductsById(productsId)
	if err != nil {
		return nil, err
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
		return nil, err
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
		return nil, err
	}

	// initialize payment
	authUrl, err := os.paymentProcessor.InitializePayment(orderRequest.Email, totalAmount, orderId)
	if err != nil {
		return nil, err
	}

	return &models.AddOrderResponse{
		OrderId:          orderId,
		AuthorizationUrl: authUrl,
	}, nil
}

func (os *OrderService) GetOrdersByUserId(userId int) (*[]models.Order, error) {
	return os.or.GetOrdersByUserId(userId)
}

func (os *OrderService) GetOrderById(orderId uuid.UUID) (*models.OrderWithItems, error) {
	order, err := os.or.GetOrderById(orderId)
	if err != nil {
		return nil, err
	}

	orderItems, err := os.or.GetOrderItemsByOrderId(orderId)
	if err != nil {
		return nil, err
	}

	productsId := make([]int, len(*orderItems))
	for index, item := range *orderItems {
		productsId[index] = item.ProductId
	}

	products, err := os.pr.GetProductsById(productsId)
	if err != nil {
		return nil, err
	}

	getOrderItems := []models.GetOrderItemsResponse{}

	for _, item := range *orderItems {
		for _, product := range *products {
			if item.ProductId == product.Id {
				getOrderItemsResponse := models.GetOrderItemsResponse{
					ImageUrl:    product.Image_url,
					ProductId:   product.Id,
					ProductName: product.Product_name,
					Price:       product.Price,
					Quantity:    item.Quantity,
				}
				getOrderItems = append(getOrderItems, getOrderItemsResponse)
			}
		}
	}

	orderWithItems := &models.OrderWithItems{
		OrderId:               order.Id,
		Status:                order.OrderStatus,
		CreatedAt:             order.OrderDate,
		GetOrderItemsResponse: getOrderItems,
	}
	return orderWithItems, nil
}
