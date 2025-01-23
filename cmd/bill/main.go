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

func GetProperPrinterAndFormatter() (utils.PrinterInterface, utils.BillFormatterInterface) {
	args := os.Args[1:]
	var printerType string

	if len(args) == 0 {
		printerType = "epsonPrinter"
	} else {
		printerType = args[0]
	}

	// TODO: check for PrinterType
	// if printerType == "" || (printerType != TerminalPrinter && printerType != EpsonPrinter) {
	// 	printerType = TerminalPrinter

	// }

	switch printerType {

	// case EpsonPrinter:
	case "epsonPrinter":
		printer := printer.NewEpsonPrinter("EPSON_TM_T20III")

		formatter := billFormatter.NewBillEpsonPrinterFormatter(
			qrCodeGenerator.NewQRCodeGenerator(),
		)

		return printer, formatter

	// case TerminalPrinter:
	case "terminalPrinter":
		printer := printer.NewTerminalPrinter(50)
		formatter := billFormatter.NewBillTerminalFormatter()

		return printer, formatter

	default:
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
