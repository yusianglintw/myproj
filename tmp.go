package main

import (
	"fmt"
)

type Circle struct {
	x, y, r float64
}

func zero(xPtr *int) {
	*xPtr++
}
func main() {
	x := 123
	zero(&x)
	fmt.Println(&x, x)
}
