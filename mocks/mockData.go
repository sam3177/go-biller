package mocks

import (
	"biller/pkg/utils"
)

var MockProducts = []utils.Product{
	{Id: "1", Name: "Product 1", UnitPrice: 1, UnitType: "kg", Stock: 40, VATCategory: utils.A},
	{Id: "2", Name: "Product 2", UnitPrice: 2, UnitType: "piece", Stock: 30, VATCategory: utils.A},
	{Id: "3", Name: "Product 3", UnitPrice: 3, UnitType: "kg", Stock: 10, VATCategory: utils.A},
}

func GetMockProductsCopy() []utils.Product {
	copied := make([]utils.Product, len(MockProducts))
	copy(copied, MockProducts) // Copy only the top-level slice
	return copied
}
