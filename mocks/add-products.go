package main

import (
	"database/sql"
	"log"
	"rayaw-api/internal/config"
	"rayaw-api/internal/models"
	"rayaw-api/internal/repositories"
)

type MockProduct struct {
	product   *models.Product
	variation *models.ProductVariation
}

func AddProducts() {
	config := config.Init()
	db, err := sql.Open("postgres", config.DbUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	productRepo := repositories.NewProductsRepository(db)

	//data

	var MockProducts = []MockProduct{
		{
			product: &models.Product{
				Image_url:           "https://res.cloudinary.com/dosnuoybn/image/upload/v1768174076/Oversized_Cutting_Half_Zip_Up_Sweatshirt_Black_jjnpld.jpg",
				Product_name:        "Running Shoes",
				Product_Description: "Lightweight running shoes",
				Price:               250.00,
				Category:            models.Hoodies,
				Product_Status:      models.Available,
			},
			variation: &models.ProductVariation{
				ProductSize: []string{"40", "41", "42"},
				Color:       []string{"black", "white"},
			},
		},
		{
			product: &models.Product{
				Image_url:           "https://res.cloudinary.com/dosnuoybn/image/upload/v1768174205/Knitwear_Trends_2025_tog0fe.jpg",
				Product_name:        "Cotton T-Shirt",
				Product_Description: "Soft cotton t-shirt",
				Price:               80.50,
				Category:            models.TShirts,
				Product_Status:      models.Available,
			},
			variation: &models.ProductVariation{
				ProductSize: []string{"S", "M", "L", "XL"},
				Color:       []string{"red", "blue", "green"},
			},
		},
	}

	for _, mockProduct := range MockProducts {
		productId, err := productRepo.AddProduct(mockProduct.product)
		if err != nil {
			log.Fatal(err)
		}

		varition := mockProduct.variation
		varition.ProductId = productId
		err = productRepo.AddProductVariation(varition)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func main() {
	AddProducts()
}
