package productRepository

import (
	"biller/mocks"
	"biller/pkg/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProductById(t *testing.T) {
	var repo = NewLocalProductRepository(mocks.MockProducts)

	tests := []struct {
		id          string
		expected    *utils.Product
		expectError bool
	}{
		{"1", &mocks.MockProducts[0], false},
		{"2", &mocks.MockProducts[1], false},
		{"4", nil, true},
	}

	for _, test := range tests {
		result, err := repo.GetProductById(test.id)
		if test.expectError {
			assert.Error(t, err)
			assert.Nil(t, result)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}

func TestIsProductValid(t *testing.T) {
	repo := NewLocalProductRepository(mocks.MockProducts)

	tests := []struct {
		id       string
		expected bool
	}{
		{"1", true},
		{"2", true},
		{"4", false},
	}

	for _, test := range tests {
		result := repo.IsProductValid(test.id)
		assert.Equal(t, test.expected, result)
	}
}

func TestIsEnoughProductInStock(t *testing.T) {
	repo := NewLocalProductRepository(mocks.MockProducts)

	tests := []struct {
		id       string
		quantity float64
		expected bool
	}{
		{"1", 30.0, true},
		{"2", 199.0, false},
		{"4", 3, false},
	}

	for _, test := range tests {
		result := repo.IsEnoughProductInStock(test.id, test.quantity)
		assert.Equal(t, test.expected, result)
	}

}

func TestUpdateStock(t *testing.T) {
	repo := NewLocalProductRepository(mocks.MockProducts)

	tests := []struct {
		id       string
		quantity float64
		expected float64
	}{
		{"2", 199.56, 30}, // it required only integer quantity, so it fails
		{"2", 30.0, 60.0},
		{"4", 3, 999},
	}

	for _, test := range tests {
		repo.UpdateStock(test.id, test.quantity)
		product, error := repo.GetProductById(test.id)
		if error != nil {
			assert.Error(t, error)
		} else {
			assert.Equal(t, test.expected, product.Stock)
		}
	}
}

func TestNewProduct(t *testing.T) {
	tests := []struct {
		name        string
		id          string
		productName string
		unitPrice   float64
		unitType    interface{} // Allow UnitType to be any type for validation
		stock       float64
		expectPanic bool
		expectedID  string
	}{
		{
			name:        "Valid Inputs",
			id:          "",
			productName: "Apple",
			unitPrice:   1.99,
			unitType:    utils.UnitKg,
			stock:       10,
			expectPanic: false,
			expectedID:  "", // Expect a generated UUID
		},
		{
			name:        "Custom ID",
			id:          "aaa",
			productName: "Orange",
			unitPrice:   2.49,
			unitType:    utils.UnitPiece,
			stock:       5,
			expectPanic: false,
			expectedID:  "aaa", // Expect custom UUID passed in
		},
		{
			name:        "Missing Name",
			id:          "",
			productName: "",
			unitPrice:   1.99,
			unitType:    utils.UnitKg,
			stock:       10,
			expectPanic: true,
			expectedID:  "",
		},
		{
			name:        "Invalid UnitPrice",
			id:          "",
			productName: "Apple",
			unitPrice:   -1,
			unitType:    utils.UnitKg,
			stock:       10,
			expectPanic: true,
			expectedID:  "",
		},
		{
			name:        "Invalid UnitType",
			id:          "",
			productName: "Apple",
			unitPrice:   1.99,
			unitType:    "invalidType",
			stock:       10,
			expectPanic: true,
			expectedID:  "",
		},
		{
			name:        "Negative Stock",
			id:          "",
			productName: "Apple",
			unitPrice:   1.99,
			unitType:    utils.UnitKg,
			stock:       -5,
			expectPanic: true,
			expectedID:  "",
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

			// Call NewProduct with the current test case data
			product := NewProduct(tt.id, tt.productName, tt.unitPrice, tt.unitType.(utils.UnitType), tt.stock)

			if !tt.expectPanic && tt.expectedID == "" && product.Id == "" {
				t.Errorf("Expected generated UUID, got empty string")
			}

			// Validate the expected ID if it's defined
			if tt.expectedID != "" && product.Id != tt.expectedID {
				t.Errorf("Expected Id '%s', got '%s'", tt.expectedID, product.Id)
			}

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
		})
	}
}
