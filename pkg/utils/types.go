package utils

type GetValidNumberFromInputOptions struct {
	ShouldBePositive bool
}

// printer
type PrinterInterface interface {
	Print(data string)
}

// productss
type ProductRepositoryInterface interface {
	GetProducts() []Product
	GetProductById(string) (*Product, error)
	UpdateStock(string, float64) error
	IsProductValid(string) bool
}

// UnitType represents a type for product units
type UnitType string

const (
	UnitPiece UnitType = "piece"
	UnitKg    UnitType = "kg"
)

type Product struct {
	Id        string
	Name      string
	UnitPrice float64
	UnitType  UnitType
	Stock     float64
}

// bills
type BillItem struct {
	Id       string
	Quantity float64
}

type BillConfig struct {
	BillsDir      string
	BillRowLength int
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
