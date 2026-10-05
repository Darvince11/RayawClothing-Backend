package repositories

import (
	"context"
	"rayaw-api/internal/models"
	"rayaw-api/internal/tests"
	"testing"

	"github.com/google/uuid"
)

func TestOrderRepository(t *testing.T) {
	db := tests.SetupTestDB(t)
	if db == nil {
		t.Fatal("Failed to set up test database")
	}

	repo := NewOrderRepository(db)

	order := models.Order{
		UserId:      1,
		TotalAmount: 200,
		OrderStatus: models.OrderStatusPending,
	}

	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		t.Fatalf("Failed to begin transaction: %v", err)
	}

	// test for add order
	orderId, err := repo.AddOrder(&order, tx)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	//test for add order itmes
	orderItems := []models.OrderItem{
		{
			OrderId:   orderId,
			ProductId: 1,
			Quantity:  2,
		},
	}

	err = repo.AddOrderItems(&orderItems, tx)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	tx.Commit()

	//test getOrdersByUserId
	orders, err := repo.GetOrdersByUserId(order.UserId)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	t.Logf("Orders: %v", orders)

	//test update order status
	err = repo.UpdateOrderStatus((*orders)[0].Id, models.OrderStatusPaid)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	//test get odrer by id
	orderS, err := repo.GetOrderById((*orders)[0].Id)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	t.Logf("Order: %v", orderS)

	//test for get order items by order id
	orderIds := (*orders)[0].Id
	items, err := repo.GetOrderItemsByOrderId([]uuid.UUID{orderIds})
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	t.Logf("Order Items: %v", items)

}
