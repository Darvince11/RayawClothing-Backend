package repositories

import (
	"context"
	"rayaw-api/internal/models"
	"rayaw-api/internal/tests"
	"testing"

	"github.com/google/uuid"
)

func TestPaymentRepository(t *testing.T) {
	db := tests.SetupTestDB(t)
	if db == nil {
		t.Fatal("Failed to set up test database")
	}

	repo := NewImplPaymentRepository(db)

	paymentHistory := &models.PaymentHistory{
		Id:            1,
		OrderId:       uuid.MustParse("7427199a-376f-4aa9-adac-ee69c8c4677b"),
		UserId:        9,
		Reference:     "RAYAW-8950A613-1758802096",
		Currency:      "GHS",
		PaymentMethod: models.PaymentMethod("mobile_money"),
		Amount:        200.00,
		PaymentStatus: models.PaymentStatus("success"),
	}

	//test for add payment history
	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		t.Fatal("Failed to begin transaction")
	}
	_, err = repo.AddPaymentHistory(paymentHistory, tx)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	tx.Commit()

	//test for get payment history
	allHistory, err := repo.GetAllPaymentHistoryByUserId(9)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	t.Logf("Payment History: %v", allHistory)

	//test for get payment history by reference
	history, err := repo.GetPaymentHistoryByReference("RAYAW-8950A613-1758802096")
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	t.Logf("Payment History by Reference: %v", history)

	//test for update payment history
	status := models.PaymentStatusCompleted
	updateReq := &models.UpdatePaymentHistoryRequest{
		PaymentStatus: &status,
	}

	orderId, err := repo.UpdatePaymentHistory(updateReq, "RAYAW-8950A613-1758802096")
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	t.Logf("Updated Payment History Order ID: %v", orderId)
}
