package pizza

type SimplePizzaFactory struct{}

func (s SimplePizzaFactory) CreatePizza(pizzaName string) Pizza {
	var pizza Pizza

	switch pizzaName {
	case "CheesePizza":
		pizza = &CheesePizza{}
	case "PepperoniPizza":
		pizza = &PepperoniPizza{}
	}
	return pizza
}
