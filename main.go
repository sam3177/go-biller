package main

import (
	"biller/pkg/cli"
	"biller/pkg/productRepository"
)

func main() {
	productRepo := productRepository.NewProductRepository()

	bill := cli.InitializeBill(productRepo)

	cli.HandleActionsOnBill(bill, productRepo)
}
