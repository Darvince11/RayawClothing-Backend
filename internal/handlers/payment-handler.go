package handlers

import (
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
	"strconv"
)

type PaymentHandler struct {
	ps     *services.PaymentService
	config *config.Config
	os     services.OrderServiceInterface
}

func NewPaymentHandler(ps *services.PaymentService, config *config.Config) *PaymentHandler {
	return &PaymentHandler{ps: ps, config: config}
}

func (ph *PaymentHandler) SetPaymentService(os services.OrderServiceInterface) {
	ph.os = os
}

func (ph *PaymentHandler) VerifyPaymentWebhook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Webhook started...")
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
	paymentMethod := models.PaymentMethod(paystackResponse.Data.Channel)
	paymentStatus := models.PaymentStatus(paystackResponse.Data.Status)

	updateReq := models.UpdatePaymentHistoryRequest{
		Reference:     &paystackResponse.Data.Reference,
		Currency:      &paystackResponse.Data.Currency,
		PaymentMethod: &paymentMethod,
		PaymentStatus: &paymentStatus,
	}

	orderId, err := ph.ps.UpdatePaymentHistory(&updateReq, paystackResponse.Data.Reference)

	if err != nil {
		http.Error(w, "Failed to update payment history", http.StatusInternalServerError)
		return
	}

	//update order status
	err = ph.os.UpdateOrderStatus(orderId, models.OrderStatusPaid)
	if err != nil {
		http.Error(w, "Failed to update order status", http.StatusInternalServerError)
		return
	}

	//If successful, return 200
	w.WriteHeader(http.StatusOK)
	fmt.Println("Payment verified successfully")
}

func (ph *PaymentHandler) GetAllPaymentHistoryByUserId(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	paymentHistory, err := ph.ps.GetAllPaymentHistoryByUserId(userId)

	res := models.Response[[]models.PaymentHistory]{
		Success: true,
		Message: "Payment history fetched successfully",
		Data:    paymentHistory,
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func (ph *PaymentHandler) GetPaymentHistoryByReference(w http.ResponseWriter, r *http.Request) {
	ref := r.URL.Query().Get("reference")
	paymentHistory, err := ph.ps.GetPaymentHistoryByReference(ref)
	if err != nil {
		http.Error(w, "Failed to fetch payment history", http.StatusInternalServerError)
		return
	}

	res := models.Response[models.PaymentHistory]{
		Success: true,
		Message: "Payment history fetched successfully",
		Data:    *paymentHistory,
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}

}

func VerifyPaymentEvents(ph *PaymentHandler, updateReq models.UpdatePaymentHistoryRequest, paystackResponse *models.PaystackVerifyResponse) {
	// _, err := ph.ps.UpdatePaymentHistory(&updateReq, paystackResponse.Data.Reference)

	// if err != nil {
	// 	log.Printf("Failed to update payment history: %v", err)
	// 	return
	// }

	//update order status
	// err := ph.os.UpdateOrderStatus(uuid.MustParse("7427199a-376f-4aa9-adac-ee69c8c4677b"), models.OrderStatusPaid)
	// if err != nil {
	// 	log.Printf("Failed to update order status: %v", err)
	// 	return
	// }
}
