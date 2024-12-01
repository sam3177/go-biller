package utils

import (
	productRepository "biller/pkg/productsRepo"
	"testing"
)

var products = []productRepository.Product{
	{Id: "1", Name: "one"},
	{Id: "2", Name: "two"},
	{Id: "3", Name: "three"},
}

func TestGetProductById(t *testing.T) {

	testCases := []struct {
		id          string
		expected    *productRepository.Product
		expectError bool
	}{
		{"1", &products[0], false},
		{"2", &products[1], false},
		{"4", nil, true},
	}

	for _, testCase := range testCases {
		result, error := GetProductById(products, testCase.id)

		if testCase.expectError {
			if error == nil {
				t.Errorf("expected error for id %v, got no error", testCase.id)
			}
		} else {
			if error != nil {
				t.Errorf("did not expect error for id '2', got %v", error)
			}

			if result == nil || *result != *testCase.expected {
				t.Errorf("Expected %v, got %v", products[1], result)
			}
		}
	}
}

func TestCheckIfProductIsValid(t *testing.T) {

	testCases := []struct {
		id       string
		expected bool
	}{
		{"1", true},
		{"2", true},
		{"4", false},
	}
	for _, testCase := range testCases {
		result := CheckIfProductIsValid(products, testCase.id)
		if result != testCase.expected {
			t.Errorf("Expected %v for id %v, got %v", testCase.expected, testCase.id, result)
		}
	}
}
