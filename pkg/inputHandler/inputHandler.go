package inputHandler

import (
	"biller/pkg/utils"
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
)

type InputHandler struct {
	reader    *bufio.Reader
	validator utils.InputValidatorInterface
}

func NewInputHandler(
	reader *bufio.Reader,
	validator utils.InputValidatorInterface,
) *InputHandler {
	return &InputHandler{
		reader:    reader,
		validator: validator,
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

func (handler *InputHandler) GetValidIntFromInput(prompt string, options utils.GetValidNumberFromInputOptions) int {
	value, _ := handler.getInput(prompt)

	if !handler.validator.ValidateInt(value) || (options.ShouldBePositive && !handler.validator.ValidatePositive(value)) {
		return handler.GetValidIntFromInput(prompt, options)
	}
	intValue, _ := strconv.ParseInt(value, 10, 0)

	return int(intValue)
}

func (handler *InputHandler) getValidFloatFromInput(prompt string, options utils.GetValidNumberFromInputOptions) float64 {
	value, _ := handler.getInput(prompt)

	if !handler.validator.ValidateFloat(value) || (options.ShouldBePositive && !handler.validator.ValidatePositive(value)) {
		return handler.getValidFloatFromInput(prompt, options)
	}
	floatValue, _ := strconv.ParseFloat(value, 64)

	return floatValue
}

func (handler *InputHandler) GetTableName() string {
	tableName, error := handler.getInput("Please, type the table name: ")

	if error != nil {
		fmt.Println("Error:", error)
		return handler.GetTableName()
	}

	return tableName
}

func (handler *InputHandler) getProductItem(products []utils.Product, action string) utils.Product {
	var promptVariant string

	if action == "add" {
		promptVariant = "add to"
	} else {
		promptVariant = "remove from"
	}

	// TODO: show only available product to remove for "remove" action (products that are already in cart)

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

func (handler *InputHandler) GetTip() float64 {
	return handler.getValidFloatFromInput("Add the tip, please: ", utils.GetValidNumberFromInputOptions{ShouldBePositive: true})
}

func (handler *InputHandler) getBillItemQuantity(productName string, action string) int {

	quantity := handler.GetValidIntFromInput(
		fmt.Sprintf("Please provide the quantity of %v you want to %v: ", productName, action),
		// TODO: maibe an interface down here
		utils.GetValidNumberFromInputOptions{ShouldBePositive: true})

	return quantity
}

func (handler *InputHandler) GetBillItem(products []utils.Product, action string) (string, int) {
	product := handler.getProductItem(products, action)
	quantity := handler.getBillItemQuantity(product.Name, action)

	return product.Id, quantity
}
