package handlers

import (
	"encoding/json"
	"net/http"
	"rayaw-api/internal/models"
	"rayaw-api/internal/services"
	"strconv"

	"github.com/google/uuid"
)

type OrderHandler struct {
	os *services.OrderService
}

func NewOrderHandler(os *services.OrderService) *OrderHandler {
	return &OrderHandler{os: os}
}

func (oh *OrderHandler) AddOrder(w http.ResponseWriter, r *http.Request) {
	var orderRequest models.CreateOrderRequest

	err := json.NewDecoder(r.Body).Decode(&orderRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	addOrderResponse, err := oh.os.AddOrder(&orderRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.Response[models.AddOrderResponse]{
		Success: true,
		Message: "Order created succesfully",
		Data:    *addOrderResponse,
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error ecoding json"+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (oh *OrderHandler) GetOrdersByUserId(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	orders, err := oh.os.GetOrdersByUserId(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.Response[[]models.Order]{
		Success: true,
		Message: "Orders fetched successfully",
		Data:    *orders,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (oh *OrderHandler) GetOrderById(w http.ResponseWriter, r *http.Request) {
	orderIdStr := r.PathValue("id")
	if orderIdStr == "" {
		http.Error(w, "order id is required", http.StatusBadRequest)
		return
	}

	orderId, err := uuid.Parse(orderIdStr)
	if err != nil {
		http.Error(w, "invalid order id", http.StatusBadRequest)
		return
	}

	orderWithItems, err := oh.os.GetOrderById(orderId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.Response[models.OrderWithItems]{
		Success: true,
		Message: "Order fetched successfully",
		Data:    *orderWithItems,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
