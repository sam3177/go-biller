package productRepository

import (
	"biller/pkg/utils"
	"fmt"
	"math"
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

func (repo *LocalProductRepository) IsEnoughProductInStock(id string, desiredQuantity float64) bool {
	product, error := repo.GetProductById(id)

	if error != nil {
		fmt.Println(error.Error())
		return false
	}

	return product.Stock >= desiredQuantity
}

func (repo *LocalProductRepository) UpdateStock(id string, quantity float64) {
	product, error := repo.GetProductById(id)

	if error != nil {
		fmt.Println(error.Error())
		return
	}

	if !repo.CanProductHaveDecimalStock(id) && quantity != math.Floor(quantity) {
		fmt.Println("This product is sold by piece, so you can not ask for a decimal point quantity.")
		return
	}

	product.Stock += quantity
}

func (repo *LocalProductRepository) CanProductHaveDecimalStock(id string) bool {
	product, error := repo.GetProductById(id)

	if error != nil {
		fmt.Println(error.Error())
		return false
	}

	return product.UnitType == utils.UnitKg
}
