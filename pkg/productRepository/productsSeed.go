package productRepository

import (
	"biller/pkg/utils"
)

var ProductsSeed []utils.Product = []utils.Product{
	{Name: "Banana", UnitPrice: 1, UnitType: "kg", Stock: 40},
	{Name: "Apple", UnitPrice: 2, UnitType: "kg", Stock: 30},
	{Name: "Orange", UnitPrice: 3, UnitType: "kg", Stock: 10},
	{Name: "Avocado", UnitPrice: 2.19, UnitType: "piece", Stock: 34},
}
