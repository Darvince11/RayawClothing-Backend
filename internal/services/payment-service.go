package services

import (
	"bytes"
	"encoding/json"
	"net/http"
	"rayaw-api/internal/config"
	"rayaw-api/internal/models"
	"rayaw-api/internal/repositories"
)

type PaymentService struct {
	pr     *repositories.ImplPaymentRepository
	config *config.Config
	client *http.Client
}

func NewPaymentService(pr *repositories.ImplPaymentRepository, config *config.Config, client *http.Client) *PaymentService {
	return &PaymentService{pr: pr, config: config, client: client}
}

func (ps *PaymentService) AddPaymentHistory(paymentHistory *models.PaymentHistory) (int, error) {
	return ps.pr.AddPaymentHistory(paymentHistory)
}

func (ps *PaymentService) GetPaymentHistoryByReference(reference string) (*models.PaymentHistory, error) {
	return ps.pr.GetPaymentHistoryByReference(reference)
}

func (ps *PaymentService) GetAllPaymentHistoryByUserId(userId int) ([]models.PaymentHistory, error) {
	return ps.pr.GetAllPaymentHistoryByUserId(userId)
}

func (ps *PaymentService) UpdatePaymentHistoryStatus(status *string) error {
	return ps.pr.UpdatePaymentHistoryStatus(status)
}

func (ps *PaymentService) InitializePayment(email string, amount float64) (string, error) {
	paymentInitData := models.PaymentInit{
		Email:        email,
		Amount:       int(amount * 100),
		Callback_Url: ps.config.PaystackCallbackUrl,
	}

	jsonBody, err := json.Marshal(&paymentInitData)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(
		"POST",
		"https://api.paystack.co/transaction/initialize",
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+ps.config.PaystackSecretKey)
	req.Header.Set("Content-Type", "application/json")

	response, err := ps.client.Do(req)
	if err != nil {
		return "", err
	}

	var paymentInitailizationRespone models.PaystackInitResponse

	err = json.NewDecoder(response.Body).Decode(&paymentInitailizationRespone)
	if err != nil {
		return "", err
	}
	return paymentInitailizationRespone.Data.AuthorizationURL, nil
}
