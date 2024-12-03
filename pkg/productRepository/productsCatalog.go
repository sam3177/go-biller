package productRepository

import (
	"biller/pkg/utils"

	"github.com/google/uuid"
)

var ProductsCatalog []utils.Product = []utils.Product{
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
