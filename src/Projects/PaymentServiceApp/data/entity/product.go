package entity

type Product struct {
	Code, Name             string
	Stock, Cost, UnitPrice float64
}

func NewProduct(code, name string, stock, cost, unitPrice float64) *Product {
	return &Product{code, name, stock, cost, unitPrice}
}

//...
