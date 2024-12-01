package cli

import (
	"bufio"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"

	"biller/pkg/bill"
	"biller/pkg/inputHandler"
	"biller/pkg/productRepository"
	"biller/pkg/utils"
)

func getBillItemFromInput(reader *bufio.Reader, action string) (string, int) {
	var promptVariant string
	if action == "add" {
		promptVariant = "add to"
	} else {
		promptVariant = "remove from"
	}

	prompt := promptui.Select{
		Label: fmt.Sprintf("Please type in the product you want to %v the bill: ", promptVariant),
		Items: func() []string {
			var names []string
			for _, product := range utils.ProductsCatalog {
				names = append(names, product.Name)
			}

			return names
		}(),
	}

	i, _, _ := prompt.Run()

	quantity := inputHandler.GetValidIntFromInput(
		reader,
		fmt.Sprintf("Please provide the quantity of %v you want to %v: ", utils.ProductsCatalog[i].Name, action),
		utils.GetValidNumberFromInputOptions{ShouldBePositive: true})

	return utils.ProductsCatalog[i].Id, quantity
}

func InitializeBill() *bill.Bill {
	reader := bufio.NewReader(os.Stdin)

	tableName, _ := inputHandler.GetInput(reader, "Please, type the table name: ")

	// TODO: refactor here
	productRepo := productRepository.NewProductRepository(utils.ProductsCatalog)

	bill := bill.NewBill(tableName, productRepo)

	return bill
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

func HandleActionsOnBill(bill *bill.Bill) {
	reader := bufio.NewReader(os.Stdin)

	promptItems := []string{utils.BILL_ACTIONS["addProduct"], utils.BILL_ACTIONS["addTip"], utils.BILL_ACTIONS["printBill"], utils.BILL_ACTIONS["saveAndExit"], utils.BILL_ACTIONS["exit"]}

	for {
		action, error := selectAction(promptItems, len(bill.Products) > 0)

		if error != nil {
			fmt.Printf("Prompt failed %v\n", error)
			return
		}

		executeAction(bill, reader, action)
	}
}

func executeAction(bill *bill.Bill, reader *bufio.Reader, action string) {
	switch action {
	case utils.BILL_ACTIONS["addProduct"]:
		name, quantity := getBillItemFromInput(reader, "add")
		bill.AddProduct(name, quantity)
		fmt.Println(bill.Products)

	case utils.BILL_ACTIONS["removeProduct"]:
		name, quantity := getBillItemFromInput(reader, "remove")
		bill.RemoveProduct(name, quantity)
		fmt.Println(bill.Products)

	case utils.BILL_ACTIONS["addTip"]:
		tip := inputHandler.GetValidFloatFromInput(reader, "Add the tip, please: ", utils.GetValidNumberFromInputOptions{ShouldBePositive: true})
		bill.SetTip(tip)
		fmt.Println(bill)

	case utils.BILL_ACTIONS["printBill"]:
		bill.PrintBill()

	case utils.BILL_ACTIONS["saveAndExit"]:
		fileName := bill.SaveBill()
		utils.OpenFileInVsCode(utils.BILLS_DIR + "/" + fileName)
		os.Exit(0)

	case utils.BILL_ACTIONS["exit"]:
		os.Exit(0)
	}
}
