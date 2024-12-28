package inputHandler_test

import (
	"biller/mocks"
	"biller/pkg/inputHandler"
	"biller/pkg/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetValidIntFromInput(t *testing.T) {
	readerMock := new(mocks.InputReaderMock)
	validatorMock := new(mocks.InputValidatorMock)
	handler := inputHandler.NewInputHandler(readerMock, validatorMock)

	options := utils.GetValidNumberFromInputOptions{ShouldBePositive: true}

	// Simulate valid input
	readerMock.On("GetInput", mock.Anything).Return("42", nil).Once()
	validatorMock.On("ValidateInt", "42").Return(true).Once()
	validatorMock.On("ValidatePositive", "42").Return(true, nil).Once()

	result := handler.GetValidIntFromInput("Enter a number:", options)
	assert.Equal(t, 42, result)

	// Simulate invalid input followed by valid input
	readerMock.On("GetInput", mock.Anything).Return("-10", nil).Once()
	readerMock.On("GetInput", mock.Anything).Return("5", nil).Once()
	validatorMock.On("ValidateInt", "-10").Return(true).Once()
	validatorMock.On("ValidatePositive", "-10").Return(false, nil).Once()
	validatorMock.On("ValidateInt", "5").Return(true).Once()
	validatorMock.On("ValidatePositive", "5").Return(true, nil).Once()

	result = handler.GetValidIntFromInput("Enter a positive number:", options)
	assert.Equal(t, 5, result)

	readerMock.AssertExpectations(t)
	validatorMock.AssertExpectations(t)
}

func TestGetValidFloatFromInput(t *testing.T) {
	readerMock := new(mocks.InputReaderMock)
	validatorMock := new(mocks.InputValidatorMock)
	handler := inputHandler.NewInputHandler(readerMock, validatorMock)

	options := utils.GetValidNumberFromInputOptions{ShouldBePositive: true}

	// Simulate valid input
	readerMock.On("GetInput", mock.Anything).Return("42.5", nil).Once()
	validatorMock.On("ValidateFloat", "42.5").Return(true).Once()
	validatorMock.On("ValidatePositive", "42.5").Return(true, nil).Once()

	result := handler.GetValidFloatFromInput("Enter a number:", options)
	assert.Equal(t, 42.5, result)

	// Simulate invalid input followed by valid input
	readerMock.On("GetInput", mock.Anything).Return("-10.7", nil).Once()
	readerMock.On("GetInput", mock.Anything).Return("15.3", nil).Once()
	validatorMock.On("ValidateFloat", "-10.7").Return(true).Once()
	validatorMock.On("ValidatePositive", "-10.7").Return(false, nil).Once()
	validatorMock.On("ValidateFloat", "15.3").Return(true).Once()
	validatorMock.On("ValidatePositive", "15.3").Return(true, nil).Once()

	result = handler.GetValidFloatFromInput("Enter a positive number:", options)
	assert.Equal(t, 15.3, result)

	readerMock.AssertExpectations(t)
	validatorMock.AssertExpectations(t)
}
