package billFormatter

import (
	"biller/pkg/utils"
	"bytes"
	"fmt"
)

type BillTerminalFormatter struct {
	BillFormatter
	buffer bytes.Buffer
}

func NewBillTerminalFormatter() *BillTerminalFormatter {
	return &BillTerminalFormatter{}
}

func (formatter *BillTerminalFormatter) FormatBill(billData utils.BillData, rowLength int) bytes.Buffer {
	defer formatter.buffer.Reset()

	billTitle := formatter.getBillTitle()
	formatter.buffer.WriteString(fmt.Sprintf("%*s", (rowLength+len(billTitle))/2, billTitle))

	formatter.buffer.WriteString(formatter.makeLineSeparator(rowLength))

	for _, product := range billData.Products {
		formatter.buffer.WriteString(formatter.formatBillProduct(product, rowLength))
	}

	formatter.buffer.WriteString(formatter.makeLineSeparator(rowLength))

	formatter.buffer.WriteString(formatter.formatSubtotal(billData.Subtotal, rowLength))

	formatter.buffer.WriteString(formatter.formatVATAmount(billData.VATAmount, rowLength))

	formatter.buffer.WriteString(formatter.makeLineSeparator(rowLength))

	formatter.buffer.WriteString(formatter.formatTotal(billData.Total, rowLength))

	return formatter.buffer
}
