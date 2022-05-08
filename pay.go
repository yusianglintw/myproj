package main

import "fmt"

type PayInterface interface {
	pay()
}

func PayCash(pi PayInterface) {
	payinrerface := pi
	payinrerface.pay()
}

type NormalPay struct {
	cash  float64
	point int
}

type VipPay struct {
	cash  float64
	point int
	vip   string
}

type PointDiscount struct {
	cash  float64
	point int
}

type AdditionPay struct {
	cash float64
}

func (p NormalPay) pay() {
	fmt.Println(p.cash)
}

func (p VipPay) pay() {
	fmt.Println(p.vip, p.cash)
}

func (p PointDiscount) pay() {
	fmt.Println(p.point, p.cash)
}

func (p AdditionPay) pay() {
	fmt.Println("100", p.cash)
}

func main() {
	p := PointDiscount{cash: 100, point: 500}
	PayCash(p)
}
