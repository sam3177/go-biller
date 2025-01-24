package mocks

import (
	"biller/pkg/utils"

	"github.com/stretchr/testify/mock"
)

// input reader
type InputReaderMock struct {
	mock.Mock
}

func (m *InputReaderMock) GetInput(prompt string) (string, error) {
	args := m.Called(prompt)
	return args.String(0), args.Error(1)
}

// input validator
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

// products JSON storage handler
type ProductsJSONStorageHandlerMock struct {
	mock.Mock
}

func (m *ProductsJSONStorageHandlerMock) GetAllProducts() ([]utils.Product, error) {
	args := m.Called()

	return args.Get(0).([]utils.Product), args.Error(1)
}

// func (m *ProductsJSONStorageHandlerMock) writeProducts(products []utils.Product) error {
// 	args := m.Called(products)

// 	return args.Error(0)
// }

func (m *ProductsJSONStorageHandlerMock) GetProduct(id string) (*utils.Product, error) {
	args := m.Called(id)

	return args.Get(0).(*utils.Product), args.Error(1)
}

func (m *ProductsJSONStorageHandlerMock) UpdateProduct(productData utils.Product) error {
	args := m.Called(productData)

	return args.Error(0)
}

func (m *ProductsJSONStorageHandlerMock) AddProduct(newProduct utils.Product) (*utils.Product, error) {
	args := m.Called(newProduct)

	return args.Get(0).(*utils.Product), args.Error(1)
}

// product Repository
type ProductRepositoryMock struct {
	mock.Mock
}

// Mock method for GetProducts
func (m *ProductRepositoryMock) GetProducts() []utils.Product {
	args := m.Called()
	return args.Get(0).([]utils.Product)
}

// Mock method for GetProductById
func (m *ProductRepositoryMock) GetProductById(id string) (*utils.Product, error) {
	args := m.Called(id)
	return args.Get(0).(*utils.Product), args.Error(1)
}

// Mock method for IsProductValid
func (m *ProductRepositoryMock) IsProductValid(id string) bool {
	args := m.Called(id)
	return args.Bool(0)
}

// Mock method for IsEnoughProductInStock
func (m *ProductRepositoryMock) IsEnoughProductInStock(id string, desiredQuantity float64) bool {
	args := m.Called(id, desiredQuantity)
	return args.Bool(0)
}

// Mock method for AddProduct
func (m *ProductRepositoryMock) AddProduct(name string, unitPrice float64, unitType utils.UnitType, stock float64, vatCategory utils.VATCategory) *utils.Product {
	args := m.Called(name, unitPrice, unitType, stock, vatCategory)
	return args.Get(0).(*utils.Product)
}

// Mock method for UpdateStock
func (m *ProductRepositoryMock) UpdateStock(id string, quantity float64) (float64, error) {
	args := m.Called(id, quantity)
	return args.Get(0).(float64), args.Error(1)
}

// Mock method for CanProductHaveDecimalStock
func (m *ProductRepositoryMock) CanProductHaveDecimalStock(id string) bool {
	args := m.Called(id)
	return args.Bool(0)
}

// bill repository
type BillRepositoryMock struct {
	mock.Mock
}

// Mock method for AddBill
func (m *BillRepositoryMock) AddBill(products []utils.BillProduct, subtotal float64, total float64) *utils.Bill {
	args := m.Called(products, subtotal, total)
	return args.Get(0).(*utils.Bill)
}

// Mock method for GetBillById
func (m *BillRepositoryMock) GetBillById(id string) (*utils.Bill, error) {
	args := m.Called(id)
	return args.Get(0).(*utils.Bill), args.Error(1)
}

// Mock method for GetBills
func (m *BillRepositoryMock) GetBills() []utils.Bill {
	args := m.Called()
	return args.Get(0).([]utils.Bill)
}

// bills JSON storage handler

type BillsJSONStorageHandlerMock struct {
	mock.Mock
}

func (m *BillsJSONStorageHandlerMock) GetAll() ([]utils.Bill, error) {
	args := m.Called()

	return args.Get(0).([]utils.Bill), args.Error(1)
}

func (m *BillsJSONStorageHandlerMock) Get(id string) (*utils.Bill, error) {
	args := m.Called(id)

	return args.Get(0).(*utils.Bill), args.Error(1)
}

func (m *BillsJSONStorageHandlerMock) Add(newBill utils.Bill) (*utils.Bill, error) {
	args := m.Called(newBill)

	return args.Get(0).(*utils.Bill), args.Error(1)
}
