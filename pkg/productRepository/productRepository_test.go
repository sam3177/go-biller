package productRepository

import (
	"biller/mocks"
	"biller/pkg/utils"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProductById(t *testing.T) {
	testProductsJsonStorageHandler := &mocks.ProductsJSONStorageHandlerMock{}
	repo := NewLocalProductRepository(testProductsJsonStorageHandler)

	tests := []struct {
		id       string
		expected *utils.Product
		error    error
	}{
		{"1", &mocks.MockProducts[0], nil},
		{"2", &mocks.MockProducts[1], nil},
		{"4", nil, errors.New("")},
	}

	for _, test := range tests {
		testProductsJsonStorageHandler.On("GetProduct", test.id).Return(test.expected, test.error).Once()

		result, err := repo.GetProductById(test.id)
		if test.error != nil {
			assert.Error(t, err)
			assert.Nil(t, result)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}

func TestIsProductValid(t *testing.T) {
	testProductsJsonStorageHandler := &mocks.ProductsJSONStorageHandlerMock{}
	repo := NewLocalProductRepository(testProductsJsonStorageHandler)

	tests := []struct {
		id       string
		expected bool
	}{
		{"1", true},
		{"2", true},
		{"4", false},
	}

	for _, test := range tests {
		var err error = nil
		if !test.expected {
			err = errors.New("")
		}
		testProductsJsonStorageHandler.On("GetProduct", test.id).Return(&utils.Product{}, err).Once()

		result := repo.IsProductValid(test.id)
		assert.Equal(t, test.expected, result)
	}
}

func TestIsEnoughProductInStock(t *testing.T) {
	testProductsJsonStorageHandler := &mocks.ProductsJSONStorageHandlerMock{}
	repo := NewLocalProductRepository(testProductsJsonStorageHandler)

	tests := []struct {
		productId    string
		productStock float64
		error        error
		quantity     float64
		expected     bool
	}{
		{"1", 40.0, nil, 30.0, true},
		{"2", 30.0, nil, 199.0, false},
		{"4", 3, errors.New(""), 5, false},
	}

	for _, test := range tests {
		testProductsJsonStorageHandler.On("GetProduct", test.productId).Return(&utils.Product{Stock: test.productStock}, test.error).Once()

		result := repo.IsEnoughProductInStock(test.productId, test.quantity)
		assert.Equal(t, test.expected, result)
	}
}

func TestUpdateStock(t *testing.T) { // 			id:          "",
	testProductsJsonStorageHandler := &mocks.ProductsJSONStorageHandlerMock{}
	repo := NewLocalProductRepository(testProductsJsonStorageHandler)

	tests := []struct {
		id             string
		quantity       float64
		initialStock   float64
		expectedStock  float64
		canHaveDecimal bool
		expectedError  error
	}{
		{"1", 10.0, 20.0, 30.0, true, nil},
		{"2", -5.0, 10.0, 5.0, true, nil},
		{"3", 1.5, 10.0, 11.5, true, nil},
		{"4", 1.5, 10.0, 10.0, false, fmt.Errorf("this product is sold by piece, so a decimal point quantity is not valid in this case")},
		{"5", -15.0, 10.0, 10.0, true, fmt.Errorf("the available stock for this product is %f, but you requested %f", 10.0, 15.0)},
		{"6", 5.0, 10.0, 15.0, false, nil},
	}

	for _, test := range tests {
		product := &utils.Product{
			Id:       test.id,
			Stock:    test.initialStock,
			UnitType: utils.UnitPiece,
		}
		if test.canHaveDecimal {
			product.UnitType = utils.UnitKg
		}

		testProductsJsonStorageHandler.On("GetProduct", test.id).Return(product, nil).Times(3)

		updatedProduct := &utils.Product{
			Id:       test.id,
			Stock:    test.expectedStock,
			UnitType: utils.UnitPiece,
		}
		if test.canHaveDecimal {
			updatedProduct.UnitType = utils.UnitKg
		}

		if test.expectedError == nil {
			testProductsJsonStorageHandler.On("UpdateProduct", *updatedProduct).Return(nil).Once()
		}

		newStock, err := repo.UpdateStock(test.id, test.quantity)
		if test.expectedError != nil {
			assert.Error(t, err)
			assert.Equal(t, test.expectedError.Error(), err.Error())
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expectedStock, newStock)
		}
	}
}

func TestAddProduct(t *testing.T) {
	testProductsJsonStorageHandler := &mocks.ProductsJSONStorageHandlerMock{}
	repo := NewLocalProductRepository(testProductsJsonStorageHandler)

	tests := []struct {
		name        string
		productName string
		unitPrice   float64
		unitType    interface{} // Allow UnitType to be any type for validation
		stock       float64
		vatCategory interface{} // Allow UnitType to be any type for validation
		expectPanic bool
	}{
		{
			name:        "Valid Inputs",
			productName: "Apple",
			unitPrice:   1.99,
			unitType:    utils.UnitKg,
			stock:       10.0,
			vatCategory: utils.A,
			expectPanic: false,
		},
		{
			name:        "Missing Product Name",
			productName: "",
			unitPrice:   1.99,
			unitType:    utils.UnitKg,
			stock:       10,
			vatCategory: utils.A,
			expectPanic: true,
		},
		{
			name:        "Invalid UnitPrice",
			productName: "Apple",
			unitPrice:   -1,
			unitType:    utils.UnitKg,
			stock:       10,
			vatCategory: utils.A,
			expectPanic: true,
		},
		{
			name:        "Invalid UnitType",
			productName: "Apple",
			unitPrice:   1.99,
			unitType:    "invalidType",
			stock:       10,
			vatCategory: utils.A,
			expectPanic: true,
		},
		{
			name:        "Negative Stock",
			productName: "Apple",
			unitPrice:   1.99,
			unitType:    utils.UnitKg,
			stock:       -5,
			vatCategory: utils.A,
			expectPanic: true,
		},
		{
			name:        "Invalid VAT Category",
			productName: "Apple",
			unitPrice:   1.99,
			unitType:    utils.UnitKg,
			stock:       -5,
			vatCategory: "invalidVATCategory",
			expectPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if r != nil && !tt.expectPanic {
					t.Errorf("Expected no panic, but got panic: %v", r)
				}
				if r == nil && tt.expectPanic {
					t.Errorf("Expected panic, but no panic occurred")
				}
			}()

			newProduct := utils.Product{
				Name:        tt.productName,
				UnitPrice:   tt.unitPrice,
				UnitType:    tt.unitType.(utils.UnitType),
				Stock:       tt.stock,
				VATCategory: tt.vatCategory.(utils.VATCategory),
			}

			if !tt.expectPanic {
				testProductsJsonStorageHandler.On("AddProduct", utils.Product{
					Name:        tt.productName,
					UnitPrice:   tt.unitPrice,
					UnitType:    tt.unitType.(utils.UnitType),
					Stock:       tt.stock,
					VATCategory: tt.vatCategory.(utils.VATCategory),
				}).Return(&newProduct, nil).Once()
			}

			// Call NewProduct with the current test case data
			product := repo.AddProduct(tt.productName, tt.unitPrice, tt.unitType.(utils.UnitType), tt.stock, tt.vatCategory.(utils.VATCategory))
			if product.Name != tt.productName {
				t.Errorf("Expected Name '%s', got '%s'", tt.productName, product.Name)
			}

			if product.UnitPrice != tt.unitPrice {
				t.Errorf("Expected UnitPrice %f, got %f", tt.unitPrice, product.UnitPrice)
			}

			if product.UnitType != tt.unitType {
				t.Errorf("Expected UnitType '%v', got '%v'", tt.unitType, product.UnitType)
			}

			if product.Stock != tt.stock {
				t.Errorf("Expected Stock %f, got %f", tt.stock, product.Stock)
			}

			if product.VATCategory != tt.vatCategory {
				t.Errorf("Expected VAT categoty %v, got %v", tt.vatCategory, product.VATCategory)
			}
		})
	}
}
