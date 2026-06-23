package interfaces

import (
	"database/sql"
	"rayaw-api/internal/models"

	"github.com/google/uuid"
)

type PaymentProcessor interface {
	InitializePayment(email string, amount float64, orderId uuid.UUID, reference string) (string, error)
	AddPaymentHistory(paymentHistory *models.PaymentHistory, tx *sql.Tx) (int, error)
}
