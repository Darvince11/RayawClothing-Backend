package repositories

import (
	"database/sql"
	"rayaw-api/internal/models"
)

type PaymentRepository interface {
	AddPaymentHistory(paymentHistory *models.PaymentHistory) (int, error)
	GetPaymentHistoryByReference(reference string) (*models.PaymentHistory, error)
	UpdatePaymentHistoryStatus(status *string) error
}

type ImplPaymentRepository struct {
	db *sql.DB
}

func NewImplPaymentRepository(db *sql.DB) *ImplPaymentRepository {
	return &ImplPaymentRepository{db: db}
}

func (pr *ImplPaymentRepository) AddPaymentHistory(paymentHistory *models.PaymentHistory) (int, error) {
	query := `INSERT INTO payments_history (order_id, payment_method, amount, payment_status) VALUES ($1, $2, $3, $4) RETURNING payment_id`
	var paymentId int
	err := pr.db.QueryRow(query, paymentHistory.OrderId, paymentHistory.PaymentMethod, paymentHistory.Amount, paymentHistory.PaymentStatus).Scan(&paymentId)
	return paymentId, err
}

func (pr *ImplPaymentRepository) GetPaymentHistoryByReference(reference string) (*models.PaymentHistory, error) {
	query := `SELECT * FROM payments_history WHERE reference = $1`

	var paymentHistory models.PaymentHistory

	err := pr.db.QueryRow(query, reference).Scan(&paymentHistory)
	if err != nil {
		return nil, err
	}

	return &paymentHistory, nil
}

func (pr *ImplPaymentRepository) UpdatePaymentHistoryStatus(status *string) error {
	query := `UPDATE payments_history
	SET payment_status = $1
	WHERE order_id= $2
	`
	_, err := pr.db.Exec(query, status)
	return err
}
