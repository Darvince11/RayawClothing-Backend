package models

import "time"

type Category string

const (
	Men        Category = "men"
	Women      Category = "women"
	Kids       Category = "kids"
	Unisex     Category = "unisex"
	TShirts    Category = "t-shirts"
	Shirts     Category = "shirts"
	Hoodies    Category = "hoodies"
	Jackets    Category = "jackets"
	Sweaters   Category = "sweaters"
	Dresses    Category = "dresses"
	Pants      Category = "pants"
	Jeans      Category = "jeans"
	Shorts     Category = "shorts"
	Skirts     Category = "skirts"
	Activewear Category = "activewear"
	Sleepwear  Category = "sleepwear"
	Underwear  Category = "underwear"
	Swimwear   Category = "swimwear"
)

type ProductStatus string

const (
	Available  ProductStatus = "available"
	OutOfStock ProductStatus = "out-of-stock"
)

type Product struct {
	Id                  int           `json:"id"`
	Image_url           string        `json:"image_url"`
	Product_name        string        `json:"product_name"`
	Product_Description string        `json:"product_description,omitempty"`
	Price               float64       `json:"price"`
	Category            Category      `json:"category"`
	Product_Status      ProductStatus `json:"product_status"`
	CreatedAt           time.Time     `json:"created_at"`
}

type ProductVariation struct {
	ProductId   int
	ProductSize []string `json:"product_size"`
	Color       []string `json:"color"`
}

type CreateProductRequest struct {
	Image_url    string           `json:"image_url"`
	Product_name string           `json:"product_name"`
	Price        int              `json:"price"`
	Description  string           `json:"description"`
	Variations   ProductVariation `json:"variations"`
}

type GetProductsByIdResponse struct {
	Product
	ProductSize []string `json:"product_size"`
	Color       []string `json:"color"`
}
