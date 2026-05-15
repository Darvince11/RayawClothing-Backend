package interfaces

import "github.com/google/uuid"

type PaymentProcessor interface {
	InitializePayment(email string, amount float64, orderId uuid.UUID) (string, error)
}
