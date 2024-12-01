package utils

import (
	productsRepository "biller/pkg/productsRepo"

	"github.com/google/uuid"
)

var ProductsCatalog []productsRepository.Product = []productsRepository.Product{
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
