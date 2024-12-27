package actionsMenuHandler

import (
	"biller/pkg/bill"
	"biller/pkg/inputHandler"
	"biller/pkg/utils"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

type ActionsMenuHandler struct {
	billingHandler *bill.BillingHandler
	inputHandler   utils.InputHandlerInterface
}

func NewActionMenuHandler(billingHandler *bill.BillingHandler, inputHandler *inputHandler.InputHandler) *ActionsMenuHandler {
	return &ActionsMenuHandler{
		billingHandler: billingHandler,
		inputHandler:   inputHandler,
	}
}

func (menuHandler *ActionsMenuHandler) selectAction(actions []string, hasProducts bool) (string, error) {
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

func (menuHandler *ActionsMenuHandler) HandleActions() {

	promptItems := []string{
		utils.BILL_ACTIONS["addProduct"],
		utils.BILL_ACTIONS["printBill"],
		utils.BILL_ACTIONS["saveAndExit"],
		utils.BILL_ACTIONS["exit"],
	}

	for {
		action, error := menuHandler.selectAction(promptItems, len(menuHandler.billingHandler.GetProducts()) > 0)

		if error != nil {
			fmt.Printf("Prompt failed %v\n", error)
			return
		}

		menuHandler.executeAction(action)
	}
}

func (menuHandler *ActionsMenuHandler) executeAction(action string) {
	products := menuHandler.billingHandler.ProductRepo.GetProducts()

	switch action {
	case utils.BILL_ACTIONS["addProduct"]:
		name, quantity := menuHandler.inputHandler.GetBillItem(products, "add")
		menuHandler.billingHandler.AddProduct(name, quantity)
		fmt.Println(menuHandler.billingHandler.GetProducts())

	case utils.BILL_ACTIONS["removeProduct"]:
		name, quantity := menuHandler.inputHandler.GetBillItem(products, "remove")
		menuHandler.billingHandler.RemoveProduct(name, quantity)
		fmt.Println(menuHandler.billingHandler.GetProducts())

	case utils.BILL_ACTIONS["printBill"]:
		menuHandler.billingHandler.PrintBill()

	case utils.BILL_ACTIONS["saveAndExit"]:
		fileName := menuHandler.billingHandler.SaveBill()
		utils.OpenFileInVsCode(menuHandler.billingHandler.BillsDir + "/" + fileName)
		os.Exit(0)

	case utils.BILL_ACTIONS["exit"]:
		os.Exit(0)
	}
}
