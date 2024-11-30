package main

import (
	"fmt"
	"slices"
)

type Product struct {
	id        string
	name      string
	unitPrice float64
}

func getProductById(id string) (*Product, error) {
	index := slices.IndexFunc(productsCatalog, func(product Product) bool {
		return product.id == id
	})

	if index == -1 {
		return nil, fmt.Errorf("product with ID %v not found", id)
	}

	return &productsCatalog[index], nil
}

func checkIfProductIsValid(id string) bool {
	index := slices.IndexFunc(productsCatalog, func(product Product) bool {
		return product.id == id
	})

	return index != -1
}
