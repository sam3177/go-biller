package productsJsonStorageHandler

import (
	"biller/pkg/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"slices"

	"github.com/google/uuid"
)

type ProductsJSONStorageHandler struct {
	FilePath string
}

func NewProductsJSONStorageHandler(filePath string) *ProductsJSONStorageHandler {
	return &ProductsJSONStorageHandler{
		FilePath: filePath,
	}
}

func (handler *ProductsJSONStorageHandler) GetAllProducts() ([]utils.Product, error) {
	file, err := os.Open(handler.FilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var products []utils.Product
	if err := json.Unmarshal(bytes, &products); err != nil {
		log.Fatalf("failed to parse JSON: %w", err)
	}

	return products, nil
}

func (handler *ProductsJSONStorageHandler) GetProduct(id string) (*utils.Product, error) {
	products, err := handler.GetAllProducts()
	if err != nil {
		return nil, fmt.Errorf("failed to read products: %w", err)
	}

	for _, product := range products {
		if product.Id == id {
			return &product, nil
		}
	}

	return nil, fmt.Errorf("product with ID %v not found", id)
}

func (handler *ProductsJSONStorageHandler) writeProducts(products []utils.Product) error {
	bytes, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to serialize products: %w", err)
	}

	// Append a newline character
	bytes = append(bytes, '\n')

	if err := os.WriteFile(handler.FilePath, bytes, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func (handler *ProductsJSONStorageHandler) UpdateProduct(productData utils.Product) error {
	products, err := handler.GetAllProducts()
	if err != nil {
		return fmt.Errorf("failed to read products: %w", err)
	}

	index := slices.IndexFunc(products, func(product utils.Product) bool {
		return product.Id == productData.Id
	})

	if index == -1 {
		return fmt.Errorf("product with ID %v not found", productData.Id)
	}

	products[index] = productData

	handler.writeProducts(products)

	return nil
}

func (handler *ProductsJSONStorageHandler) AddProduct(newProduct utils.Product) (*utils.Product, error) {
	products, err := handler.GetAllProducts()
	if err != nil {
		return nil, fmt.Errorf("failed to read products: %w", err)
	}

	for _, product := range products {
		if product.Name == newProduct.Name {
			return nil, fmt.Errorf("the product '%v' already exists", newProduct.Name)
		}
	}

	newProduct.Id = uuid.NewString()

	products = append(products, newProduct)

	if err := handler.writeProducts(products); err != nil {
		return nil, fmt.Errorf("failed to write updated products: %w", err)
	}

	return &newProduct, nil
}

func (handler *ProductsJSONStorageHandler) SeedJSONFile(productsCatalog []utils.Product) error {
	products, err := handler.GetAllProducts()
	if err != nil {
		return fmt.Errorf("failed to read products: %w", err)
	}
	if len(products) > 0 {
		return fmt.Errorf("products list is already populated")
	}

	productsToBeAdded := []utils.Product{}
	for _, product := range productsCatalog {
		product.Id = uuid.NewString()
		productsToBeAdded = append(productsToBeAdded, product)
	}

	handler.writeProducts(productsToBeAdded)

	return nil
}
