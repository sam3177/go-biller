package productRepository

import (
	"biller/pkg/utils"
)

var ProductsSeed []utils.Product = []utils.Product{
	{Name: "Banana", UnitPrice: 1, UnitType: "kg", Stock: 40, VATCategory: utils.A},
	{Name: "Apple", UnitPrice: 2, UnitType: "kg", Stock: 30, VATCategory: utils.A},
	{Name: "Orange", UnitPrice: 3, UnitType: "kg", Stock: 10, VATCategory: utils.A},
	{Name: "Avocado", UnitPrice: 2.19, UnitType: "piece", Stock: 34, VATCategory: utils.A},
	{Name: "Toilet Paper", UnitPrice: .1, UnitType: "piece", Stock: 34, VATCategory: utils.B},
}
