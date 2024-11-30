package main

import (
	"fmt"
	"strings"
)

type Product struct {
	id        string
	name      string
	unitPrice float64
}

type BillProduct struct {
	name     string
	quantity int
}

type Bill struct {
	tableName string
	products  []BillProduct
	tip       float64
}

//create a Bill

func NewBill(tableName string) *Bill {
	return &Bill{
		tableName: tableName,
		products:  []BillProduct{},
		tip:       0,
	}
}

// methods:

// addProduct
func (bill *Bill) AddProduct(name string, quantity int) {
	for i, value := range bill.products {
		if value.name == name {
			bill.products[i].quantity += quantity
			return
		}
	}
	bill.products = append(bill.products, BillProduct{name: name, quantity: quantity})
}

// removeProduct
func (bill *Bill) RemoveProduct(name string, quantity int) {
	for i, value := range bill.products {
		if value.name == name {
			bill.products[i].quantity -= quantity
			if bill.products[i].quantity <= 0 {
				bill.products = append(bill.products[:i], bill.products[i+1:]...)
			}
			return
		}
	}
}

// setTip
func (bill *Bill) SetTip(tip float64) {
	bill.tip = tip
}

// makeTotal
func (bill *Bill) calculateTotal() float64 {
	var total float64 = 0

	for _, value := range bill.products {
		total += float64(value.quantity) * productsCatalog[value.name]
	}

	return total
}

// formatBill // dots for spacing on format fn (extra)
func (bill *Bill) formatBill() string {
	makeFooterLine := func(name string, amount float64) string {
		newLine := "\n" + name

		formattedAmount := fmt.Sprintf("%0.2f", amount)
		newLine += fmt.Sprintf("%*v \n",
			BILL_ROW_LENGTH-len(name),
			formattedAmount,
		)

		return newLine
	}

	billTitle := "----Bill----"
	dottedLine := strings.Repeat("-", BILL_ROW_LENGTH) + "\n"
	formattedBill := fmt.Sprintf("%*s \n", (BILL_ROW_LENGTH+len(billTitle))/2, billTitle)

	formattedBill += fmt.Sprintf("Table %v \n", bill.tableName)
	formattedBill += dottedLine

	fmt.Println()
	for _, value := range bill.products {
		formattedBill += fmt.Sprintf(value.name + "\n")

		formattedQuantityTimesUnitPrice := fmt.Sprintf("%*s",
			BILL_ROW_LENGTH/2,
			fmt.Sprintf("%v X %0.2f", value.quantity, productsCatalog[value.name]),
		)

		formattedBill += formattedQuantityTimesUnitPrice

		totalCost := fmt.Sprintf("%0.2f", float64(value.quantity)*productsCatalog[value.name])
		formattedBill += fmt.Sprintf("%*v \n",
			BILL_ROW_LENGTH/2,
			totalCost,
		)

	}

	formattedBill += dottedLine
	formattedBill += makeFooterLine("Subtotal", bill.calculateTotal())
	formattedBill += makeFooterLine("Tip", bill.tip)
	formattedBill += dottedLine
	formattedBill += makeFooterLine("Total", bill.calculateTotal()+bill.tip)

	return formattedBill
}

// printBill
func (bill *Bill) PrintBill() {
	fmt.Print(bill.formatBill())
}
