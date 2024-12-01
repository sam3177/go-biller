package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"

	"biller/pkg/bill"
	"biller/pkg/utils"
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

	quantity := getValidIntFromInput(reader, fmt.Sprintf("Please provide the quantity of %v you want to %v: ", utils.ProductsCatalog[i].Name, action))

	return utils.ProductsCatalog[i].Id, quantity
}

func getValidIntFromInput(reader *bufio.Reader, prompt string) int {
	value, _ := getInput(reader, prompt)

	//check if quantity is a int
	intValue, error := strconv.ParseInt(value, 10, 0)
	if error != nil {
		fmt.Println("Error:", error)

		return getValidIntFromInput(reader, prompt)
	}

	return int(intValue)
}

func getValidFloatFromInput(reader *bufio.Reader, prompt string) float64 {
	value, _ := getInput(reader, prompt)

	//check if quantity is a float64
	floatValue, error := strconv.ParseFloat(value, 64)
	if error != nil {
		fmt.Println("Error:", error)

		return getValidFloatFromInput(reader, prompt)
	}

	return floatValue
}

func InitializeBillWithNumber() *bill.Bill {
	reader := bufio.NewReader(os.Stdin)

	tableName, _ := getInput(reader, "Please, type the table name:")

	bill := bill.NewBill(tableName)

	return bill
}

func HandleActionsOnBill(bill *bill.Bill) {
	reader := bufio.NewReader(os.Stdin)

	promptItems := []string{utils.BILL_ACTIONS["addProduct"], utils.BILL_ACTIONS["addTip"], utils.BILL_ACTIONS["printBill"], utils.BILL_ACTIONS["saveAndExit"], utils.BILL_ACTIONS["exit"]}
	promptItemsWithRemoveOption := append(promptItems[0:1], append([]string{utils.BILL_ACTIONS["removeProduct"]}, promptItems[1:]...)...)

	var prompt promptui.Select

	for {
		fmt.Println(utils.ProductsCatalog)

		var items []string

		if len(bill.Products) > 0 {
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
		case utils.BILL_ACTIONS["addProduct"]:
			name, quantity := getBillItemFromInput(reader, "add")
			bill.AddProduct(name, quantity)
			fmt.Println(bill.Products)

		case utils.BILL_ACTIONS["removeProduct"]:
			name, quantity := getBillItemFromInput(reader, "remove")
			bill.RemoveProduct(name, quantity)
			fmt.Println(bill.Products)

		case utils.BILL_ACTIONS["addTip"]:
			tip := getValidFloatFromInput(reader, "Add the tip, please: ")
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
}
