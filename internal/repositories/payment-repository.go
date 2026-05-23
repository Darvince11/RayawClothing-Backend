package repositories

import (
	"database/sql"
	"fmt"
	"rayaw-api/internal/models"
	"strings"
)

type PaymentRepository interface {
	AddPaymentHistory(paymentHistory *models.PaymentHistory, tx *sql.Tx) (int, error)
	GetPaymentHistoryByReference(reference string) (*models.PaymentHistory, error)
	UpdatePaymentHistory(updateReq *models.UpdatePaymentHistoryRequest, reference string) error
}

type ImplPaymentRepository struct {
	db *sql.DB
}

func NewImplPaymentRepository(db *sql.DB) *ImplPaymentRepository {
	return &ImplPaymentRepository{db: db}
}

func (pr *ImplPaymentRepository) AddPaymentHistory(paymentHistory *models.PaymentHistory, tx *sql.Tx) (int, error) {
	query := `INSERT INTO payments_history (order_id, user_id, reference, currency, payment_method, amount, payment_status, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING payment_id`
	var paymentId int
	err := tx.QueryRow(query, paymentHistory.OrderId, paymentHistory.UserId, paymentHistory.Reference, paymentHistory.Currency, paymentHistory.PaymentMethod, paymentHistory.Amount, paymentHistory.PaymentStatus, paymentHistory.CreatedAt, paymentHistory.UpdatedAt).Scan(&paymentId)
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

func (pr *ImplPaymentRepository) UpdatePaymentHistory(updateReq *models.UpdatePaymentHistoryRequest, reference string) error {
	clauses := []string{}
	args := []any{}
	argIndex := 1

	if updateReq.Currency != nil {
		clauses = append(clauses, fmt.Sprintf("currency = $%d", argIndex))
		args = append(args, *updateReq.Currency)
		argIndex++
	}
	if updateReq.PaymentMethod != nil {
		clauses = append(clauses, fmt.Sprintf("payment_method = $%d", argIndex))
		args = append(args, *updateReq.PaymentMethod)
		argIndex++
	}
	if updateReq.PaymentStatus != nil {
		clauses = append(clauses, fmt.Sprintf("payment_status = $%d", argIndex))
		args = append(args, *updateReq.PaymentStatus)
		argIndex++
	}

	if len(clauses) == 0 {
		return nil // No fields to update
	}

	args = append(args, reference)
	argIndex++

	query := fmt.Sprint(`
	UPDATE payments_history
	SET `, strings.Join(clauses, ", "), fmt.Sprintf("WHERE reference = $%d", argIndex))

	_, err := pr.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}
