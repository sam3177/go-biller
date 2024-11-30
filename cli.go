package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
)

func getInput(reader *bufio.Reader, prompt string) (string, error) {
	fmt.Print(prompt)
	value, error := reader.ReadString('\n')

	if error != nil {
		fmt.Println("Error:", error)
	}

	return strings.TrimSpace(value), error
}

func getBillItemFromInput(reader *bufio.Reader, action string) (string, int) {
	var promptVariant string
	if action == "add" {
		promptVariant = "add to"
	} else {
		promptVariant = "remove from"
	}

	productName, _ := getInput(reader, fmt.Sprintf("Please type in the product you want to %v the bill: ", promptVariant))
	var quantity int

	validQuantity := false
	for !validQuantity {

		inputQuantity, _ := getInput(reader, fmt.Sprintf("Please provide the quantity of %v you want to %v: ", productName, action))

		//check if quantity is a float64
		intValue, error := strconv.ParseInt(inputQuantity, 10, 0)
		if error != nil {
			fmt.Println("Error:", error)
			continue
		}
		validQuantity = true
		quantity = int(intValue)
	}
	return productName, quantity
}

func initializeBillWithNumber() *Bill {
	reader := bufio.NewReader(os.Stdin)

	tableNumber, _ := getInput(reader, "Please, type the table number:")

	bill := NewBill(tableNumber)
	return bill
}

func handleActionsOnBill(bill *Bill) {
	reader := bufio.NewReader(os.Stdin)

	promptItems := []string{BILL_ACTIONS["addProduct"], BILL_ACTIONS["addTip"], BILL_ACTIONS["printBill"], BILL_ACTIONS["saveAndExit"], BILL_ACTIONS["exit"]}
	promptItemsWithRemoveOption := append(promptItems[0:1], append([]string{BILL_ACTIONS["removeProduct"]}, promptItems[1:]...)...)

	var prompt promptui.Select

	for {
		fmt.Println(productsCatalog)

		var items []string

		if len(bill.products) > 0 {
			items = promptItemsWithRemoveOption
		} else {
			items = promptItems

		}

		prompt = promptui.Select{
			Label: "Select an Option",
			Items: items,
		}

		_, action, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		switch action {
		case BILL_ACTIONS["addProduct"]:
			name, quantity := getBillItemFromInput(reader, "add")
			bill.AddProduct(name, quantity)
			fmt.Println(bill.products)
		case BILL_ACTIONS["removeProduct"]:
			name, quantity := getBillItemFromInput(reader, "remove")
			bill.RemoveProduct(name, quantity)
			fmt.Println(bill.products)
		case BILL_ACTIONS["addTip"]:
			validTip := false
			for !validTip {
				tip, _ := getInput(reader, "Add the tip, please: ")
				floatValue, error := strconv.ParseFloat(tip, 64)
				if error != nil {
					fmt.Println("Error:", error)
					continue
				}
				validTip = true
				bill.SetTip(floatValue)
			}
			fmt.Println(bill)
		case BILL_ACTIONS["printBill"]:
			bill.PrintBill()
		case BILL_ACTIONS["exit"]:
			os.Exit(0)
		}
	}
}
