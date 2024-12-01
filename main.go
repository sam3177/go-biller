package main

import "biller/pkg/cli"

func main() {
	bill := cli.InitializeBill()
	cli.HandleActionsOnBill(bill)
}
