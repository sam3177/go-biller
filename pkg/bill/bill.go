package bill

import (
	"bytes"
	"fmt"
	"os"

	"github.com/google/uuid"

	"biller/pkg/utils"
)

// TODO: date on the bill

type Bill struct {
	products    []utils.BillItem
	ProductRepo utils.ProductRepositoryInterface
	Printer     utils.PrinterInterface
	Formatter   utils.BillFormatterInterface
	BillsDir    string
}

func NewBill(
	productRepo utils.ProductRepositoryInterface,
	printer utils.PrinterInterface,
	formatter utils.BillFormatterInterface,
	billsDir string,
) *Bill {
	return &Bill{
		products:    []utils.BillItem{},
		ProductRepo: productRepo,
		Formatter:   formatter,
		Printer:     printer,
		BillsDir:    billsDir,
	}
}

func (bill *Bill) AddProduct(id string, quantity float64) {
	if !bill.ProductRepo.IsProductValid(id) {
		fmt.Printf("Product with ID %v is not a valid product in the system.", id)
		return
	}
	if quantity <= 0 {
		return
	}

	_, removeFromStockError := bill.ProductRepo.UpdateStock(id, quantity*-1)

	if removeFromStockError != nil {
		fmt.Println(removeFromStockError)
		return
	}

	for i, value := range bill.products {
		if value.Id == id {
			bill.products[i].Quantity += quantity
			return
		}
	}

	bill.products = append(bill.products, utils.BillItem{Id: id, Quantity: quantity})
}

func (bill *Bill) RemoveProduct(id string, quantity float64) {
	if !bill.ProductRepo.IsProductValid(id) {
		fmt.Printf("Product with ID %v is not a valid product in the system.", id)

		return
	}

	if quantity <= 0 {
		return
	}

	for i, value := range bill.products {
		if value.Id == id {
			billQuantityIsGreaterThanRemoveQuantity := bill.products[i].Quantity > quantity
			quantityToAddBackToDB := quantity
			if !billQuantityIsGreaterThanRemoveQuantity {
				quantityToAddBackToDB = bill.products[i].Quantity
			}
			_, addbackToStockError := bill.ProductRepo.UpdateStock(id, quantityToAddBackToDB)

			if addbackToStockError != nil {
				fmt.Println(addbackToStockError)
				return
			}

			bill.products[i].Quantity -= quantity
			if bill.products[i].Quantity <= 0 {
				bill.products = append(bill.products[:i], bill.products[i+1:]...)
			}

			return
		}
	}
}

func (bill *Bill) GetProducts() []utils.BillItem {
	return bill.products
}

func (bill *Bill) CalculateTotal() float64 {
	var total float64 = 0

	for _, value := range bill.products {
		product, _ := bill.ProductRepo.GetProductById(value.Id)
		total += value.Quantity * product.UnitPrice
	}

	return total
}

func (bill *Bill) getProductsWithInfosForFormatter() []utils.ProductWithQuantityFromBill {
	products := []utils.ProductWithQuantityFromBill{}

	for _, value := range bill.products {
		product, _ := bill.ProductRepo.GetProductById(value.Id)

		products = append(products, utils.ProductWithQuantityFromBill{
			Product:  *product,
			Quantity: value.Quantity,
		})
	}

	return products
}

func (bill *Bill) FormatBill() bytes.Buffer {

	// Create a BillData DTO
	billData := utils.BillData{
		Products: bill.getProductsWithInfosForFormatter(),
		Subtotal: bill.CalculateTotal(),
		//VAT to be added in the future
		Total: bill.CalculateTotal(),
	}

	formattedBill := bill.Formatter.FormatBill(billData, bill.Printer.GetRowLength())

	return formattedBill
}

func (bill *Bill) PrintBill() {
	formattedBill := bill.FormatBill()

	// Print the formatted bill
	bill.Printer.Print(formattedBill)
}

func (bill *Bill) SaveBill() string {
	data := bill.FormatBill()

	// TODO: problems on saved file if using the printer formatter
	fileName := "bill_" + uuid.NewString() + ".txt"

	error := os.WriteFile(bill.BillsDir+"/"+fileName, data.Bytes(), 0644)

	if error != nil {
		fmt.Println("Error", error)
		panic(error)
	}

	return fileName
}
