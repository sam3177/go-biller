package productRepository

import (
	"biller/pkg/utils"
	"fmt"
	"math"
	"slices"

	"github.com/google/uuid"
)

type LocalProductRepository struct {
	products []utils.Product
}

func NewProduct(
	id string,
	name string,
	unitPrice float64,
	unitType utils.UnitType,
	stock float64,
) *utils.Product {

	//id check
	var productId string
	if id != "" {
		productId = id
	} else {
		productId = uuid.NewString()
	}

	// name check
	if name == "" {
		panic("Product name is mandatory.")
	}

	// unit price check
	if unitPrice <= 0 {
		panic("Product unitPrice must be greater than 0.")
	}

	// unit type check
	if unitType != utils.UnitPiece && unitType != utils.UnitKg {
		panic("Product unitType must be 'piece' or 'kg'.")
	}

	// stock check
	if stock < 0 {
		panic("Product stock must be 0 or greater.")
	}

	return &utils.Product{
		Id:        productId,
		Name:      name,
		UnitPrice: unitPrice,
		UnitType:  unitType,
		Stock:     stock,
	}
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

func (repo *LocalProductRepository) UpdateStock(id string, quantity float64) error {
	product, error := repo.GetProductById(id)

	if error != nil {
		return error
	}

	if !repo.CanProductHaveDecimalStock(id) && quantity != math.Floor(quantity) {
		return fmt.Errorf("this product is sold by piece, so a decimal point quantity is not valid in this case")
	}

	if quantity < 0 && !repo.IsEnoughProductInStock(id, quantity*-1) {

		return fmt.Errorf("the available stock for this product is %f, but you requested %f", product.Stock, quantity*-1)
	}

	product.Stock += quantity

	return nil
}

func (repo *LocalProductRepository) CanProductHaveDecimalStock(id string) bool {
	product, error := repo.GetProductById(id)

	if error != nil {
		fmt.Println(error.Error())
		return false
	}

	return product.UnitType == utils.UnitKg
}
