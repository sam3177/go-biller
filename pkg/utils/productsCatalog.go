package utils

import (
	productRepository "biller/pkg/productsRepo"

	"github.com/google/uuid"
)

var ProductsCatalog []productRepository.Product = []productRepository.Product{
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
