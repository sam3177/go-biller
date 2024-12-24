package billFormatter

import (
	"biller/pkg/utils"
	"bytes"
	"fmt"
)

type BillTerminalFormatter struct {
	BillFormatter
}

func NewBillTerminalFormatter() *BillTerminalFormatter {
	return &BillTerminalFormatter{}
}

func (formatter *BillTerminalFormatter) FormatBill(billData utils.BillData, rowLength int) bytes.Buffer {
	var buffer bytes.Buffer

	billTitle := formatter.getBillTitle()
	buffer.WriteString(fmt.Sprintf("%*s", (rowLength+len(billTitle))/2, billTitle))

	buffer.WriteString(fmt.Sprintf("Table name: %v \n", billData.TableName))

	buffer.WriteString(formatter.makeDottedLine(rowLength))

	for _, product := range billData.Products {
		buffer.WriteString(formatter.formatBillProduct(product, rowLength))
	}

	buffer.WriteString(formatter.makeDottedLine(rowLength))

	buffer.WriteString(formatter.formatSubtotal(billData.Subtotal, rowLength))

	buffer.WriteString(formatter.formatTotal(billData.Total, rowLength))

	return buffer
}
