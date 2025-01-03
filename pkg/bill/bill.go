package bill

import (
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"

	"biller/pkg/utils"
)

type Bill struct {
	tableName   string
	products    []utils.BillItem
	tip         float64
	ProductRepo utils.ProductRepositoryInterface
	utils.BillConfig
}

func NewBill(
	tableName string,
	productRepo utils.ProductRepositoryInterface,
	config utils.BillConfig,
) *Bill {
	return &Bill{
		tableName:   tableName,
		products:    []utils.BillItem{},
		tip:         0,
		ProductRepo: productRepo,
		BillConfig:  config,
	}
}

func (bill *Bill) AddProduct(id string, quantity int) {
	if !bill.ProductRepo.IsProductValid(id) {
		fmt.Printf("Product with ID %v is not a valid product in the system.", id)

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

func (bill *Bill) RemoveProduct(id string, quantity int) {
	if !bill.ProductRepo.IsProductValid(id) {
		fmt.Printf("Product with ID %v is not a valid product in the system.", id)

		return
	}

	for i, value := range bill.products {
		if value.Id == id {
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

func (bill *Bill) SetTip(tip float64) {
	bill.tip = tip
}

func (bill *Bill) CalculateTotal() float64 {
	var total float64 = 0

	for _, value := range bill.products {
		product, _ := bill.ProductRepo.GetProductById(value.Id)
		total += float64(value.Quantity) * product.UnitPrice
	}

	return total
}

func (bill *Bill) FormatBill() string {
	makeFooterLine := func(name string, amount float64) string {
		newLine := "\n" + name

		formattedAmount := fmt.Sprintf("%0.2f", amount)
		newLine += fmt.Sprintf("%*v \n",
			bill.BillRowLength-len(name),
			formattedAmount,
		)

		return newLine
	}

	billTitle := "----Bill----"
	dottedLine := strings.Repeat("-", bill.BillRowLength) + "\n"
	formattedBill := fmt.Sprintf("%*s \n", (bill.BillRowLength+len(billTitle))/2, billTitle)

	formattedBill += fmt.Sprintf("Table name: %v \n", bill.tableName)
	formattedBill += dottedLine

	for _, value := range bill.products {
		product, _ := bill.ProductRepo.GetProductById(value.Id)

		formattedBill += fmt.Sprintf(product.Name + "\n")

		formattedQuantityTimesUnitPrice := fmt.Sprintf("%*s",
			bill.BillRowLength/2,
			fmt.Sprintf("%v X %0.2f", value.Quantity, product.UnitPrice),
		)

		formattedBill += formattedQuantityTimesUnitPrice

		totalCost := fmt.Sprintf("%0.2f", float64(value.Quantity)*product.UnitPrice)
		formattedBill += fmt.Sprintf("%*v \n",
			bill.BillRowLength/2,
			totalCost,
		)

	}

	formattedBill += dottedLine
	formattedBill += makeFooterLine("Subtotal", bill.CalculateTotal())
	formattedBill += makeFooterLine("Tip", bill.tip)
	formattedBill += dottedLine
	formattedBill += makeFooterLine("Total", bill.CalculateTotal()+bill.tip)
	formattedBill += "\n"

	return formattedBill
}

func (bill *Bill) PrintBill() {
	fmt.Print(bill.FormatBill())
}

func (bill *Bill) SaveBill() string {
	data := []byte(bill.FormatBill())

	fileName := "table_" + bill.tableName + "_" + uuid.NewString() + ".txt"

	error := os.WriteFile(bill.BillsDir+"/"+fileName, data, 0644)

	if error != nil {
		fmt.Println("Error", error)
		panic(error)
	}

	return fileName
}
