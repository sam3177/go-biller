package main

import "github.com/google/uuid"

var productsCatalog []Product = []Product{
	{
		id:        uuid.NewString(),
		name:      "Banana",
		unitPrice: 1,
	},
	{
		id:        uuid.NewString(),
		name:      "Apple",
		unitPrice: 1.2,
	},
}
