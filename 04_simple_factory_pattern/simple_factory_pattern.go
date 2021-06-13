package main

import "design-patterns-go/04_simple_factory_pattern/pizza"

func main() {
	pizzaStore := &pizza.Store{
		Factory: pizza.SimplePizzaFactory{},
	}
	pizza := pizzaStore.OrderPizza("CheesePizza")
	pizza = pizzaStore.OrderPizza("PepperoniPizza")
	_ = pizza
}
