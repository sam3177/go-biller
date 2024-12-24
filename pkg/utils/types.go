package utils

import "bytes"

type GetValidNumberFromInputOptions struct {
	ShouldBePositive bool
}

// printer
type PrinterInterface interface {
	Print(data bytes.Buffer)
	GetRowLength() int
}

// productss
type ProductRepositoryInterface interface {
	GetProducts() []Product
	GetProductById(string) (*Product, error)
	UpdateStock(string, float64) (float64, error)
	IsProductValid(string) bool
	AddProduct(
		name string,
		unitPrice float64,
		unitType UnitType,
		stock float64,
	) *Product
}

// UnitType represents a type for product units
type UnitType string

const (
	UnitPiece UnitType = "piece"
	UnitKg    UnitType = "kg"
)

type Product struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	UnitPrice float64  `json:"unitPrice"`
	UnitType  UnitType `json:"unitType"`
	Stock     float64  `json:"stock"`
}

type ProductWithQuantityFromBill struct {
	Product
	Quantity float64
}

// bills

type BillData struct {
	TableName string
	Products  []ProductWithQuantityFromBill
	Tip       float64
	Subtotal  float64
	Total     float64
}

type BillItem struct {
	Id       string
	Quantity float64
}

// input validator
type InputValidatorInterface interface {
	ValidateFloat(string) bool
	ValidateInt(string) bool
	ValidatePositive(string) bool
	ValidateMinLength(string, int) bool
	ValidateMaxLength(string, int) bool
}

// input handler
type InputHandlerInterface interface {
	GetTableName() string
	GetBillItem([]Product, string) (string, float64)
	GetTip() float64
}

type InputReaderInterface interface {
	GetInput(string) (string, error)
}

// Products Storage Handler
type ProductsStorageHandlerInterface interface {
	GetAllProducts() ([]Product, error)
	GetProduct(string) (*Product, error)
	UpdateProduct(Product) error
	AddProduct(Product) (*Product, error)
}

// Bill Formatter
type BillFormatterInterface interface {
	FormatBill(billData BillData, rowLength int) bytes.Buffer
}
