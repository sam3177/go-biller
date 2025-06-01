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
	"flag"
	"fmt"
	"os"
)

func GetProperPrinterAndFormatter() (utils.PrinterInterface, utils.BillFormatterInterface) {
	// Define command line flags
	printerType := flag.String("printer", "terminalPrinter", "Type of printer to use (terminalPrinter or epsonPrinter)")
	flag.Parse()

	switch *printerType {
	case "epsonPrinter":
		printer := printer.NewEpsonPrinter("EPSON_TM_T20III")
		formatter := billFormatter.NewBillEpsonPrinterFormatter(
			qrCodeGenerator.NewQRCodeGenerator(),
		)
		return printer, formatter

	case "terminalPrinter":
		printer := printer.NewTerminalPrinter(50)
		formatter := billFormatter.NewBillTerminalFormatter()
		return printer, formatter

	default:
		fmt.Printf("Warning: Unknown printer type '%s'. Using terminal printer as default.\n", *printerType)
		printer := printer.NewTerminalPrinter(50)
		formatter := billFormatter.NewBillTerminalFormatter()
		return printer, formatter
	}
}

func main() {
	productsJSONHandler := productsJsonStorageHandler.NewProductsJSONStorageHandler(utils.GetAbsolutePath("./data/products.json"))
	billsJSONHandler := billsJsonStorageHandler.NewBillsJSONStorageHandler(utils.GetAbsolutePath("./data/bills.json"))

	error := productsJSONHandler.SeedJSONFile(productRepository.ProductsSeed)

	if error != nil {
		fmt.Println(error)
	}

	printer, formatter := GetProperPrinterAndFormatter()

	productRepo := productRepository.NewLocalProductRepository(productsJSONHandler)
	billRepo := billRepository.NewLocalBillRepository(billsJSONHandler)

	bill := bill.NewBillingHandler(
		productRepo,
		billRepo,
		printer,
		formatter,
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
