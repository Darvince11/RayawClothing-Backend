package routes

import (
	"net/http"
	"rayaw-api/internal/handlers"
)

type ProductsRoutes struct {
	mux            *http.ServeMux
	productHandler *handlers.ProductsHandler
}

func NewProductsRoutes(mux *http.ServeMux, productHandler *handlers.ProductsHandler) *ProductsRoutes {
	return &ProductsRoutes{mux: mux, productHandler: productHandler}
}

func (pr *ProductsRoutes) RegisterRoutes() {
	pr.mux.HandleFunc("GET /products", pr.productHandler.GetAllProductsHandler)
}
