package utils

import (
	"bytes"
	"image"
)

type GetValidNumberFromInputOptions struct {
	ShouldBePositive bool
	FloatPrecision   int
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
		vatCategory VATCategory,
	) *Product
}

// UnitType represents a type for product units
type UnitType string

const (
	UnitPiece UnitType = "piece"
	UnitKg    UnitType = "kg"
)

// VAT category
type VATCategory string

const (
	A VATCategory = "A"
	B VATCategory = "B"
)

// Define a custom type for actions
type BillAction string

// Use constants for predefined actions
const (
	AddProduct    BillAction = "addProduct"
	RemoveProduct BillAction = "removeProduct"
	PrintBill     BillAction = "printBill"
	SaveAndExit   BillAction = "saveAndExit"
	Exit          BillAction = "exit"
)

type Product struct {
	Id          string      `json:"id"`
	Name        string      `json:"name"`
	UnitPrice   float64     `json:"unitPrice"`
	UnitType    UnitType    `json:"unitType"`
	Stock       float64     `json:"stock"`
	VATCategory VATCategory `json:"vatCategory"`
}

type ProductWithQuantityFromBill struct {
	Product
	Quantity float64
}

// bills

type BillData struct {
	Id        string
	Products  []ProductWithQuantityFromBill
	Subtotal  float64
	VATAmount float64
	Total     float64
	CreatedAt string
}

type BillProduct struct {
	Id       string
	Quantity float64
}

type Bill struct {
	Id        string        `json:"id"`
	Products  []BillProduct `json:"products"` // not ok, bill product repo needs to be put in place
	Subtotal  float64       `json:"subtotal"`
	Total     float64       `json:"total"`
	CreatedAt string        `json:"createdAt"`
}

type BillRepositoryInterface interface {
	GetBills() []Bill
	GetBillById(string) (*Bill, error)
	AddBill(
		products []BillProduct,
		subtotal float64,
		total float64,
		createdAt string,
		id string,
	) *Bill
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
	GetBillItem([]Product, string) (string, float64)
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

// Bill Storage Handler
type BillStorageHandlerInterface interface {
	GetAll() ([]Bill, error)
	Get(string) (*Bill, error)
	Add(Bill) (*Bill, error)
}

// barcode Generator
type BarcodeGeneratorInterface interface {
	GenerateCode(data string, width int, height int) (image.Image, error)
}

// File Handler
type FileHandlerInterface interface {
	Save(data *bytes.Buffer, fileName string) string
	OpenFile(filePath string, command string)
}
