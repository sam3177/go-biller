package bill

// import (
// 	"biller/pkg/utils"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// type MockUtils struct {
// 	mock.Mock
// }

// func (m *MockUtils) IsProductValid(products []utils.Product, id string) bool {
// 	args := m.Called(products, id)
// 	return args.Bool(0)
// }

// func TestAddProduct(t *testing.T) {
// 	mockUtils := new(MockUtils)

// 	productsCatalog := []utils.Product{
// 		{Id: "1", Name: "Product 1"},
// 		{Id: "2", Name: "Product 2"},
// 	}

// 	mockUtils.On("IsProductValid", productsCatalog, "1").Return(true)
// 	mockUtils.On("IsProductValid", productsCatalog, "2").Return(true)
// 	mockUtils.On("IsProductValid", productsCatalog, "3").Return(false)

// 	bill := &Bill{
// 		TableName: "34",
// 		Products:  []BillItem{},
// 	}

// 	// Test adding a valid product
// 	bill.AddProduct("1", 2)
// 	assert.Equal(t, 1, len(bill.Products))
// 	assert.Equal(t, "1", bill.Products[0].Id)
// 	assert.Equal(t, 2, bill.Products[0].Quantity)
// }

// // // Test adding the same product again
// // bill.AddProduct("1", 3)
// // assert.Equal(t, 1, len(bill.Products))
// // assert.Equal(t, "1", bill.Products[0].Id)
// // assert.Equal(t, 5, bill.Products[0].Quantity)

// // // Test adding another valid product
// // bill.AddProduct("2", 1)
// // assert.Equal(t, 2, len(bill.Products))
// // assert.Equal(t, "2", bill.Products[1].Id)
// // assert.Equal(t, 1, bill.Products[1].Quantity)

// // // Test adding an invalid product
// // bill.AddProduct("3", 1)
// // assert.Equal(t, 2, len(bill.Products)) // No change in length
// // }
