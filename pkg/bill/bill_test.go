package bill

import (
	"biller/pkg/productRepository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddProduct(t *testing.T) {

	productRepo := productRepository.NewProductRepository()

	bill := &Bill{
		tableName:   "34",
		products:    []BillItem{},
		ProductRepo: productRepo,
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
