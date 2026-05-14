package routes

import (
	"net/http"
	"rayaw-api/internal/handlers"
)

type OrderRoutes struct {
	mux          *http.ServeMux
	orderHandler *handlers.OrderHandler
}

func NewOrderRoutes(mux *http.ServeMux, orderHandler *handlers.OrderHandler) *OrderRoutes {
	return &OrderRoutes{mux: mux, orderHandler: orderHandler}
}

func (or *OrderRoutes) RegisterRoutes() {
	or.mux.HandleFunc("POST /orders", or.orderHandler.AddOrder)
	or.mux.HandleFunc("GET /orders/user/{id}", or.orderHandler.GetOrdersByUserId)
}
