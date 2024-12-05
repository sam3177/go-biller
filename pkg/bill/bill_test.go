package bill

import (
	"biller/pkg/printer"
	"biller/pkg/productRepository"
	"biller/pkg/utils"
	"biller/tests"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testProductsRepo = productRepository.NewLocalProductRepository(tests.MockProducts)

var testBillConfig = utils.BillConfig{
	BillsDir:      "./bills",
	BillRowLength: utils.BILL_ROW_LENGTH,
}
var testTermimalPrinter = printer.NewTerminalPrinter()

func TestAddProduct(t *testing.T) {

	bill := NewBill(testProductsRepo, testTermimalPrinter, testBillConfig)

	// Test adding a valid product
	bill.AddProduct("1", 2)
	assert.Equal(t, 1, len(bill.GetProducts()))
	assert.Equal(t, "1", bill.GetProducts()[0].Id)
	assert.Equal(t, 2, bill.GetProducts()[0].Quantity)

	// Test adding the same product again
	bill.AddProduct("1", 3)
	assert.Equal(t, 1, len(bill.GetProducts()))
	assert.Equal(t, "1", bill.GetProducts()[0].Id)
	assert.Equal(t, 5, bill.GetProducts()[0].Quantity)

	// Test adding another valid product
	bill.AddProduct("2", 1)
	assert.Equal(t, 2, len(bill.GetProducts()))
	assert.Equal(t, "2", bill.GetProducts()[1].Id)
	assert.Equal(t, 1, bill.GetProducts()[1].Quantity)

	// Test adding an invalid product
	bill.AddProduct("4", 1)
	assert.Equal(t, 2, len(bill.GetProducts())) // No change in length
}

func TestRemoveProduct(t *testing.T) {
	bill := NewBill(testProductsRepo, testTermimalPrinter, testBillConfig)

	// Add 3 products to the bill
	bill.AddProduct("1", 4)
	bill.AddProduct("2", 45)
	bill.AddProduct("3", 7)

	// Test removing a valid product with quantity less than existing
	bill.RemoveProduct("1", 2)
	assert.Equal(t, 3, len(bill.GetProducts()))
	assert.Equal(t, 2, bill.GetProducts()[0].Quantity)

	// Test removing a valid product with quantity equal to existing
	bill.RemoveProduct("2", 45)
	assert.Equal(t, 2, len(bill.GetProducts()))

	// Test removing a valid product with quantity more than existing
	bill.RemoveProduct("3", 10)
	assert.Equal(t, 1, len(bill.GetProducts()))

	// Test removing an invalid product
	bill.RemoveProduct("4", 1)
	assert.Equal(t, 1, len(bill.GetProducts())) // No change in length

	// Test removing a valid product that does not exist in the bill
	bill.RemoveProduct("2", 1)
	assert.Equal(t, 1, len(bill.GetProducts())) // No change in length
}

func TestCalculateTotal(t *testing.T) {
	bill := NewBill(testProductsRepo, testTermimalPrinter, testBillConfig)

	bill.AddProduct("1", 4)
	bill.AddProduct("2", 3)
	bill.AddProduct("3", 1)

	// Calculate the total
	total := bill.CalculateTotal()

	// Assert the expected total
	assert.Equal(t, 13.0, total)
}

func TestFormatBill(t *testing.T) {
	bill := NewBill(testProductsRepo, testTermimalPrinter, testBillConfig)
	bill.SetTableName("Table 1")

	// Add some products to the bill
	bill.AddProduct("1", 4)
	bill.AddProduct("2", 3)

	// Set the tip
	bill.SetTip(34.6)

	// Format the bill
	formattedBill := bill.FormatBill()

	expectedText := `              ----Bill---- 
Table name: Table 1 
----------------------------------------
Product 1
            4 X 1.00                4.00 
Product 2
            3 X 2.00                6.00 
----------------------------------------

Subtotal                           10.00 

Tip                                34.60 
----------------------------------------

Total                              44.60 

`

	// Assert the expected formatted bill
	assert.Equal(t, expectedText, formattedBill)
}

func TestSaveBill(t *testing.T) {
	bill := NewBill(testProductsRepo, testTermimalPrinter, testBillConfig)

	//make bills folder and cleanup at the end with defer
	os.Mkdir("bills", 0755)
	defer os.RemoveAll(bill.BillsDir)

	fileName := bill.SaveBill()

	file, error := os.Open(bill.BillsDir + "/" + fileName)

	if error != nil {
		t.Errorf("file not found")
	}

	defer file.Close()
}
