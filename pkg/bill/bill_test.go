package bill

import (
	"biller/pkg/productRepository"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockProductRepository struct {
	products []productRepository.Product
}

func (m *MockProductRepository) GetProducts() []productRepository.Product {
	return m.products
}

func (m *MockProductRepository) GetProductById(id string) (*productRepository.Product, error) {
	for _, product := range m.products {
		if product.Id == id {
			return &product, nil
		}
	}
	return nil, fmt.Errorf("product with ID %v not found", id)
}

func (m *MockProductRepository) IsProductValid(id string) bool {
	for _, product := range m.products {
		if product.Id == id {
			return true
		}
	}
	return false
}

func TestAddProduct(t *testing.T) {

	testProductsRepo := &MockProductRepository{
		products: []productRepository.Product{
			{Id: "1", Name: "Product 1"},
			{Id: "2", Name: "Product 2"},
		},
	}

	bill := &Bill{
		tableName:   "34",
		products:    []BillItem{},
		ProductRepo: testProductsRepo,
	}

	// Test adding a valid product
	bill.AddProduct("1", 2)
	assert.Equal(t, 1, len(bill.GetProducts()))
	assert.Equal(t, "1", bill.GetProducts()[0].Id)
	assert.Equal(t, 2, bill.GetProducts()[0].Quantity)

	// Test adding the same product again
	bill.AddProduct("1", 3)
	assert.Equal(t, 1, len(bill.GetProducts()))
	assert.Equal(t, "1", bill.GetProducts()[0].Id)
	assert.Equal(t, 5, bill.GetProducts()[0].Quantity)

	// Test adding another valid product
	bill.AddProduct("2", 1)
	assert.Equal(t, 2, len(bill.GetProducts()))
	assert.Equal(t, "2", bill.GetProducts()[1].Id)
	assert.Equal(t, 1, bill.GetProducts()[1].Quantity)

	// Test adding an invalid product
	bill.AddProduct("3", 1)
	assert.Equal(t, 2, len(bill.GetProducts())) // No change in length
}
