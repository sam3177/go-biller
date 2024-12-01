package utils

import (
	productsRepository "biller/pkg/productsRepo"
	"fmt"
	"os"
	"os/exec"
	"slices"
)

func GetProductById(products []productsRepository.Product, id string) (*productsRepository.Product, error) {
	index := slices.IndexFunc(products, func(product productsRepository.Product) bool {
		return product.Id == id
	})

	if index == -1 {
		return nil, fmt.Errorf("product with ID %v not found", id)
	}

	return &products[index], nil
}

func CheckIfProductIsValid(products []productsRepository.Product, id string) bool {
	index := slices.IndexFunc(products, func(product productsRepository.Product) bool {
		return product.Id == id
	})

	return index != -1
}

func OpenFileInVsCode(filePath string) {
	cmd := exec.Command("code", filePath) // `code` is the VSCode CLI command
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if error := cmd.Run(); error != nil {
		fmt.Printf("Failed to open file in VSCode: %v", error)
		panic(error)
	}
}
