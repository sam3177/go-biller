package bill

import (
	"biller/pkg/productRepository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddProduct(t *testing.T) {

	productsCatalog := []productRepository.Product{
		{Id: "1", Name: "Product 1"},
		{Id: "2", Name: "Product 2"},
	}

	productRepo := productRepository.NewProductRepository(productsCatalog)

	bill := &Bill{
		TableName:   "34",
		Products:    []BillItem{},
		ProductRepo: productRepo,
	}

	// Test adding a valid product
	bill.AddProduct("1", 2)
	assert.Equal(t, 1, len(bill.Products))
	assert.Equal(t, "1", bill.Products[0].Id)
	assert.Equal(t, 2, bill.Products[0].Quantity)

	// Test adding the same product again
	bill.AddProduct("1", 3)
	assert.Equal(t, 1, len(bill.Products))
	assert.Equal(t, "1", bill.Products[0].Id)
	assert.Equal(t, 5, bill.Products[0].Quantity)

	// Test adding another valid product
	bill.AddProduct("2", 1)
	assert.Equal(t, 2, len(bill.Products))
	assert.Equal(t, "2", bill.Products[1].Id)
	assert.Equal(t, 1, bill.Products[1].Quantity)

	// Test adding an invalid product
	bill.AddProduct("3", 1)
	assert.Equal(t, 2, len(bill.Products)) // No change in length
}
