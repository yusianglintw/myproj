package main

import "fmt"

type Product struct {
	price float64
}

type Member struct {
	level string
	cash  float64
	point float64
}

func NewProduct(price float64) *Product {
	product := new(Product)
	product.setPrice(price)
	return product

}

func NewMember(level string, cash float64, point float64) *Member {
	member := new(Member)
	member.setLevel(level)
	member.setCash(cash)
	member.setPoint(point)
	return member
}

func (member *Member) setLevel(level string) {
	member.level = level
}
func (member *Member) setCash(cash float64) {
	member.cash = cash
}
func (member *Member) setPoint(point float64) {
	member.point = point
}

func (member *Member) GetLevel() string {
	return member.level
}

func (member *Member) GetCash() float64 {
	return member.cash
}

func (member *Member) GetPoint() float64 {
	return member.point
}

func (product *Product) setPrice(price float64) {
	product.price = price
}

func (product *Product) GetPrice() float64 {
	return product.price
}

type Pay interface {
	pay(m *Member)
}

func performPay(p Pay, m *Member) {
	p.pay(m)
}

type NormalPay struct {
	cash float64
}

type VipPay struct {
	cash  float64
	point float64
	vip   string
}

type PointDiscount struct {
	cash  float64
	point float64
}

type AdditionPay struct {
	cash  float64
	point float64
}

func (p NormalPay) pay(m *Member) {
	cash := m.GetCash() - p.cash
	m.setCash((cash))
}

func (p VipPay) pay(m *Member) {
	fmt.Println(p.vip, p.cash)
}

func (p PointDiscount) pay(m *Member) {
	point := m.GetPoint() - p.point
	cash := m.GetCash() - p.cash + p.point
	m.setCash(cash)
	m.setPoint(point)
}

func (p AdditionPay) pay(m *Member) {
	if m.GetLevel() != "Vip0" {
		if p.point >= 100 {
			point := m.GetPoint() - p.point
			cash := m.GetCash() - (p.cash - p.point*0.9)
			m.setCash(cash)
			m.setPoint(point)
		}
	} else {
		point := m.GetPoint() - p.point
		cash := m.GetCash() - (p.cash - p.point)
		m.setCash(cash)
		m.setPoint(point)
	}
}

func main() {
	m1 := NewMember("Vip1", 10000, 10000)

	// p1 := NewProduct(50)
	// p2 := NewProduct(100)
	NP := NormalPay{cash: 100}
	performPay(NP, m1)
	NP.pay2()
	fmt.Println(m1)
	// fmt.Println("This is my product", p1.GetPrice(), p2.GetPrice())
}
