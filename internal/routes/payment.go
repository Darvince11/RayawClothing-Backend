package routes

import (
	"net/http"
	"rayaw-api/internal/handlers"
)

type PaymentRoutes struct {
	mux            *http.ServeMux
	paymentHandler *handlers.PaymentHandler
}

func NewPaymentRoutes(mux *http.ServeMux, paymentHandler *handlers.PaymentHandler) *PaymentRoutes {
	return &PaymentRoutes{mux: mux, paymentHandler: paymentHandler}
}

func (pr *PaymentRoutes) RegisterRoutes() {
	pr.mux.HandleFunc("POST /payment-verify-webhook", pr.paymentHandler.VerifyPaymentWebhook)
	pr.mux.HandleFunc("GET /payments/users/{id}", pr.paymentHandler.GetAllPaymentHistoryByUserId)
}
