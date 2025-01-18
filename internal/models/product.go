package models

type ProductCount struct {
	Product Product
	Count   uint32
}

type Product struct {
	Name string
}
