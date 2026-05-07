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
	query := `INSERT INTO payments_history (order_id, reference, currency, payment_method, amount, payment_status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING payment_id`
	var paymentId int
	err := pr.db.QueryRow(query, paymentHistory.OrderId, paymentHistory.Reference, paymentHistory.Currency, paymentHistory.PaymentMethod, paymentHistory.Amount, paymentHistory.PaymentStatus, paymentHistory.CreatedAt, paymentHistory.UpdatedAt).Scan(&paymentId)
	return paymentId, err
}

func (pr *ImplPaymentRepository) GetPaymentHistoryByReference(reference string) (*models.PaymentHistory, error) {
	query := `SELECT * FROM payments_history WHERE reference = $1`

	var paymentHistory models.PaymentHistory

	err := pr.db.QueryRow(query, reference).Scan(&paymentHistory.Id, &paymentHistory.OrderId, &paymentHistory.Reference, &paymentHistory.Currency, &paymentHistory.PaymentMethod, &paymentHistory.Amount, &paymentHistory.PaymentStatus, &paymentHistory.CreatedAt, &paymentHistory.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &paymentHistory, nil
}

func (pr *ImplPaymentRepository) GetAllPaymentHistoryByUserId(userId int) ([]models.PaymentHistory, error) {
	query := `SELECT (reference, payment_method, amount, payment_status, created_at) FROM PAYMENTS_HISTORY WHERE user_id=$1`
	rows, err := pr.db.Query(query, userId)

	var PaymentHistory []models.PaymentHistory
	for rows.Next() {
		var paymentHistory models.PaymentHistory
		rows.Scan(&paymentHistory.Reference, &paymentHistory.PaymentMethod, &paymentHistory.Amount, &paymentHistory.PaymentStatus, &paymentHistory.CreatedAt)
		PaymentHistory = append(PaymentHistory, paymentHistory)
	}
	return PaymentHistory, err
}

func (pr *ImplPaymentRepository) UpdatePaymentHistoryStatus(status *string) error {
	query := `UPDATE payments_history
	SET payment_status = $1
	WHERE order_id= $2
	`
	_, err := pr.db.Exec(query, status)
	return err
}
