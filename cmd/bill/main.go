package main

import (
	"biller/pkg/bill"
	"biller/pkg/cli"
	"biller/pkg/inputHandler"
	"biller/pkg/productRepository"
	"biller/pkg/utils"
)

func main() {
	productRepo := productRepository.NewLocalProductRepository(productRepository.ProductsCatalog)

	tableName := inputHandler.GetTableName()

	bill := bill.NewBill(
		tableName,
		productRepo,
		utils.BillConfig{
			BillsDir:      utils.GetBillsDir(),
			BillRowLength: utils.BILL_ROW_LENGTH,
		},
	)

	cli.HandleActionsOnBill(bill, productRepo)
}
