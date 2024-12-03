package productRepository

import (
	"biller/pkg/utils"
	"fmt"
	"slices"
)

type LocalProductRepository struct {
	products []utils.Product
}

func NewLocalProductRepository(products []utils.Product) *LocalProductRepository {
	return &LocalProductRepository{products: products}
}

func (repo *LocalProductRepository) GetProducts() []utils.Product {
	return repo.products
}

func (repo *LocalProductRepository) GetProductById(id string) (*utils.Product, error) {
	index := slices.IndexFunc(repo.GetProducts(), func(product utils.Product) bool {
		return product.Id == id
	})

	if index == -1 {
		return nil, fmt.Errorf("product with ID %v not found", id)
	}

	return &repo.GetProducts()[index], nil
}

func (repo *LocalProductRepository) IsProductValid(id string) bool {
	index := slices.IndexFunc(repo.GetProducts(), func(product utils.Product) bool {
		return product.Id == id
	})

	return index != -1
}
