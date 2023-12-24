package entity

type Payment struct {
	Id                    int
	Username, ProductCode string
	Amount, UnitPrice     float64
}

func NewPayment(id int, username, productCode string, amount, unitPrice float64) *Payment {
	return &Payment{id, username, productCode, amount, unitPrice}
}

//...
