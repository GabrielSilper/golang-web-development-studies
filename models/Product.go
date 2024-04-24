package models

type Product struct {
	Id                int
	Name, Description string
	Price             float64
	Quantity          int
}

func NewProduct(id, quantity int, name, description string, price float64) Product {
	return Product{id, name, description, price, quantity}
}
