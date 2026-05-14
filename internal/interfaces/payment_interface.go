package interfaces

type PaymentProcessor interface {
	InitializePayment(email string, amount float64) (string, error)
}
