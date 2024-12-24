package inputHandler

import (
	"biller/pkg/utils"
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
)

type InputHandler struct {
	reader    utils.InputReaderInterface
	validator utils.InputValidatorInterface
}

func NewInputHandler(
	reader utils.InputReaderInterface,
	validator utils.InputValidatorInterface,
) *InputHandler {
	return &InputHandler{
		reader:    reader,
		validator: validator,
	}
}

func (handler *InputHandler) GetValidIntFromInput(prompt string, options utils.GetValidNumberFromInputOptions) int {
	value, _ := handler.reader.GetInput(prompt)

	if !handler.validator.ValidateInt(value) || (options.ShouldBePositive && !handler.validator.ValidatePositive(value)) {
		return handler.GetValidIntFromInput(prompt, options)
	}
	intValue, _ := strconv.ParseInt(value, 10, 0)

	return int(intValue)
}

func (handler *InputHandler) GetValidFloatFromInput(prompt string, options utils.GetValidNumberFromInputOptions) float64 {
	value, _ := handler.reader.GetInput(prompt)

	if !handler.validator.ValidateFloat(value) || (options.ShouldBePositive && !handler.validator.ValidatePositive(value)) {
		return handler.GetValidFloatFromInput(prompt, options)
	}
	floatValue, _ := strconv.ParseFloat(value, 64)

	return floatValue
}

func (handler *InputHandler) GetTableName() string {
	tableName, error := handler.reader.GetInput("Please, type the table name: ")

	if error != nil {
		fmt.Println("Error:", error)
	}

	if !handler.validator.ValidateMinLength(tableName, 1) {
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
	return handler.GetValidFloatFromInput("Add the tip, please: ", utils.GetValidNumberFromInputOptions{ShouldBePositive: true})
}

func (handler *InputHandler) getBillItemQuantity(productName string, action string) float64 {
	return handler.GetValidFloatFromInput(
		fmt.Sprintf("Please provide the quantity of %v you want to %v: ", productName, action),
		// TODO: maibe an interface down here
		utils.GetValidNumberFromInputOptions{ShouldBePositive: true})
}

func (handler *InputHandler) GetBillItem(products []utils.Product, action string) (string, float64) {
	product := handler.getProductItem(products, action)
	quantity := handler.getBillItemQuantity(product.Name, action)

	return product.Id, quantity
}
