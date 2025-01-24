package productRepository

import (
	"biller/pkg/utils"
	"fmt"
	"math"
)

type LocalProductRepository struct {
	dataHandler utils.ProductsStorageHandlerInterface
}

func NewLocalProductRepository(
	dataHandler utils.ProductsStorageHandlerInterface,
) *LocalProductRepository {
	return &LocalProductRepository{
		dataHandler: dataHandler,
	}
}

func (repo *LocalProductRepository) GetProducts() []utils.Product {
	products, error := repo.dataHandler.GetAllProducts()

	if error != nil {
		fmt.Println(error)
		return nil
	}

	return products
}

func (repo *LocalProductRepository) GetProductById(id string) (*utils.Product, error) {
	product, error := repo.dataHandler.GetProduct(id)

	if error != nil {
		fmt.Println(error)
		return nil, error
	}

	return product, nil
}

func (repo *LocalProductRepository) IsProductValid(id string) bool {
	product, error := repo.GetProductById(id)

	return product != nil && error == nil
}

func (repo *LocalProductRepository) IsEnoughProductInStock(id string, desiredQuantity float64) bool {
	product, error := repo.GetProductById(id)

	if error != nil {
		fmt.Println(error.Error())
		return false
	}

	return product.Stock >= desiredQuantity
}

func (repo *LocalProductRepository) AddProduct(
	name string,
	unitPrice float64,
	unitType utils.UnitType,
	stock float64,
	vatCategory utils.VATCategory,
) *utils.Product {
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

	// check VAT category
	if vatCategory != utils.A && vatCategory != utils.B {
		panic("Product VAT category must be 'A' or 'B'.")
	}

	newProduct, error := repo.dataHandler.AddProduct(utils.Product{
		Name:        name,
		UnitPrice:   unitPrice,
		UnitType:    unitType,
		Stock:       stock,
		VATCategory: vatCategory,
	})

	if error != nil {
		fmt.Println(error)
		return nil
	}

	return newProduct
}

func (repo *LocalProductRepository) UpdateStock(id string, quantity float64) (float64, error) {
	product, error := repo.GetProductById(id)

	if error != nil {
		return 0, error
	}

	if !repo.CanProductHaveDecimalStock(id) && quantity != math.Floor(quantity) {
		return 0, fmt.Errorf("this product is sold by piece, so a decimal point quantity is not valid in this case")
	}

	if quantity < 0 && !repo.IsEnoughProductInStock(id, quantity*-1) {

		return 0, fmt.Errorf("the available stock for this product is %f, but you requested %f", product.Stock, quantity*-1)
	}

	product.Stock = utils.RoundToGivenDecimals(product.Stock+quantity, 3)

	updateProductError := repo.dataHandler.UpdateProduct(*product)

	if updateProductError != nil {
		return 0, updateProductError
	}

	return product.Stock, nil
}

func (repo *LocalProductRepository) CanProductHaveDecimalStock(id string) bool {
	product, error := repo.GetProductById(id)

	if error != nil {
		fmt.Println(error.Error())
		return false
	}

	return product.UnitType == utils.UnitKg
}
