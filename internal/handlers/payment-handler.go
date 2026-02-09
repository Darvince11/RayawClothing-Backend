package handlers

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"rayaw-api/internal/config"
	"rayaw-api/internal/models"
	"rayaw-api/internal/services"
)

type PaymentHandler struct {
	ps     *services.PaymentService
	config *config.Config
}

func NewPaymentHandler(ps *services.PaymentService) *PaymentHandler {
	return &PaymentHandler{ps: ps}
}

func (ph *PaymentHandler) InitializePayment(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}

	//decode data
	type InitData struct {
		models.PaymentInitRequest
		Callback_Url string `json:"callback_url"`
	}
	var initData InitData
	err := json.NewDecoder(r.Body).Decode(&initData)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	fmt.Println("Received payment initialization request:", initData)

	initData.Callback_Url = ph.config.PaystackCallbackUrl

	jsonBody, err := json.Marshal(initData)
	if err != nil {
		http.Error(w, "Failed to marshal request body", http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest(
		"POST",
		"https://api.paystack.co/transaction/initialize",
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Authorization", "Bearer "+ph.config.PaystackSecretKey)
	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to send request", http.StatusInternalServerError)
		return
	}

	var paymentInitailizationRespone models.PaystackInitResponse

	err = json.NewDecoder(response.Body).Decode(&paymentInitailizationRespone)
	if err != nil {
		http.Error(w, "Failed to decode response", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(paymentInitailizationRespone)
}

func (ph *PaymentHandler) VerifyPaymentWebhook(w http.ResponseWriter, r *http.Request) {
	signature := r.Header.Get("X-Paystack-Signature")
	if signature == "" {
		http.Error(w, "Missing signature", http.StatusBadRequest)
		return
	}

	paystackResByte, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	//Verify payment signature
	mac := hmac.New(sha512.New, []byte(ph.config.PaystackSecretKey))
	_, err = mac.Write(paystackResByte)
	if err != nil {
		http.Error(w, "Failed to write HMAC", http.StatusInternalServerError)
		return
	}
	expectedSignatureByte := mac.Sum(nil)

	expectedSignature := hex.EncodeToString(expectedSignatureByte)

	if signature != expectedSignature {
		http.Error(w, "Invalid signature", http.StatusBadRequest)
		return
	}

	//Parse the payment details from paystack into a variable
	var paystackResponse models.PaystackVerifyResponse

	err = json.Unmarshal(paystackResByte, &paystackResponse)
	if err != nil {
		http.Error(w, "Failed to unmarshal paystack response", http.StatusInternalServerError)
		return
	}

	//Update payment history status in the database
	status := paystackResponse.Data.Status
	err = ph.ps.UpdatePaymentHistoryStatus(&status)

	//If successful, return 200
	w.WriteHeader(http.StatusOK)
	fmt.Println("Payment verified successfully")
}

func (ph *PaymentHandler) VerifyPayment(w http.ResponseWriter, r *http.Request) {
	ref := r.URL.Query().Get("reference")
	//fetch the payment history
	paymentHistory, err := ph.ps.GetPaymentHistoryByReference(ref)
	if err != nil {
		http.Error(w, "Failed to fetch payment history", http.StatusInternalServerError)
		return
	}
	//check status
	if paymentHistory.PaymentStatus != "success" {
		http.Error(w, "Payment not successful", http.StatusPaymentRequired)
		return
	}

	//return response
	err = json.NewEncoder(w).Encode(paymentHistory)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
