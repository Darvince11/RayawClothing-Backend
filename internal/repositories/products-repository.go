package repositories

import (
	"database/sql"
	"fmt"
	"rayaw-api/internal/models"

	"github.com/lib/pq"
)

type ProductsRepository interface {
	AddProduct(product *models.Product) (int, error)
	AddProductVariation(variation *models.ProductVariation) error
	GetAllProducts(cursor, limit int) (*[]models.Product, error)
	GetProductsVariation(productId []int) (*[]models.ProductVariation, error)
	GetProductsById(productsId []int) (*[]models.Product, error)
	UpdateProduct(productId int, newProduct *models.Product) error
	DeleteProduct(productId int) error
}

type ImplProductsRepository struct {
	db *sql.DB
}

func NewProductsRepository(db *sql.DB) ProductsRepository {
	return &ImplProductsRepository{db: db}
}

func (pr *ImplProductsRepository) AddProduct(product *models.Product) (int, error) {
	query := `INSERT INTO products (image_url, product_name, product_description, price, category, product_status)
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id
	`
	var productId int
	err := pr.db.QueryRow(query, product.Image_url, product.Product_name, product.Product_Description, product.Price, product.Category, product.Product_Status).Scan(&productId)
	return productId, err
}

func (pr *ImplProductsRepository) AddProductVariation(variation *models.ProductVariation) error {
	query := `INSERT INTO product_variants (product_id, product_size, color)
	VALUES ($1, $2, $3);`
	_, err := pr.db.Exec(query, variation.ProductId, pq.Array(variation.ProductSize), pq.Array(variation.Color))
	return err
}

func (pr *ImplProductsRepository) GetAllProducts(cursor int, limit int) (*[]models.Product, error) {
	query := `SELECT id, image_url, product_name, price, category, product_status, created_at FROM products 
	WHERE id > $1
	ORDER BY id ASC
	LIMIT $2;
	`
	var products []models.Product
	results, err := pr.db.Query(query, cursor, limit)
	if err != nil {
		return nil, err
	}

	for results.Next() {
		var product models.Product
		results.Scan(&product.Id, &product.Image_url, &product.Product_name, &product.Price, &product.Category, &product.Product_Status, &product.CreatedAt)
		products = append(products, product)
	}

	return &products, nil
}
func (pr *ImplProductsRepository) GetProductsById(productsId []int) (*[]models.Product, error) {
	placeholders := []string{}
	for index := range productsId {
		placeholders = append(placeholders, fmt.Sprintf("$%v", index+1))
	}

	query := `SELECT * FROM products
	WHERE id = ANY($1);
	`
	var products []models.Product
	rows, err := pr.db.Query(query, pq.Array(productsId))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.Id, &product.Image_url, &product.Product_name, &product.Product_Description, &product.Price, &product.Category, &product.Product_Status, &product.CreatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return &products, err
}

func (pr *ImplProductsRepository) GetProductsVariation(productId []int) (*[]models.ProductVariation, error) {
	query := `SELECT product_size, color FROM product_variants
	WHERE product_id = ANY($1) ;
	`
	prodVars := []models.ProductVariation{}
	row, err := pr.db.Query(query, pq.Array(productId))

	for row.Next() {
		var prodVar models.ProductVariation
		err := row.Scan(pq.Array(&prodVar.ProductSize), pq.Array(&prodVar.Color))
		if err != nil {
			return nil, err
		}
		prodVars = append(prodVars, prodVar)
	}

	return &prodVars, err
}

func (pr *ImplProductsRepository) UpdateProduct(productId int, newProduct *models.Product) error {
	query := `UPDATE products
	SET image_url=$1, product_name=$2, product_description=$3, price=$4, category=$5, product_status=$6
	WHERE id=$7;
	`
	_, err := pr.db.Exec(query, newProduct.Image_url, newProduct.Product_name, newProduct.Product_Description, newProduct.Price, newProduct.Category, newProduct.Product_Status, productId)
	return err
}
func (pr *ImplProductsRepository) DeleteProduct(productId int) error {
	query := `DELETE FROM products
	WHERE id=$1;
	`
	_, err := pr.db.Exec(query, productId)
	return err
}
