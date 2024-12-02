package main

import (
	"biller/pkg/cli"
	"biller/pkg/productRepository"
)

func main() {
	productRepo := productRepository.NewLocalProductRepository(productRepository.ProductsCatalog)

	bill := cli.InitializeBill(productRepo)

	cli.HandleActionsOnBill(bill, productRepo)
}
