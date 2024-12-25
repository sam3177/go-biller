package main

import (
	"biller/pkg/actionsMenuHandler"
	"biller/pkg/bill"
	"biller/pkg/billFormatter"
	"biller/pkg/inputHandler"
	"biller/pkg/inputReader"
	"biller/pkg/inputValidator"
	"biller/pkg/printer"
	"biller/pkg/productRepository"
	"biller/pkg/productsJsonStorageHandler"
	"biller/pkg/utils"
	"bufio"
	"fmt"
	"os"
)

func main() {
	productsJSONHandler := productsJsonStorageHandler.NewProductsJSONStorageHandler("./data/products.json")

	error := productsJSONHandler.SeedJSONFile(productRepository.ProductsSeed)

	if error != nil {
		fmt.Println(error)
	}

	productRepo := productRepository.NewLocalProductRepository(productsJSONHandler)

	// epsonPrinter := printer.NewEpsonPrinter("EPSON_TM_T20III")
	termimalPrinter := printer.NewTerminalPrinter(50)

	// epsonPrinterFormatter := billFormatter.NewBillEpsonPrinterFormatter()
	terminalPrinterFormatter := billFormatter.NewBillTerminalFormatter()

	bill := bill.NewBill(
		productRepo,
		termimalPrinter,
		terminalPrinterFormatter,
		utils.GetBillsDir(),
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
