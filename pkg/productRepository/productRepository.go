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
	GetProducts() []Product
	GetProductById(id string) (*Product, error)
	IsProductValid(id string) bool
}

type LocalProductRepository struct {
	products []Product
}

func NewLocalProductRepository(products []Product) *LocalProductRepository {
	return &LocalProductRepository{products: products}
}

func (repo *LocalProductRepository) GetProducts() []Product {
	return repo.products
}

func (repo *LocalProductRepository) GetProductById(id string) (*Product, error) {
	index := slices.IndexFunc(repo.GetProducts(), func(product Product) bool {
		return product.Id == id
	})

	if index == -1 {
		return nil, fmt.Errorf("product with ID %v not found", id)
	}

	return &repo.GetProducts()[index], nil
}

func (repo *LocalProductRepository) IsProductValid(id string) bool {
	index := slices.IndexFunc(repo.GetProducts(), func(product Product) bool {
		return product.Id == id
	})

	return index != -1
}
