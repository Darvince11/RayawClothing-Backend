package services

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"rayaw-api/internal/interfaces"
	"rayaw-api/internal/models"
	"rayaw-api/internal/repositories"
	"strings"

	"github.com/google/uuid"
)

type OrderServiceInterface interface {
	AddOrder(orderRequest *models.CreateOrderRequest) (*models.AddOrderResponse, error)
	GetOrdersByUserId(userId int) (*[]models.GetOrderByUserIdResponse, error)
	GetOrderById(orderId uuid.UUID) (*models.OrderWithItems, error)
	UpdateOrderStatus(orderId uuid.UUID, status models.OrderStatus) error
}

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

	orderIdStr, err := generateOrderId()
	if err != nil {
		return nil, err
	}

	order := &models.Order{
		OrderId:     orderIdStr,
		UserId:      orderRequest.UserId,
		TotalAmount: totalAmount,
		OrderStatus: models.OrderStatusPending,
	}

	tx, err := os.db.BeginTx(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	Id, err := os.or.AddOrder(order, tx)
	if err != nil {
		return nil, err
	}

	// add order items
	orderItems := []models.OrderItem{}
	for _, product := range orderRequest.Products {
		orderItem := models.OrderItem{
			OrderId:   Id,
			ProductId: product.ProdctID,
			Quantity:  product.Quantity,
		}
		orderItems = append(orderItems, orderItem)
	}
	err = os.or.AddOrderItems(&orderItems, tx)
	if err != nil {
		return nil, err
	}

	//add payment history
	reference, err := GenerateReference()
	if err != nil {
		return nil, err
	}

	paymentId, err := os.paymentProcessor.GeneratePaymentId()
	if err != nil {
		return nil, err
	}

	paymentHistory := &models.PaymentHistory{
		PaymentId:     paymentId,
		OrderId:       Id,
		UserId:        orderRequest.UserId,
		Reference:     reference,
		Amount:        totalAmount,
		PaymentStatus: models.PaymentStatusPending,
	}

	_, err = os.paymentProcessor.AddPaymentHistory(paymentHistory, tx)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	// initialize payment
	authUrl, err := os.paymentProcessor.InitializePayment(orderRequest.Email, totalAmount, Id, reference)
	if err != nil {
		return nil, err
	}

	return &models.AddOrderResponse{
		OrderId:          Id,
		AuthorizationUrl: authUrl,
	}, nil
}

func (os *OrderService) GetOrdersByUserId(userId int) (*[]models.GetOrderByUserIdResponse, error) {
	//get orders
	orders, err := os.or.GetOrdersByUserId(userId)
	if err != nil {
		return nil, err
	}

	orderIds := make([]uuid.UUID, len(*orders))
	for index, order := range *orders {
		orderIds[index] = order.Id
	}

	// get order items
	orderItems, err := os.or.GetOrderItemsByOrderId(orderIds)
	if err != nil {
		return nil, err
	}

	//create a list of product ids
	var productIds []int
	for key := range *orderItems {
		for _, item := range (*orderItems)[key] {
			productIds = append(productIds, item.ProductId)
		}
	}

	// use porduct id from order items to get product names
	products, err := os.pr.GetProductsById(productIds)
	if err != nil {
		return nil, err
	}

	orderIdToProductName := make(map[uuid.UUID][]string)
	for key, items := range *orderItems {
		for _, product := range *products {
			for _, orderItem := range items {
				if orderItem.ProductId == product.Id {
					orderIdToProductName[key] = append(orderIdToProductName[key], product.Product_name)
					break
				}
			}
		}
	}

	var orderResponses []models.GetOrderByUserIdResponse

	for _, order := range *orders {
		orderResponse := models.GetOrderByUserIdResponse{
			Id:              order.Id,
			OrderId:         order.OrderId,
			UserId:          order.UserId,
			TotalAmount:     order.TotalAmount,
			OrderStatus:     order.OrderStatus,
			OrderDate:       order.OrderDate,
			OrderItemsNames: orderIdToProductName[order.Id],
		}
		orderResponses = append(orderResponses, orderResponse)
	}
	return &orderResponses, nil
}

func (os *OrderService) GetOrderById(orderId uuid.UUID) (*models.OrderWithItems, error) {
	order, err := os.or.GetOrderById(orderId)
	if err != nil {
		return nil, err
	}

	orderItems, err := os.or.GetOrderItemsByOrderId([]uuid.UUID{orderId})
	if err != nil {
		return nil, err
	}

	productsId := []int{}
	for _, item := range *orderItems {
		for _, i := range item {
			productsId = append(productsId, i.ProductId)
		}
	}

	products, err := os.pr.GetProductsById(productsId)
	if err != nil {
		return nil, err
	}

	getOrderItems := []models.GetOrderItemsResponse{}

	for _, items := range *orderItems {
		for _, item := range items {
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

	}

	orderWithItems := &models.OrderWithItems{
		OrderId:               order.Id,
		Status:                order.OrderStatus,
		CreatedAt:             order.OrderDate,
		GetOrderItemsResponse: getOrderItems,
	}
	return orderWithItems, nil
}

func (os *OrderService) UpdateOrderStatus(orderId uuid.UUID, status models.OrderStatus) error {
	return os.or.UpdateOrderStatus(orderId, status)
}

func GenerateReference() (string, error) {
	bytes := make([]byte, 8)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	reference := base64.RawURLEncoding.EncodeToString(bytes)
	return reference, nil
}

func generateOrderId() (string, error) {
	idList := []string{"OD"}
	var randBytes [4]byte
	_, err := rand.Read(randBytes[:])

	if err != nil {
		return "", err
	}

	randStr := hex.EncodeToString(randBytes[:])
	randStr = strings.ToUpper(randStr)
	idList = append(idList, randStr[0:4], randStr[4:])

	return strings.Join(idList, "-"), nil
}
