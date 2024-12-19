package productsJsonStorageHandler

import (
	"biller/mocks"
	"biller/pkg/utils"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTempFileWithContent(t *testing.T, content []byte) string {
	t.Helper()
	tempFile, err := os.CreateTemp("", "products_test_*.json")
	assert.NoError(t, err)

	if content != nil {
		_, err := tempFile.Write(content)
		assert.NoError(t, err)
	}

	assert.NoError(t, tempFile.Close())
	return tempFile.Name()
}

func setupBeforeTest(t *testing.T) ([]utils.Product, string) {
	products := mocks.GetMockProductsCopy()
	content, _ := json.Marshal(products)
	tempFile := createTempFileWithContent(t, content)

	return products, tempFile
}

func TestGetAllProducts(t *testing.T) {
	// Arrange

	products, tempFile := setupBeforeTest(t)
	defer os.Remove(tempFile)

	handler := NewProductsJSONStorageHandler(tempFile)

	// Act
	result, err := handler.GetAllProducts()

	// Assert
	assert.NoError(t, err)
	assert.Len(t, result, 3)
	assert.Equal(t, products[0].Name, result[0].Name)
	assert.Equal(t, products[1].Name, result[1].Name)
}

func TestGetProduct(t *testing.T) {
	// Arrange
	_, tempFile := setupBeforeTest(t)
	defer os.Remove(tempFile)

	handler := NewProductsJSONStorageHandler(tempFile)

	// Act
	product, err := handler.GetProduct("1")

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "Product 1", product.Name)

	// Test for non-existing product
	_, err = handler.GetProduct("999")
	assert.Error(t, err)
}

func TestAddProduct(t *testing.T) {
	// Arrange
	_, tempFile := setupBeforeTest(t)
	defer os.Remove(tempFile)

	handler := NewProductsJSONStorageHandler(tempFile)

	newProduct := utils.Product{
		Name:      "NewProduct",
		UnitPrice: 15.0,
		Stock:     30,
	}

	// Act
	addedProduct, err := handler.AddProduct(newProduct)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, addedProduct)
	assert.Equal(t, "NewProduct", addedProduct.Name)

	// Verify file content
	allProducts, err := handler.GetAllProducts()
	assert.NoError(t, err)
	assert.Len(t, allProducts, 4)
	assert.Equal(t, "NewProduct", allProducts[3].Name)

	// Attempt adding a duplicate product
	_, err = handler.AddProduct(newProduct)
	assert.Error(t, err)
}

func TestUpdateProduct(t *testing.T) {
	// Arrange
	products, tempFile := setupBeforeTest(t)
	defer os.Remove(tempFile)

	handler := NewProductsJSONStorageHandler(tempFile)

	updatedProduct := products[1]
	updatedProduct.Stock = 50.0

	// Act
	err := handler.UpdateProduct(updatedProduct)

	// Assert
	assert.NoError(t, err)

	// Verify file content
	allProducts, err := handler.GetAllProducts()
	assert.NoError(t, err)
	assert.Len(t, allProducts, 3)
	assert.Equal(t, allProducts[1].Stock, 50.0)
	assert.Equal(t, allProducts[1].Name, "Product 2")

	// Test updating a non-existent product
	err = handler.UpdateProduct(utils.Product{Id: "999", Name: "NonExistent"})
	assert.Error(t, err)
}

func TestSeedJSONFile(t *testing.T) {
	// Arrange
	emptyFile := createTempFileWithContent(t, []byte("[]"))
	defer os.Remove(emptyFile)

	handler := NewProductsJSONStorageHandler(emptyFile)

	productsCatalog := []utils.Product{
		{Name: "SeedProduct1", UnitPrice: 10.0, Stock: 100},
		{Name: "SeedProduct2", UnitPrice: 20.0, Stock: 200},
	}

	// Act
	err := handler.SeedJSONFile(productsCatalog)

	// Assert
	assert.NoError(t, err)

	// Verify file content
	allProducts, err := handler.GetAllProducts()
	assert.NoError(t, err)
	assert.Len(t, allProducts, 2)
	assert.Equal(t, "SeedProduct1", allProducts[0].Name)
	assert.Equal(t, "SeedProduct2", allProducts[1].Name)

	// Test seeding an already populated file
	err = handler.SeedJSONFile(productsCatalog)
	assert.Error(t, err)
}
