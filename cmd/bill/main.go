package main

import (
	"biller/pkg/actionsMenuHandler"
	"biller/pkg/bill"
	"biller/pkg/inputHandler"
	"biller/pkg/inputReader"
	"biller/pkg/inputValidator"
	"biller/pkg/printer"
	"biller/pkg/productRepository"
	"biller/pkg/utils"
	"bufio"
	"os"
)

func main() {
	productRepo := productRepository.NewLocalProductRepository(productRepository.MakeProductsCatalog(productRepository.ProductsSlice))
	termimalPrinter := printer.NewTerminalPrinter()

	bill := bill.NewBill(
		productRepo,
		termimalPrinter,
		utils.BillConfig{
			BillsDir:      utils.GetBillsDir(),
			BillRowLength: utils.BILL_ROW_LENGTH,
		},
	)

	inputValidator := inputValidator.NewInputValidator()
	inputReader := inputReader.NewInputReader(bufio.NewReader(os.Stdin))

	inputHandler := inputHandler.NewInputHandler(
		inputReader,
		inputValidator,
	)

	actionsMenu := actionsMenuHandler.NewActionMenuHandler(
		bill, inputHandler,
	)

	actionsMenu.HandleActions()
}
