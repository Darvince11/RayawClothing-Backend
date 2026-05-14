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
	var orderRequest models.CreateOrderRequest

	err := json.NewDecoder(r.Body).Decode(&orderRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	authUrl, err := oh.os.AddOrder(&orderRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.Response[map[string]string]{
		Success: true,
		Message: "Order created succesfully",
		Data:    map[string]string{"authorization_url": authUrl},
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
