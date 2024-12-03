package cli

import (
	"bufio"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"

	"biller/pkg/bill"
	"biller/pkg/inputHandler"
	"biller/pkg/utils"
)

// refactor here
func getBillItemFromInput(reader *bufio.Reader, productRepo utils.ProductRepositoryInterface, action string) (string, int) {
	var promptVariant string
	if action == "add" {
		promptVariant = "add to"
	} else {
		promptVariant = "remove from"
	}

	products := productRepo.GetProducts()

	prompt := promptui.Select{
		Label: fmt.Sprintf("Please type in the product you want to %v the bill: ", promptVariant),
		Items: func() []string {
			var names []string
			for _, product := range products {
				names = append(names, product.Name)
			}

			return names
		}(),
	}

	i, _, _ := prompt.Run()

	quantity := inputHandler.GetValidIntFromInput(
		reader,
		fmt.Sprintf("Please provide the quantity of %v you want to %v: ", products[i].Name, action),
		utils.GetValidNumberFromInputOptions{ShouldBePositive: true})

	return products[i].Id, quantity
}

func selectAction(actions []string, hasProducts bool) (string, error) {
	if hasProducts {
		actions = append(actions[0:1], append([]string{utils.BILL_ACTIONS["removeProduct"]}, actions[1:]...)...)
	}

	prompt := promptui.Select{
		Label: "Select an Option",
		Items: actions,
	}

	_, action, error := prompt.Run()

	return action, error
}

func HandleActionsOnBill(bill *bill.Bill, productRepo utils.ProductRepositoryInterface) {
	reader := bufio.NewReader(os.Stdin)

	promptItems := []string{utils.BILL_ACTIONS["addProduct"], utils.BILL_ACTIONS["addTip"], utils.BILL_ACTIONS["printBill"], utils.BILL_ACTIONS["saveAndExit"], utils.BILL_ACTIONS["exit"]}

	for {
		action, error := selectAction(promptItems, len(bill.GetProducts()) > 0)

		if error != nil {
			fmt.Printf("Prompt failed %v\n", error)
			return
		}

		executeAction(bill, productRepo, reader, action)
	}
}

func executeAction(bill *bill.Bill, productRepo utils.ProductRepositoryInterface, reader *bufio.Reader, action string) {
	switch action {
	case utils.BILL_ACTIONS["addProduct"]:
		name, quantity := getBillItemFromInput(reader, productRepo, "add")
		bill.AddProduct(name, quantity)
		fmt.Println(bill.GetProducts())

	case utils.BILL_ACTIONS["removeProduct"]:
		name, quantity := getBillItemFromInput(reader, productRepo, "remove")
		bill.RemoveProduct(name, quantity)
		fmt.Println(bill.GetProducts())

	case utils.BILL_ACTIONS["addTip"]:
		tip := inputHandler.GetValidFloatFromInput(reader, "Add the tip, please: ", utils.GetValidNumberFromInputOptions{ShouldBePositive: true})
		bill.SetTip(tip)
		fmt.Println(bill)

	case utils.BILL_ACTIONS["printBill"]:
		bill.PrintBill()

	case utils.BILL_ACTIONS["saveAndExit"]:
		fileName := bill.SaveBill()
		utils.OpenFileInVsCode(bill.BillsDir + "/" + fileName)
		os.Exit(0)

	case utils.BILL_ACTIONS["exit"]:
		os.Exit(0)
	}
}
