package inputValidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateFloat(t *testing.T) {
	validator := NewInputValidator()

	tests := []struct {
		input    string
		expected bool
	}{
		{"123.45", true},
		{"-123.45", true},
		{"abc", false},
		{"", false},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := validator.ValidateFloat(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestValidateInt(t *testing.T) {
	validator := NewInputValidator()

	tests := []struct {
		input    string
		expected bool
	}{
		{"123", true},
		{"-123", true},
		{"abc", false},
		{"", false},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := validator.ValidateInt(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestValidatePositive(t *testing.T) {
	validator := NewInputValidator()

	tests := []struct {
		input    string
		expected bool
	}{
		{"123", true},
		{"0", false},
		{"-123", false},
		{"abc", false},
		{"", false},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := validator.ValidatePositive(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestValidateMinLength(t *testing.T) {
	validator := NewInputValidator()

	tests := []struct {
		input    string
		length   int
		expected bool
	}{
		{"hello", 3, true},
		{"hi", 3, false},
		{"", 1, false},
		{"world", 5, true},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := validator.ValidateMinLength(test.input, test.length)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestValidateMaxLength(t *testing.T) {
	validator := NewInputValidator()

	tests := []struct {
		input    string
		length   int
		expected bool
	}{
		{"hello", 10, true},
		{"hello", 5, false},
		{"", 1, true},
		{"world!", 5, false},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := validator.ValidateMaxLength(test.input, test.length)
			assert.Equal(t, test.expected, result)
		})
	}
}
