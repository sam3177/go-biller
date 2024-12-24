package billFormatter

import (
	"biller/pkg/utils"
	"fmt"
	"strings"
)

type BillFormatter struct {
}

func (formatter *BillFormatter) makeFooterLine(name string, amount float64, rowLength int) string {
	newLine := name

	formattedAmount := fmt.Sprintf("%0.2f", amount)
	newLine += fmt.Sprintf("%*v \n",
		rowLength-len(name),
		formattedAmount,
	)

	return newLine
}
func (formatter *BillFormatter) makeLineSeparator(rowLength int) string {
	return strings.Repeat("-", rowLength) + "\n"
}

func (formatter *BillFormatter) formatSubtotal(subtotal float64, rowLength int) string {
	subtotalRow := ""
	subtotalRow += formatter.makeFooterLine("Subtotal", subtotal, rowLength)

	return subtotalRow
}

func (formatter *BillFormatter) formatTotal(total float64, rowLength int) string {
	totalRow := ""
	totalRow += formatter.makeFooterLine("Total", total, rowLength)

	return totalRow
}

func (formatter *BillFormatter) formatBillProduct(billProduct utils.ProductWithQuantityFromBill, rowLength int) string {
	formattedProduct := ""
	formattedProduct += fmt.Sprintf(billProduct.Name + "\n")

	formattedQuantityTimesUnitPrice := fmt.Sprintf("%*s",
		rowLength/2,
		fmt.Sprintf("%v X %0.2f", billProduct.Quantity, billProduct.UnitPrice),
	)

	formattedProduct += formattedQuantityTimesUnitPrice

	totalCost := fmt.Sprintf("%0.2f", float64(billProduct.Quantity)*billProduct.UnitPrice)
	formattedProduct += fmt.Sprintf("%*v \n",
		rowLength/2,
		totalCost,
	)
	return formattedProduct
}

func (formatter *BillFormatter) getBillTitle() string {
	return "----Bill----\n"
}
