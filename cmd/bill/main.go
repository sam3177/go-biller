package main

import (
	"biller/pkg/actionsMenuHandler"
	"biller/pkg/bill"
	"biller/pkg/billFormatter"
	"biller/pkg/billRepository"
	"biller/pkg/billsJsonStorageHandler"
	"biller/pkg/inputHandler"
	"biller/pkg/inputReader"
	"biller/pkg/inputValidator"
	"biller/pkg/printer"
	"biller/pkg/productRepository"
	"biller/pkg/productsJsonStorageHandler"
	"biller/pkg/qrCodeGenerator"
	"biller/pkg/utils"
	"bufio"
	"fmt"
	"os"
)

func main() {
	productsJSONHandler := productsJsonStorageHandler.NewProductsJSONStorageHandler(utils.GetAbsolutePath("./data/products.json"))
	billsJSONHandler := billsJsonStorageHandler.NewBillsJSONStorageHandler(utils.GetAbsolutePath("./data/bills.json"))

	error := productsJSONHandler.SeedJSONFile(productRepository.ProductsSeed)

	if error != nil {
		fmt.Println(error)
	}

	productRepo := productRepository.NewLocalProductRepository(productsJSONHandler)
	billRepo := billRepository.NewLocalBillRepository(billsJSONHandler)

	epsonPrinter := printer.NewEpsonPrinter("EPSON_TM_T20III")
	// termimalPrinter := printer.NewTerminalPrinter(50)

	epsonPrinterFormatter := billFormatter.NewBillEpsonPrinterFormatter(
		qrCodeGenerator.NewQRCodeGenerator(),
	)
	// terminalPrinterFormatter := billFormatter.NewBillTerminalFormatter()

	bill := bill.NewBillingHandler(
		productRepo,
		billRepo,
		epsonPrinter,
		epsonPrinterFormatter,
		utils.GetAbsolutePath("./bills"),
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
