package models

type ProductQuantity struct {
	Product Product
	Count   uint32
}

type Product struct {
	Name string
}
