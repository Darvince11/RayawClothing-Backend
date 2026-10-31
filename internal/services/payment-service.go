package services

import (
	"bytes"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"rayaw-api/internal/config"
	"rayaw-api/internal/models"
	"rayaw-api/internal/repositories"
	"time"

	"github.com/google/uuid"
)

type PaymentService struct {
	pr     *repositories.ImplPaymentRepository
	config *config.Config
	client *http.Client
}

func NewPaymentService(pr *repositories.ImplPaymentRepository, config *config.Config, client *http.Client) *PaymentService {
	return &PaymentService{pr: pr, config: config, client: client}
}

func (ps *PaymentService) AddPaymentHistory(paymentHistory *models.PaymentHistory, tx *sql.Tx) (int, error) {
	return ps.pr.AddPaymentHistory(paymentHistory, tx)
}

func (ps *PaymentService) GetPaymentHistoryByReference(reference string) (*models.PaymentHistory, error) {
	return ps.pr.GetPaymentHistoryByReference(reference)
}

func (ps *PaymentService) GetAllPaymentHistoryByUserId(userId int) ([]models.PaymentHistory, error) {
	fmt.Println("Service user Id: ", userId)
	return ps.pr.GetAllPaymentHistoryByUserId(userId)
}

func (ps *PaymentService) UpdatePaymentHistory(updateReq *models.UpdatePaymentHistoryRequest, reference string) (uuid.UUID, error) {
	return ps.pr.UpdatePaymentHistory(updateReq, reference)
}

func (ps *PaymentService) InitializePayment(email string, amount float64, orderId uuid.UUID, reference string) (string, error) {
	paymentInitData := models.PaymentInit{
		Email:        email,
		Amount:       int(amount * 100),
		Reference:    reference,
		Callback_Url: ps.config.PaystackCallbackUrl + orderId.String(),
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

func (ps *PaymentService) GeneratePaymentId() (string, error) {
	paymentId := "RW"
	timestamp := time.Now().Unix()
	var randomBytes [2]byte

	_, err := rand.Read(randomBytes[:])
	if err != nil {
		return "", err
	}

	randomStr := hex.EncodeToString(randomBytes[:])
	return fmt.Sprintf("%s-%d-%s", paymentId, timestamp, randomStr), nil
}
