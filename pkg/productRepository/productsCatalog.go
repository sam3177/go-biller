package productRepository

import (
	"biller/pkg/utils"

	"github.com/google/uuid"
)

var ProductsSlice []utils.Product = []utils.Product{
	{
		Id:        uuid.NewString(),
		Name:      "Banana",
		UnitPrice: 1,
	},
	{
		Id:        uuid.NewString(),
		Name:      "Apple",
		UnitPrice: 1.2,
	},
	{
		Id:        uuid.NewString(),
		Name:      "Orange",
		UnitPrice: .7,
	},
}

func MakeProductsCatalog(products []utils.Product) []utils.Product {
	var productsCatalog = []utils.Product{}

	for _, item := range products {
		newProduct := NewProduct(
			item.Id,
			item.Name,
			item.UnitPrice,
			item.UnitType,
			item.Stock,
		)

		productsCatalog = append(productsCatalog, *newProduct)
	}

	return productsCatalog
}
