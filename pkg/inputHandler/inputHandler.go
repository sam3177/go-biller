package inputHandler

import (
	"biller/pkg/bill"
	"biller/pkg/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
)

type InputHandler struct {
	reader      *bufio.Reader
	productRepo utils.ProductRepositoryInterface
	bill        *bill.Bill
}

func NewInputHandler(
	reader *bufio.Reader,
	productRepo utils.ProductRepositoryInterface,
	bill *bill.Bill,
) *InputHandler {
	return &InputHandler{
		reader:      reader,
		productRepo: productRepo,
		bill:        bill,
	}
}

func (handler *InputHandler) getInput(prompt string) (string, error) {
	fmt.Print(prompt)
	value, error := handler.reader.ReadString('\n')

	if error != nil {
		fmt.Println("Error:", error)
	}

	return strings.TrimSpace(value), error
}

func (handler *InputHandler) getValidIntFromInput(prompt string, options utils.GetValidNumberFromInputOptions) int {
	value, _ := handler.getInput(prompt)

	intValue, error := strconv.ParseInt(value, 10, 0)
	if error != nil {
		fmt.Println("Error:", error)
		return handler.getValidIntFromInput(prompt, options)
	} else if options.ShouldBePositive && intValue <= 0 {
		fmt.Println("Error: You must enter a positive integer number")
		return handler.getValidIntFromInput(prompt, options)
	}

	return int(intValue)
}

func (handler *InputHandler) getValidFloatFromInput(prompt string, options utils.GetValidNumberFromInputOptions) float64 {
	value, _ := handler.getInput(prompt)

	floatValue, error := strconv.ParseFloat(value, 64)
	if error != nil {
		fmt.Println("Error:", error)
		return handler.getValidFloatFromInput(prompt, options)
	} else if options.ShouldBePositive && floatValue <= 0 {
		fmt.Println("Error: You must enter a positive decimal number")
		return handler.getValidFloatFromInput(prompt, options)
	}

	return floatValue
}

func (handler *InputHandler) getTableName() string {
	tableName, error := handler.getInput("Please, type the table name: ")

	if error != nil {
		fmt.Println("Error:", error)
		return handler.getTableName()
	}

	return tableName
}

func (handler *InputHandler) getBillItem(action string) (string, int) {
	product := handler.getProductItem(action)
	quantity := handler.getBillItemQuantity(product.Name, action)

	return product.Id, quantity
}

func (handler *InputHandler) getProductItem(action string) utils.Product {
	var promptVariant string

	if action == "add" {
		promptVariant = "add to"
	} else {
		promptVariant = "remove from"
	}

	products := handler.productRepo.GetProducts()

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

	return products[i]
}

func (handler *InputHandler) getBillItemQuantity(productName string, action string) int {

	quantity := handler.getValidIntFromInput(
		fmt.Sprintf("Please provide the quantity of %v you want to %v: ", productName, action),
		// TODO: maibe an interface down here
		utils.GetValidNumberFromInputOptions{ShouldBePositive: true})

	return quantity
}

func (handler *InputHandler) selectAction(actions []string, hasProducts bool) (string, error) {
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

func (handler *InputHandler) HandleActions() {
	tableName := handler.getTableName()

	handler.bill.SetTableName(tableName)

	promptItems := []string{utils.BILL_ACTIONS["addProduct"], utils.BILL_ACTIONS["addTip"], utils.BILL_ACTIONS["printBill"], utils.BILL_ACTIONS["saveAndExit"], utils.BILL_ACTIONS["exit"]}

	for {
		action, error := handler.selectAction(promptItems, len(handler.bill.GetProducts()) > 0)

		if error != nil {
			fmt.Printf("Prompt failed %v\n", error)
			return
		}

		handler.executeAction(action)
	}
}

func (handler *InputHandler) executeAction(action string) {
	switch action {
	case utils.BILL_ACTIONS["addProduct"]:
		name, quantity := handler.getBillItem("add")
		handler.bill.AddProduct(name, quantity)
		fmt.Println(handler.bill.GetProducts())

	case utils.BILL_ACTIONS["removeProduct"]:
		name, quantity := handler.getBillItem("remove")
		handler.bill.RemoveProduct(name, quantity)
		fmt.Println(handler.bill.GetProducts())

	case utils.BILL_ACTIONS["addTip"]:
		tip := handler.getValidFloatFromInput("Add the tip, please: ", utils.GetValidNumberFromInputOptions{ShouldBePositive: true})
		handler.bill.SetTip(tip)
		fmt.Println(handler.bill)

	case utils.BILL_ACTIONS["printBill"]:
		handler.bill.PrintBill()

	case utils.BILL_ACTIONS["saveAndExit"]:
		fileName := handler.bill.SaveBill()
		utils.OpenFileInVsCode(handler.bill.BillsDir + "/" + fileName)
		os.Exit(0)

	case utils.BILL_ACTIONS["exit"]:
		os.Exit(0)
	}
}
