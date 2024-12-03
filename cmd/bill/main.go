package main

import (
	"biller/pkg/bill"
	"biller/pkg/cli"
	"biller/pkg/inputHandler"
	"biller/pkg/printer"
	"biller/pkg/productRepository"
	"biller/pkg/utils"
)

func main() {
	productRepo := productRepository.NewLocalProductRepository(productRepository.ProductsCatalog)
	termimalPrinter := printer.NewTerminalPrinter()

	tableName := inputHandler.GetTableName()

	bill := bill.NewBill(
		tableName,
		productRepo,
		termimalPrinter,
		utils.BillConfig{
			BillsDir:      utils.GetBillsDir(),
			BillRowLength: utils.BILL_ROW_LENGTH,
		},
	)

	cli.HandleActionsOnBill(bill, productRepo)
}
