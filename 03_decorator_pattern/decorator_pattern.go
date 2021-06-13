package main

import (
	"fmt"

	"design-patterns-go/03_decorator_pattern/beverage"
)

func main() {
	var b beverage.Beverage
	b = beverage.NewExpresso(beverage.LARGE)
	fmt.Printf("%s $%f\n", b.GetDescription(), b.Cost())

	b = beverage.NewMocha(b)
	fmt.Printf("%s $%f\n", b.GetDescription(), b.Cost())

	b = beverage.NewMocha(b)
	fmt.Printf("%s $%f\n", b.GetDescription(), b.Cost())
}
