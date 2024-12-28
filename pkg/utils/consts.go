package utils

var BILL_ROW_LENGTH = 46

var BILL_ACTIONS = map[BillAction]string{
	AddProduct:    "Add product",
	RemoveProduct: "Remove product",
	PrintBill:     "Print bill",
	SaveAndExit:   "Save the bill and exit",
	Exit:          "Exit without saving",
}
