package utils

type Product struct {
	Id        string
	Name      string
	UnitPrice float64
}

type GetValidNumberFromInputOptions struct {
	ShouldBePositive bool
}
