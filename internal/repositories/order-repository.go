package repositories

import (
	"database/sql"
	"rayaw-api/internal/models"

	"github.com/google/uuid"
)

type OrderRepository interface {
	GetOrdersByUserId(userId int) (*[]models.Order, error)
	AddOrder(order *models.Order) (int, error)
	AddOrderItems(orderItems *[]models.OrderItem) error
	GetOrderById(orderId uuid.UUID) (*models.Order, error)
	GetOrderItemsByOrderId(orderId uuid.UUID) (*[]models.OrderItem, error)
}

type ImplOrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *ImplOrderRepository {
	return &ImplOrderRepository{db: db}
}

func (r *ImplOrderRepository) GetOrdersByUserId(userId int) (*[]models.Order, error) {
	query := `SELECT id, user_id, total_amount, order_status, order_date
	 FROM orders 
	 WHERE user_id = $1`

	var orders []models.Order

	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.Id, &order.UserId, &order.TotalAmount, &order.OrderStatus, &order.OrderDate)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return &orders, nil
}

func (r *ImplOrderRepository) AddOrder(order *models.Order) (uuid.UUID, error) {
	query := `INSERT INTO orders (user_id, total_amount, order_status)
	 VALUES ($1, $2, $3) RETURNING id`
	var id uuid.UUID
	err := r.db.QueryRow(query, order.UserId, order.TotalAmount, order.OrderStatus).Scan(&id)
	return id, err
}

func (r *ImplOrderRepository) AddOrderItems(orderItems *[]models.OrderItem) error {
	query := `INSERT INTO order_items (order_id, product_id, quantity)
	VALUES ($1, $2, $3)`

	for _, item := range *orderItems {
		_, err := r.db.Exec(query, item.OrderId, item.ProductId, item.Quantity)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *ImplOrderRepository) GetOrderById(orderId uuid.UUID) (*models.Order, error) {
	query := `SELECT id, user_id, total_amount, order_status, order_date
	 FROM orders 
	 WHERE id = $1`
	var order models.Order
	err := r.db.QueryRow(query, orderId).Scan(&order.Id, &order.UserId, &order.TotalAmount, &order.OrderStatus, &order.OrderDate)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *ImplOrderRepository) GetOrderItemsByOrderId(orderId uuid.UUID) (*[]models.OrderItem, error) {
	query := `SELECT id, order_id, product_id, quantity
	 FROM order_items 
	 WHERE order_id = $1`

	var orderItems []models.OrderItem

	rows, err := r.db.Query(query, orderId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.OrderItem
		err := rows.Scan(&item.Id, &item.OrderId, &item.ProductId, &item.Quantity)
		if err != nil {
			return nil, err
		}
		orderItems = append(orderItems, item)
	}
	return &orderItems, nil
}
