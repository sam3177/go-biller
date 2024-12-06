package mocks

import "github.com/stretchr/testify/mock"

type InputReaderMock struct {
	mock.Mock
}

func (m *InputReaderMock) GetInput(prompt string) (string, error) {
	args := m.Called(prompt)
	return args.String(0), args.Error(1)
}

type InputValidatorMock struct {
	mock.Mock
}

func (m *InputValidatorMock) ValidateInt(value string) bool {
	args := m.Called(value)
	return args.Bool(0)
}

func (m *InputValidatorMock) ValidateFloat(value string) bool {
	args := m.Called(value)
	return args.Bool(0)
}

func (m *InputValidatorMock) ValidatePositive(value string) bool {
	args := m.Called(value)
	return args.Bool(0)
}

func (m *InputValidatorMock) ValidateMinLength(value string, length int) bool {
	args := m.Called(value, length)
	return args.Bool(0)
}

func (m *InputValidatorMock) ValidateMaxLength(value string, length int) bool {
	args := m.Called(value, length)
	return args.Bool(0)
}
