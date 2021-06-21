package main

import "design-patterns-go/11_iterator_pattern/menu"

func main() {
	w := menu.Waitress{
		PancakeHouse: menu.NewPancakeHouseMenu(),
		DinerMenu:    menu.NewDinerMenu(),
	}
	w.PrintMenu()
}
