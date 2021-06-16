package main

import "design-patterns-go/05_factory_method_pattern/pizza"

func main() {
	nyStore := &pizza.NYStore{}
	nyStore.OrderPizza("CheesePizza")

	californiaStore := &pizza.CaliforniaStore{}
	californiaStore.OrderPizza("CheesePizza")
}
