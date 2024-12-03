package tests

import (
	"biller/pkg/productRepository"
	"biller/pkg/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProductById(t *testing.T) {
	var repo = productRepository.NewLocalProductRepository(mockProducts)

	tests := []struct {
		id          string
		expected    *utils.Product
		expectError bool
	}{
		{"1", &mockProducts[0], false},
		{"2", &mockProducts[1], false},
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
	repo := productRepository.NewLocalProductRepository(mockProducts)

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
