package productRepository

import (
	"biller/pkg/utils"
	"biller/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProductById(t *testing.T) {
	var repo = NewLocalProductRepository(tests.MockProducts)

	tests := []struct {
		id          string
		expected    *utils.Product
		expectError bool
	}{
		{"1", &tests.MockProducts[0], false},
		{"2", &tests.MockProducts[1], false},
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
	repo := NewLocalProductRepository(tests.MockProducts)

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
