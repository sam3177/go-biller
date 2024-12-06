package inputValidator

import (
	"fmt"
	"strconv"
)

type InputValidator struct{}

func NewInputValidator() *InputValidator {
	return &InputValidator{}
}

func (validator *InputValidator) ValidateFloat(value string) bool {
	_, error := strconv.ParseFloat(value, 64)
	if error != nil {
		fmt.Println("Error:", error)
	}

	return error == nil
}

func (validator *InputValidator) ValidateInt(value string) bool {
	_, error := strconv.ParseInt(value, 10, 0)
	if error != nil {
		fmt.Println("Error:", error)
	}
	return error == nil
}

func (validator *InputValidator) ValidatePositive(value string) bool {
	errorMessage := "Error: The value must be a positive number"
	num, error := strconv.ParseFloat(value, 64)

	if error != nil {
		fmt.Println("Error:", error)
		return false
	}

	if num < 0 {
		fmt.Println(errorMessage)
	}

	return num > 0
}

func (validator *InputValidator) ValidateMinLength(value string, length int) bool {
	fmt.Println("asdfasd", length)
	if len(value) < length {
		fmt.Printf("The value must have a minimum length of %d.\n", length)
	}

	return len(value) >= length
}
func (validator *InputValidator) ValidateMaxLength(value string, length int) bool {
	if len(value) >= length {
		fmt.Printf("The value must have a maximum length of %d.\n", length)
	}
	return len(value) < length
}
