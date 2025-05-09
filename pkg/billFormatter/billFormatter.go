package billFormatter

import (
	"biller/pkg/utils"
	"fmt"
	"strings"
)

type BillFormatter struct{}

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

func (formatter *BillFormatter) formatVATAmount(subtotal float64, rowLength int) string {
	VATRow := ""
	VATRow += formatter.makeFooterLine("VAT", subtotal, rowLength)

	return VATRow
}

func (formatter *BillFormatter) formatTotal(total float64, rowLength int) string {
	totalRow := ""
	totalRow += formatter.makeFooterLine("Total", total, rowLength)

	return totalRow
}

func (formatter *BillFormatter) formatBillProduct(billProduct utils.ProductWithQuantityFromBill, rowLength int) string {
	formattedProduct := ""
	formattedProduct += fmt.Sprintf(billProduct.Name)

	formattedProduct += fmt.Sprintf("%*v \n",
		rowLength-len(billProduct.Name),
		billProduct.VATCategory,
	)

	formattedQuantityTimesUnitPrice := fmt.Sprintf("%*s",
		rowLength/2,
		fmt.Sprintf("%v %s X %0.2f", billProduct.Quantity, strings.ToUpper(string(billProduct.UnitType)), billProduct.UnitPrice),
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

func (formatter *BillFormatter) formatBillDate(rowLength int, createdAt string) string {
	if createdAt == "" {
		return ""
	}

	dateSection := "\n"
	dateSection += formatter.makeLineSeparator(rowLength)
	dateSection += createdAt
	dateSection += "\n"

	return dateSection
}
