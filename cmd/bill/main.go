package main

import (
	"biller/pkg/bill"
	"biller/pkg/inputHandler"
	"biller/pkg/printer"
	"biller/pkg/productRepository"
	"biller/pkg/utils"
	"bufio"
	"os"
)

func main() {
	productRepo := productRepository.NewLocalProductRepository(productRepository.ProductsCatalog)
	termimalPrinter := printer.NewTerminalPrinter()

	bill := bill.NewBill(
		productRepo,
		termimalPrinter,
		utils.BillConfig{
			BillsDir:      utils.GetBillsDir(),
			BillRowLength: utils.BILL_ROW_LENGTH,
		},
	)

	inputHandler := inputHandler.NewInputHandler(
		bufio.NewReader(os.Stdin),
		productRepo,
		bill,
	)

	inputHandler.HandleActions()
}
