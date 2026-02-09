package repositories

import (
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
		Id:          uuid.New(),
		UserId:      1,
		TotalAmount: 200,
		OrderStatus: models.OrderStatusPending,
	}

	_, err := repo.AddOrder(&order)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	orders, err := repo.GetOrdersByUserId(order.UserId)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	t.Logf("Orders: %v", orders)
}
