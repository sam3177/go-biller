package bill

import (
	"biller/mocks"
	"biller/pkg/billFormatter"
	"biller/pkg/printer"
	"biller/pkg/utils"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testTermimalPrinter = printer.NewTerminalPrinter(utils.BILL_ROW_LENGTH)

var testFormatter = billFormatter.NewBillTerminalFormatter()

func findProductByID(products []utils.BillProduct, id string) *utils.BillProduct {
	for _, product := range products {
		if product.Id == id {
			return &product
		}
	}
	return nil
}

func TestAddProduct(t *testing.T) {
	testProductsRepo := &mocks.ProductRepositoryMock{}
	testBillRepo := &mocks.BillRepositoryMock{}
	bill := NewBillingHandler(testProductsRepo, testBillRepo, testTermimalPrinter, testFormatter, "./bills")

	tests := []struct {
		name              string
		productID         string
		quantity          float64
		isProductValid    bool
		updateStockReturn float64
		updateStockError  error
		expectedLength    int
		expectedQuantity  float64
		expectedError     bool
	}{
		{
			name:              "Add valid product",
			productID:         "1",
			quantity:          2,
			isProductValid:    true,
			updateStockReturn: 5.0,
			updateStockError:  nil,
			expectedLength:    1,
			expectedQuantity:  2,
		},
		{
			name:              "Add same product again",
			productID:         "1",
			quantity:          3,
			isProductValid:    true,
			updateStockReturn: 5.0,
			updateStockError:  nil,
			expectedLength:    1,
			expectedQuantity:  5,
		},
		{
			name:              "Add another valid product",
			productID:         "2",
			quantity:          1,
			isProductValid:    true,
			updateStockReturn: 10.0,
			updateStockError:  nil,
			expectedLength:    2,
			expectedQuantity:  1,
		},
		{
			name:              "Add product with floating point quantity (unitType kg)",
			productID:         "3",
			quantity:          5.56,
			isProductValid:    true,
			updateStockReturn: 15.0,
			updateStockError:  nil,
			expectedLength:    3,
			expectedQuantity:  5.56,
		},
		{
			name:              "Add product with floating point quantity (unitType piece, fails)",
			productID:         "2",
			quantity:          5.56,
			isProductValid:    true,
			updateStockReturn: 0,
			updateStockError:  fmt.Errorf("this product is sold by piece, so a decimal point quantity is not valid in this case"),
			expectedLength:    3,
			expectedQuantity:  1, // No change
		},
		{
			name:              "Add valid product but quantity exceeds stock (fails)",
			productID:         "2",
			quantity:          999,
			isProductValid:    true,
			updateStockReturn: 0,
			updateStockError:  fmt.Errorf("the available stock for this product is whatever, but you requested 999"),
			expectedLength:    3,
			expectedQuantity:  1, // No change
		},
		{
			name:              "Add invalid product",
			productID:         "4",
			quantity:          1,
			isProductValid:    false,
			updateStockReturn: 0,
			updateStockError:  nil,
			expectedLength:    3, // No change
			expectedQuantity:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange: set expectations
			testProductsRepo.On("IsProductValid", tt.productID).Return(tt.isProductValid).Once()
			if tt.isProductValid {
				testProductsRepo.On("UpdateStock", tt.productID, -tt.quantity).Return(tt.updateStockReturn, tt.updateStockError).Once()
			}

			// Act
			bill.AddProduct(tt.productID, tt.quantity)

			// Assert
			assert.Equal(t, tt.expectedLength, len(bill.GetProducts()))
			if tt.expectedLength > 0 {
				product := findProductByID(bill.GetProducts(), tt.productID)
				if product != nil {
					assert.Equal(t, tt.expectedQuantity, product.Quantity)
				}
			}

			// Verify mock expectations
			testProductsRepo.AssertExpectations(t)
		})
	}
}

func populateBillWithProducts(bill *BillingHandler, testProductsRepo *mocks.ProductRepositoryMock) {
	productsToAdd := []struct {
		productId string
		quantity  float64
	}{
		{
			productId: "1",
			quantity:  4,
		},
		{
			productId: "2",
			quantity:  5,
		},
		{
			productId: "3",
			quantity:  7,
		},
	}

	for _, product := range productsToAdd {
		testProductsRepo.On("IsProductValid", product.productId).Return(true).Once()
		testProductsRepo.On("UpdateStock", product.productId, -product.quantity).Return(45.5, nil).Once()

		bill.AddProduct(product.productId, product.quantity)
	}
}

func mockGetProductByIdForAddedProducts(bill *BillingHandler, testProductsRepo *mocks.ProductRepositoryMock) {
	repoProducts := []utils.Product{
		{Name: "Product 1", UnitPrice: 4.0, UnitType: "kg", VATCategory: utils.A},
		{Name: "Product 2", UnitPrice: 4.2, UnitType: "piece", VATCategory: utils.A},
		{Name: "Product 3", UnitPrice: 34.1, UnitType: "kg", VATCategory: utils.A},
	} // 4, 5, 7
	for idx, product := range bill.products {
		testProductsRepo.On("GetProductById", product.Id).Return(
			&repoProducts[idx], nil).Once()
	}
}

func TestRemoveProduct(t *testing.T) {
	testProductsRepo := &mocks.ProductRepositoryMock{}
	testBillRepo := &mocks.BillRepositoryMock{}
	bill := NewBillingHandler(testProductsRepo, testBillRepo, testTermimalPrinter, testFormatter, "./bills")

	// Arrange: Add initial products to the bill
	populateBillWithProducts(bill, testProductsRepo)

	// Define test cases
	tests := []struct {
		name                  string
		productID             string
		quantity              float64
		expectedLength        int
		expectedQuantity      float64
		quantityToAddBackToDB float64
		isProductValid        bool
		error                 error
	}{
		{
			name:                  "Remove valid product with quantity less than existing",
			productID:             "1",
			quantity:              2.0,
			expectedLength:        3,
			expectedQuantity:      2.0,
			quantityToAddBackToDB: 2.0,
			isProductValid:        true,
		},
		{
			name:                  "Remove valid product with quantity equal to existing",
			productID:             "2",
			quantity:              5.0,
			expectedLength:        2,
			expectedQuantity:      0, // Product should be removed
			quantityToAddBackToDB: 5.0,
			isProductValid:        true,
		},
		{
			name:                  "Remove valid product with quantity more than existing",
			productID:             "3",
			quantity:              10.0,
			expectedLength:        1,
			expectedQuantity:      0, // Product should be removed
			quantityToAddBackToDB: 7.0,
			isProductValid:        true,
		},
		{
			name:                  "Remove invalid product",
			productID:             "4",
			quantity:              1.0,
			expectedLength:        1, // No change
			expectedQuantity:      0,
			quantityToAddBackToDB: 0,
			isProductValid:        false,
			error:                 errors.New(""),
		},
		{
			name:                  "Remove valid product not in bill",
			productID:             "2",
			quantity:              1.0,
			expectedLength:        1, // No change
			expectedQuantity:      0,
			quantityToAddBackToDB: 0,
			isProductValid:        false,
			error:                 errors.New(""),
		},
	}

	// Execute test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange: set expectations
			testProductsRepo.On("IsProductValid", tt.productID).Return(tt.isProductValid).Once()
			if tt.isProductValid {
				testProductsRepo.On("UpdateStock", tt.productID, tt.quantityToAddBackToDB).Return(tt.expectedQuantity, tt.error).Once()
			}
			// Act
			bill.RemoveProduct(tt.productID, tt.quantity)

			// Assert
			assert.Equal(t, tt.expectedLength, len(bill.GetProducts()))

			if tt.expectedLength > 0 {
				product := findProductByID(bill.GetProducts(), tt.productID)
				if product != nil {
					assert.Equal(t, tt.expectedQuantity, product.Quantity)
				}
			}
		})
	}

}

func TestCalculateTotal(t *testing.T) {
	testProductsRepo := &mocks.ProductRepositoryMock{}
	testBillRepo := &mocks.BillRepositoryMock{}
	bill := NewBillingHandler(testProductsRepo, testBillRepo, testTermimalPrinter, testFormatter, "./bills")

	// Arrange: Add initial products to the bill
	populateBillWithProducts(bill, testProductsRepo)

	mockGetProductByIdForAddedProducts(bill, testProductsRepo)

	// Calculate the total
	total := bill.CalculateTotal()

	// Assert the expected total
	assert.Equal(t, 275.7, total)
}

func TestCalculateVAT(t *testing.T) {
	testProductsRepo := &mocks.ProductRepositoryMock{}
	testBillRepo := &mocks.BillRepositoryMock{}
	bill := NewBillingHandler(testProductsRepo, testBillRepo, testTermimalPrinter, testFormatter, "./bills")

	// Arrange: Add initial products to the bill
	populateBillWithProducts(bill, testProductsRepo)

	mockGetProductByIdForAddedProducts(bill, testProductsRepo)

	// Calculate the total
	totalVAT := bill.CalculateVAT()
	// Assert the expected total
	assert.Equal(t, 24.81, totalVAT)
}

func TestSaveBill(t *testing.T) {
	testProductsRepo := &mocks.ProductRepositoryMock{}
	testBillRepo := &mocks.BillRepositoryMock{}

	bill := NewBillingHandler(testProductsRepo, testBillRepo, testTermimalPrinter, testFormatter, "./bills")

	//make bills folder and cleanup at the end with defer
	os.Mkdir("bills", 0755)
	defer os.RemoveAll(bill.BillsDir)

	testBillRepo.On("AddBill", bill.products, bill.CalculateTotal(), bill.CalculateTotal()).Return(&utils.Bill{}).Once()

	fileName := bill.SaveBill()

	file, error := os.Open(bill.BillsDir + "/" + fileName)

	if error != nil {
		t.Errorf("file not found")
	}

	defer file.Close()
}
