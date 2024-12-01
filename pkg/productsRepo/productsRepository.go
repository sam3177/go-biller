package productsRepository

type Product struct {
	Id        string
	Name      string
	UnitPrice float64
}

type ProductsRepositoryInterface interface {
	GetProductByID(id string) (Product, error) // Fetch product details by ID
	IsProductValid(id string) bool             // Check if a product is valid
	ListAllProducts() []Product                // Optional: List all products
}
