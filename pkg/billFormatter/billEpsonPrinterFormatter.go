package billFormatter

import (
	"biller/pkg/utils"
	"bytes"
	"fmt"
)

type BillEpsonPrinterFormatter struct {
	BillFormatter
	buffer bytes.Buffer
}

func NewBillEpsonPrinterFormatter() *BillEpsonPrinterFormatter {
	return &BillEpsonPrinterFormatter{}
}

func (formatter *BillEpsonPrinterFormatter) FormatBill(billData utils.BillData, rowLength int) bytes.Buffer {
	formatter.buffer.Write([]byte{0x1B, 0x40})             // ESC @ (initialize)
	formatter.buffer.Write([]byte{0x1D, 0x4C, 0x0A, 0x00}) // 10 dots left margin (1 character width)

	formatter.makeBillHeader(billData.TableName, rowLength)

	for _, product := range billData.Products {
		formatter.buffer.WriteString(formatter.formatBillProduct(product, rowLength))
	}

	formatter.makeSeparator(rowLength)

	formatter.buffer.WriteString(formatter.formatSubtotal(billData.Subtotal, rowLength))

	formatter.enableBold()

	formatter.buffer.WriteString(formatter.formatTotal(billData.Total, rowLength))

	formatter.buffer.Write([]byte{0x1B, 0x64, 0x06}) // Feed 6 lines
	formatter.buffer.Write([]byte{0x1D, 0x56, 0x00}) // Cut the paper

	return formatter.buffer
}

func (formatter *BillEpsonPrinterFormatter) makeBillHeader(tableName string, rowLength int) {
	formatter.alignCenter()
	formatter.enableBold()
	formatter.buffer.Write([]byte{0x1D, 0x21, 0x11}) // Double width and height font

	formatter.buffer.WriteString(formatter.getBillTitle())
	formatter.buffer.Write([]byte{0x1B, 0x64, 0x02}) // Feed 2 lines

	formatter.alignLeft()
	formatter.disableBold()
	formatter.buffer.Write([]byte{0x1D, 0x21, 0x00}) // Reset font size

	formatter.buffer.WriteString(fmt.Sprintf("Table name: %v \n", tableName))
	formatter.makeSeparator(rowLength)
}

func (formatter *BillEpsonPrinterFormatter) alignCenter() {
	formatter.buffer.Write([]byte{0x1B, 0x61, 0x01}) //center align
}

func (formatter *BillEpsonPrinterFormatter) alignLeft() {
	formatter.buffer.Write([]byte{0x1B, 0x61, 0x00}) // left align
}

func (formatter *BillEpsonPrinterFormatter) enableBold() {
	formatter.buffer.Write([]byte{0x1B, 0x45, 0x01}) // ESC E 1 - Enable bold
}

func (formatter *BillEpsonPrinterFormatter) disableBold() {
	formatter.buffer.Write([]byte{0x1B, 0x45, 0x00}) // ESC E 0 - Disable bold
}

func (formatter *BillEpsonPrinterFormatter) makeSeparator(rowLength int) {
	formatter.buffer.WriteString(formatter.makeDottedLine(rowLength))
}
