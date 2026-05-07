package handlers

import (
	"encoding/json"
	"net/http"
	"rayaw-api/internal/models"
	"rayaw-api/internal/services"
	"strconv"
)

type OrderHandler struct {
	os *services.OrderService
}

func NewOrderHandler(os *services.OrderService) *OrderHandler {
	return &OrderHandler{os: os}
}

func (oh *OrderHandler) AddOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order

	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = oh.os.AddOrder(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
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

func (oh *OrderHandler) AddOrderItems(w http.ResponseWriter, r *http.Request) {
	var orderItems []models.OrderItem
	err := json.NewDecoder(r.Body).Decode(&orderItems)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = oh.os.AddOrderItems(&orderItems)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
