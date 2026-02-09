package services

import (
	"rayaw-api/internal/models"
	"rayaw-api/internal/repositories"
)

type PaymentService struct {
	pr *repositories.ImplPaymentRepository
}

func NewPaymentService(pr *repositories.ImplPaymentRepository) *PaymentService {
	return &PaymentService{pr: pr}
}

func (ps *PaymentService) AddPaymentHistory(paymentHistory *models.PaymentHistory) (int, error) {
	return ps.pr.AddPaymentHistory(paymentHistory)
}

func (ps *PaymentService) GetPaymentHistoryByReference(reference string) (*models.PaymentHistory, error) {
	return ps.pr.GetPaymentHistoryByReference(reference)
}

func (ps *PaymentService) UpdatePaymentHistoryStatus(status *string) error {
	return ps.pr.UpdatePaymentHistoryStatus(status)
}
