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
	GetProductById(id string) (*Product, error)
	IsProductValid(id string) bool
}

type Product struct {
	Id        string
	Name      string
	UnitPrice float64
}

// bills
type BillItem struct {
	Id       string
	Quantity int
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
}

// input handler
type InputHandlerInterface interface {
	GetTableName() string
	GetBillItem([]Product, string) (string, int)
	GetTip() float64
}
