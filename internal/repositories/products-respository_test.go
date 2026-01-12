package repositories

import (
	"rayaw-api/internal/models"
	"rayaw-api/internal/tests"
	"testing"
)

func TestProductsRepository(t *testing.T) {
	db := tests.SetupTestDB(t)
	if db == nil {
		t.Error("db is nil")
	}

	productsRepository := NewProductsRepository(db)

	product := models.Product{
		Image_url:           "image url",
		Product_name:        "Dress",
		Product_Description: "A beautiful dress",
		Price:               120,
		Category:            models.Dresses,
		Product_Status:      models.Available,
	}

	//test AddProduct
	productId, err := productsRepository.AddProduct(&product)
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
	t.Logf("product id: %v", productId)

	//test GetAllProducts
	products, err := productsRepository.GetAllProducts(1, 10)
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
	t.Logf("products: %v", products)

	productsRepository.GetProductById(1)

	//test GetProductById
	productQuery, err := productsRepository.GetProductById(productId)
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
	t.Logf("product: %v", productQuery)
}
