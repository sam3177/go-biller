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
