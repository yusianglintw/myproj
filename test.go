package main

import "fmt"

type Product struct {
	price float64
}

func NewProduct(price float64) *Product {
	product := new(Product)
	product.setPrice(price)
	return product

}
func (product *Product) setPrice(price float64) {
	product.price = price
}

func (product *Product) GetPrice() float64 {
	return product.price
}

func main() {
	p1 := NewProduct(100)
	p2 := NewProduct(200)
	var val float64 = 0
	// p3 := NewProduct(300)
	// p4 := NewProduct(400)
	p_l := [2]*Product{p1, p2}
	for _, p := range p_l {
		val += p.GetPrice()
	}
	fmt.Println(val)
}
