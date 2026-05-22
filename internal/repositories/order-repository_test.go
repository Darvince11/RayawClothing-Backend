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
		Id:          uuid.New(),
		UserId:      1,
		TotalAmount: 200,
		OrderStatus: models.OrderStatusPending,
	}

	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		t.Fatalf("Failed to begin transaction: %v", err)
	}

	_, err = repo.AddOrder(&order, tx)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	tx.Commit()

	orders, err := repo.GetOrdersByUserId(order.UserId)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	t.Logf("Orders: %v", orders)
}
