package main

import "biller/pkg/cli"

func main() {
	bill := cli.InitializeBillWithNumber()
	cli.HandleActionsOnBill(bill)
}
