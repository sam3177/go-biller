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
	}

	if num < 0 {
		fmt.Println(errorMessage)
	}

	return num > 0
}
