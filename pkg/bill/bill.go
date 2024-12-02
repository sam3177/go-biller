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
	tableName   string
	products    []BillItem
	tip         float64
	ProductRepo productRepository.ProductRepositoryInterface
}

func NewBill(tableName string, productRepo *productRepository.ProductRepository) *Bill {
	return &Bill{
		tableName:   tableName,
		products:    []BillItem{},
		tip:         0,
		ProductRepo: productRepo,
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

	bill.products = append(bill.products, BillItem{Id: id, Quantity: quantity})
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

func (bill *Bill) GetProducts() []BillItem {
	return bill.products
}

func (bill *Bill) SetTip(tip float64) {
	bill.tip = tip
}

func (bill *Bill) calculateTotal() float64 {
	var total float64 = 0

	for _, value := range bill.products {
		product, _ := bill.ProductRepo.GetProductById(value.Id)
		total += float64(value.Quantity) * product.UnitPrice
	}

	return total
}

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

	formattedBill += fmt.Sprintf("Table name: %v \n", bill.tableName)
	formattedBill += dottedLine

	for _, value := range bill.products {
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
	formattedBill += makeFooterLine("Tip", bill.tip)
	formattedBill += dottedLine
	formattedBill += makeFooterLine("Total", bill.calculateTotal()+bill.tip)
	formattedBill += "\n"

	return formattedBill
}

func (bill *Bill) PrintBill() {
	fmt.Print(bill.formatBill())
}

func (bill *Bill) SaveBill() string {
	data := []byte(bill.formatBill())

	fileName := "table_" + bill.tableName + "_" + uuid.NewString() + ".txt"

	error := os.WriteFile(utils.BILLS_DIR+"/"+fileName, data, 0644)

	if error != nil {
		fmt.Println("Error", error)
		panic(error)
	}

	return fileName
}
