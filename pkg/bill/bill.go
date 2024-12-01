package bill

import (
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"

	"biller/pkg/productRepository"
	"biller/pkg/utils"
)

type BillItem struct {
	Id       string
	Quantity int
}

type Bill struct {
	TableName   string
	Products    []BillItem
	Tip         float64
	ProductRepo productRepository.ProductRepositoryInterface
}

func NewBill(tableName string, productRepo *productRepository.ProductRepository) *Bill {
	return &Bill{
		TableName:   tableName,
		Products:    []BillItem{},
		Tip:         0,
		ProductRepo: productRepo,
	}
}

// methods:

// addProduct
func (bill *Bill) AddProduct(id string, quantity int) {
	if !bill.ProductRepo.IsProductValid(id) {
		fmt.Printf("Product with ID %v is not a valid product in the system.", id)

		return
	}

	for i, value := range bill.Products {
		if value.Id == id {
			bill.Products[i].Quantity += quantity
			return
		}
	}

	bill.Products = append(bill.Products, BillItem{Id: id, Quantity: quantity})
}

// removeProduct
func (bill *Bill) RemoveProduct(id string, quantity int) {
	if !bill.ProductRepo.IsProductValid(id) {
		fmt.Printf("Product with ID %v is not a valid product in the system.", id)

		return
	}

	for i, value := range bill.Products {
		if value.Id == id {
			bill.Products[i].Quantity -= quantity
			if bill.Products[i].Quantity <= 0 {
				bill.Products = append(bill.Products[:i], bill.Products[i+1:]...)
			}
			return
		}
	}
}

// setTip
func (bill *Bill) SetTip(tip float64) {
	bill.Tip = tip
}

// makeTotal
func (bill *Bill) calculateTotal() float64 {
	var total float64 = 0

	for _, value := range bill.Products {
		product, _ := bill.ProductRepo.GetProductById(value.Id)
		total += float64(value.Quantity) * product.UnitPrice
	}

	return total
}

// formatBill // dots for spacing on format fn (extra)
func (bill *Bill) formatBill() string {
	makeFooterLine := func(name string, amount float64) string {
		newLine := "\n" + name

		formattedAmount := fmt.Sprintf("%0.2f", amount)
		newLine += fmt.Sprintf("%*v \n",
			utils.BILL_ROW_LENGTH-len(name),
			formattedAmount,
		)

		return newLine
	}

	billTitle := "----Bill----"
	dottedLine := strings.Repeat("-", utils.BILL_ROW_LENGTH) + "\n"
	formattedBill := fmt.Sprintf("%*s \n", (utils.BILL_ROW_LENGTH+len(billTitle))/2, billTitle)

	formattedBill += fmt.Sprintf("Table name: %v \n", bill.TableName)
	formattedBill += dottedLine

	for _, value := range bill.Products {
		product, _ := bill.ProductRepo.GetProductById(value.Id)

		formattedBill += fmt.Sprintf(product.Name + "\n")

		formattedQuantityTimesUnitPrice := fmt.Sprintf("%*s",
			utils.BILL_ROW_LENGTH/2,
			fmt.Sprintf("%v X %0.2f", value.Quantity, product.UnitPrice),
		)

		formattedBill += formattedQuantityTimesUnitPrice

		totalCost := fmt.Sprintf("%0.2f", float64(value.Quantity)*product.UnitPrice)
		formattedBill += fmt.Sprintf("%*v \n",
			utils.BILL_ROW_LENGTH/2,
			totalCost,
		)

	}

	formattedBill += dottedLine
	formattedBill += makeFooterLine("Subtotal", bill.calculateTotal())
	formattedBill += makeFooterLine("Tip", bill.Tip)
	formattedBill += dottedLine
	formattedBill += makeFooterLine("Total", bill.calculateTotal()+bill.Tip)
	formattedBill += "\n"

	return formattedBill
}

// printBill
func (bill *Bill) PrintBill() {
	fmt.Print(bill.formatBill())
}

func (bill *Bill) SaveBill() string {
	data := []byte(bill.formatBill())

	fileName := "table_" + bill.TableName + "_" + uuid.NewString() + ".txt"

	error := os.WriteFile(utils.BILLS_DIR+"/"+fileName, data, 0644)

	if error != nil {
		fmt.Println("Error", error)
		panic(error)
	}

	return fileName
}
