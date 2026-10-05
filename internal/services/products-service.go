package services

import (
	"rayaw-api/internal/models"
	"rayaw-api/internal/repositories"
)

type ProductService struct {
	productRepo repositories.ProductsRepository
}

func NewProductService(productRepo repositories.ProductsRepository) *ProductService {
	return &ProductService{productRepo: productRepo}
}

func (ps *ProductService) AddProductToStore(product *models.Product, variant *models.ProductVariation) error {
	productId, err := ps.productRepo.AddProduct(product)
	if err != nil {
		return err
	}
	variant.ProductId = productId
	return ps.productRepo.AddProductVariation(variant)
}

func (ps *ProductService) GetAllProducts(cursor, limit int) (*[]models.Product, error) {
	return ps.productRepo.GetAllProducts(cursor, limit)
}

func (ps *ProductService) GetProductsById(productsId []int) (*[]models.GetProductsByIdResponse, error) {
	productsResponse := []models.GetProductsByIdResponse{}
	products, err := ps.productRepo.GetProductsById(productsId)
	if err != nil {
		return nil, err
	}
	variants, err := ps.productRepo.GetProductsVariation(productsId)
	if err != nil {
		return nil, err
	}

	for pIndex, product := range *products {
		var productResponse models.GetProductsByIdResponse
		productResponse.Product = product
		productResponse.ProductSize = (*variants)[pIndex].ProductSize
		productResponse.Color = (*variants)[pIndex].Color
		productsResponse = append(productsResponse, productResponse)
	}
	return &productsResponse, nil
}

func (ps *ProductService) UpdateProduct(productId int, newProduct *models.Product) error {
	return ps.productRepo.UpdateProduct(productId, newProduct)
}

func (ps *ProductService) DeleteProduct(productId int) error {
	return ps.productRepo.DeleteProduct(productId)
}
