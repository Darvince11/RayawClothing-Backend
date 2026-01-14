package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rayaw-api/internal/models"
	"rayaw-api/internal/services"
	"strconv"
)

type ProductsHandler struct {
	productService *services.ProductService
}

func NewProductsHandler(productService *services.ProductService) *ProductsHandler {
	return &ProductsHandler{productService: productService}
}

func GetNextCursor(currentCursor, limit int) int {
	return currentCursor + limit
}

func (ph *ProductsHandler) GetAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	//get the query params
	cursor := r.URL.Query().Get("cursor")
	limit := r.URL.Query().Get("limit")

	var cursorInt int
	//convert to int
	if cursor == "" {
		cursorInt = 0
	} else {
		cursor, err := strconv.Atoi(cursor)
		cursorInt = cursor
		if err != nil {
			http.Error(w, "Invalid cursor", http.StatusBadRequest)
			return
		}
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		http.Error(w, "Invalid limit", http.StatusBadRequest)
		return
	}

	//call the get products service and pass in the values
	products, err := ph.productService.GetAllProducts(cursorInt, limitInt)
	if err != nil {
		http.Error(w, "Error fetching products", http.StatusInternalServerError)
		return
	}

	//return response
	response := models.Response[[]models.Product]{
		Success: true,
		Message: "retrieved products successfully",
		Data:    *products,
		Meta: map[string]any{
			"next":  GetNextCursor(cursorInt, limitInt),
			"limit": limitInt,
		},
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error econding response")
		return
	}
}

func (ph *ProductsHandler) GetProductByIdHandler(w http.ResponseWriter, r *http.Request) {
	//get the poduct id from the url
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid limit", http.StatusBadRequest)
		return
	}
	//fetch the product
	product, err := ph.productService.GetProductById(id)
	if err != nil {
		http.Error(w, "Error fetching product", http.StatusInternalServerError)
		return
	}

	//return response
	response := models.Response[models.GetProductsByIdResponse]{
		Success: true,
		Message: "retrieved product successfully",
		Data:    *product,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error econding response")
		return
	}
}
