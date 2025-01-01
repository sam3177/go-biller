package billFormatter

import (
	"biller/pkg/utils"
	"bytes"
	"fmt"
	"image"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

type BillEpsonPrinterFormatter struct {
	BillFormatter
	buffer bytes.Buffer
}

func NewBillEpsonPrinterFormatter() *BillEpsonPrinterFormatter {
	return &BillEpsonPrinterFormatter{}
}

func (formatter *BillEpsonPrinterFormatter) FormatBill(billData utils.BillData, rowLength int) bytes.Buffer {
	defer formatter.buffer.Reset()

	formatter.buffer.Write([]byte{0x1B, 0x40})             // ESC @ (initialize)
	formatter.buffer.Write([]byte{0x1D, 0x4C, 0x0A, 0x00}) // 10 dots left margin (1 character width)

	formatter.makeBillHeader(rowLength)

	for _, product := range billData.Products {
		formatter.buffer.WriteString(formatter.formatBillProduct(product, rowLength))
	}

	formatter.buffer.WriteString(formatter.makeLineSeparator(rowLength))

	formatter.buffer.WriteString(formatter.formatSubtotal(billData.Subtotal, rowLength))

	formatter.buffer.WriteString(formatter.makeLineSeparator(rowLength))

	formatter.enableBold()

	formatter.buffer.WriteString(formatter.formatTotal(billData.Total, rowLength))
	formatter.alignCenter()

	formatter.GetBarcodeBuffer()

	formatter.buffer.Write([]byte{0x1B, 0x64, 0x06}) // Feed 6 lines
	formatter.buffer.Write([]byte{0x1D, 0x56, 0x00}) // Cut the paper

	return formatter.buffer
}

func (formatter *BillEpsonPrinterFormatter) makeBillHeader(rowLength int) {
	formatter.alignCenter()
	formatter.enableBold()

	formatter.buffer.WriteString(formatter.getBillTitle())
	formatter.buffer.Write([]byte{0x1B, 0x64, 0x02}) // Feed 2 lines

	formatter.alignLeft()
	formatter.disableBold()

	formatter.buffer.WriteString(formatter.makeLineSeparator(rowLength))
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

func (formatter *BillFormatter) CreateBarcode(id string) (barcode.Barcode, error) {

	// Generate a Code128 barcode
	code, err := qr.Encode(id, qr.M, qr.Auto)
	if err != nil {
		fmt.Println("Error generating barcode:", err)

		return nil, err
	}
	code, _ = barcode.Scale(code, 300, 300)

	return code, nil
}

func (formatter *BillEpsonPrinterFormatter) GetBarcodeBuffer() {

	barcode, err := formatter.CreateBarcode("https://www.instagram.com/silviu.rvn/")
	if err != nil {
		fmt.Println("Error generating barcode:", err)
		return
	}
	bytes, err := convertImageToRaster(barcode)

	if err != nil {
		fmt.Println("Error converting image to raster:", err)
		return
	}

	formatter.buffer.Write(bytes)
}

func convertImageToRaster(img image.Image) ([]byte, error) {
	var buffer bytes.Buffer

	// Convert the image to a monochrome raster format
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	widthBytes := (width + 7) / 8 // 1 byte = 8 pixels

	fmt.Println("widthBytes", widthBytes)
	//print hieght

	fmt.Println("height", height)
	//print width
	fmt.Println("width", width)

	// ESC/POS raster command header
	buffer.Write([]byte{0x1D, 0x76, 0x30, 0x00})
	buffer.Write([]byte{byte(widthBytes), byte(widthBytes >> 8)})
	buffer.Write([]byte{byte(height), byte(height >> 8)})

	// Convert image pixels to binary data
	for y := 0; y < height; y++ {
		for x := 0; x < width; x += 8 {
			var byteData byte
			for bit := 0; bit < 8; bit++ {
				if x+bit < width {
					r, g, b, _ := img.At(x+bit, y).RGBA()
					// Convert to grayscale and check if it's closer to black or white
					gray := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
					if gray < 128 {
						byteData |= (1 << (7 - bit))
					}
				}
			}
			buffer.WriteByte(byteData)
		}
	}

	return buffer.Bytes(), nil
}
