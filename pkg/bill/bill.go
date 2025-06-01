package bill

import (
	"bytes"
	"fmt"
	"time"

	"github.com/google/uuid"

	"biller/pkg/utils"
)

type BillingHandler struct {
	products    []utils.BillProduct
	ProductRepo utils.ProductRepositoryInterface
	BillRepo    utils.BillRepositoryInterface
	Printer     utils.PrinterInterface
	Formatter   utils.BillFormatterInterface
	FileHandler utils.FileHandlerInterface
}

func NewBillingHandler(
	productRepo utils.ProductRepositoryInterface,
	billRepo utils.BillRepositoryInterface,
	printer utils.PrinterInterface,
	formatter utils.BillFormatterInterface,
	fileHandler utils.FileHandlerInterface,
) *BillingHandler {
	return &BillingHandler{
		products:    []utils.BillProduct{},
		ProductRepo: productRepo,
		BillRepo:    billRepo,
		Formatter:   formatter,
		Printer:     printer,
		FileHandler: fileHandler,
	}
}

func (billingHandler *BillingHandler) AddProduct(id string, quantity float64) {
	if !billingHandler.ProductRepo.IsProductValid(id) {
		fmt.Printf("Product with ID %v is not a valid product in the system.", id)
		return
	}
	if quantity <= 0 {
		return
	}

	_, removeFromStockError := billingHandler.ProductRepo.UpdateStock(id, quantity*-1)

	if removeFromStockError != nil {
		fmt.Println(removeFromStockError)
		return
	}

	for i, value := range billingHandler.products {
		if value.Id == id {
			billingHandler.products[i].Quantity += quantity
			return
		}
	}

	billingHandler.products = append(billingHandler.products, utils.BillProduct{Id: id, Quantity: quantity})
}

func (billingHandler *BillingHandler) RemoveProduct(id string, quantity float64) {
	if !billingHandler.ProductRepo.IsProductValid(id) {
		fmt.Printf("Product with ID %v is not a valid product in the system.", id)

		return
	}

	if quantity <= 0 {
		return
	}

	for i, value := range billingHandler.products {
		if value.Id == id {
			billQuantityIsGreaterThanRemoveQuantity := billingHandler.products[i].Quantity > quantity
			quantityToAddBackToDB := quantity
			if !billQuantityIsGreaterThanRemoveQuantity {
				quantityToAddBackToDB = billingHandler.products[i].Quantity
			}
			_, addbackToStockError := billingHandler.ProductRepo.UpdateStock(id, quantityToAddBackToDB)

			if addbackToStockError != nil {
				fmt.Println(addbackToStockError)
				return
			}

			billingHandler.products[i].Quantity -= quantity
			if billingHandler.products[i].Quantity <= 0 {
				billingHandler.products = append(billingHandler.products[:i], billingHandler.products[i+1:]...)
			}

			return
		}
	}
}

func (billingHandler *BillingHandler) GetProducts() []utils.BillProduct {
	return billingHandler.products
}

func (bill *BillingHandler) CalculateTotal() float64 {
	var total float64 = 0

	for _, value := range bill.products {
		product, _ := bill.ProductRepo.GetProductById(value.Id)
		total += value.Quantity * product.UnitPrice
	}

	return utils.RoundToGivenDecimals(total, 2)
}

func (bill *BillingHandler) CalculateVAT() float64 {
	var totalVAT float64 = 0

	for _, value := range bill.products {
		product, _ := bill.ProductRepo.GetProductById(value.Id)
		VATPercentage := utils.VAT_PERCENTAGES_PER_CATEGORY[product.VATCategory]
		VATPerUnit := product.UnitPrice * (VATPercentage / 100)
		totalVAT += value.Quantity * VATPerUnit
	}

	return utils.RoundToGivenDecimals(totalVAT, 2)
}

func (billingHandler *BillingHandler) GetProductsWithInfos() []utils.Product {
	products := []utils.Product{}

	for _, value := range billingHandler.products {
		product, _ := billingHandler.ProductRepo.GetProductById(value.Id)

		products = append(products, *product)
	}

	return products
}

func (billingHandler *BillingHandler) GetProductsWithInfosForFormatter() []utils.ProductWithQuantityFromBill {
	products := []utils.ProductWithQuantityFromBill{}

	for _, value := range billingHandler.products {
		product, _ := billingHandler.ProductRepo.GetProductById(value.Id)

		products = append(products, utils.ProductWithQuantityFromBill{
			Product:  *product,
			Quantity: value.Quantity,
		})
	}

	return products
}

func (billingHandler *BillingHandler) FormatBill(uuid string, createdAt string) *bytes.Buffer {

	// Create a BillData DTO
	billData := utils.BillData{
		Products:  billingHandler.GetProductsWithInfosForFormatter(),
		Subtotal:  billingHandler.CalculateTotal() - billingHandler.CalculateVAT(),
		VATAmount: billingHandler.CalculateVAT(),
		Total:     billingHandler.CalculateTotal(),
		Id:        uuid,
		CreatedAt: createdAt,
	}

	formattedBill := billingHandler.Formatter.FormatBill(billData, billingHandler.Printer.GetRowLength())

	return &formattedBill
}

func (billingHandler *BillingHandler) PrintBill() {
	// we will send empty string for id and date, because we are just viewing the draft bill here,
	// so we don't have an id and date for the bill. We will provide them on saving the bill
	formattedBill := billingHandler.FormatBill("", "")

	// Print the formatted bill
	billingHandler.Printer.Print(*formattedBill)
}

func (billingHandler *BillingHandler) SaveBill() string {
	uuid := uuid.NewString()
	createdAt := time.Now().Format("02-01-2006 15:04:05")
	data := billingHandler.FormatBill(uuid, createdAt)

	billingHandler.Printer.Print(*data)

	billingHandler.BillRepo.AddBill(
		billingHandler.products,
		billingHandler.CalculateTotal()-billingHandler.CalculateVAT(),
		billingHandler.CalculateTotal(),
		createdAt,
		uuid,
	)

	fileName := "bill_" + uuid + ".txt"

	billingHandler.FileHandler.Save(data, fileName)

	return fileName
}
