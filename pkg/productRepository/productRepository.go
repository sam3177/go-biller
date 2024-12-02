package productRepository

import (
	"fmt"
	"slices"
)

type Product struct {
	Id        string
	Name      string
	UnitPrice float64
}

type ProductRepositoryInterface interface {
	GetProductById(id string) (*Product, error) // Fetch product details by ID
	IsProductValid(id string) bool              // Check if a product is valid
	GetProducts() []Product                     // Optional: List all products
}

type ProductRepository struct {
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (repo *ProductRepository) GetProducts() []Product {
	return ProductsCatalog
}

func (repo *ProductRepository) GetProductById(id string) (*Product, error) {
	index := slices.IndexFunc(repo.GetProducts(), func(product Product) bool {
		return product.Id == id
	})

	if index == -1 {
		return nil, fmt.Errorf("product with ID %v not found", id)
	}

	return &repo.GetProducts()[index], nil
}

func (repo *ProductRepository) IsProductValid(id string) bool {
	index := slices.IndexFunc(repo.GetProducts(), func(product Product) bool {
		return product.Id == id
	})

	return index != -1
}
