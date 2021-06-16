package pizza

type store interface {
	createPizza(string) Pizza
}

type Store struct {
	store
}

func (s *Store) OrderPizza(pizzaName string) Pizza {
	pizza := s.createPizza(pizzaName)
	pizza.Prepare()
	pizza.Cut()
	pizza.Box()
	return pizza
}

type NYStore struct {
	Store
}

func (store *NYStore) createPizza(pizzaName string) Pizza {
	var pizza Pizza
	switch pizzaName {
	case "CheesePizza":
		pizza = &NYCheesePizza{}
	case "PepperoniPizza":
		pizza = &NYPepperoniPizza{}
	}
	return pizza
}

type CaliforniaStore struct {
	Store
}

func (store *CaliforniaStore) createPizza(pizzaName string) Pizza {
	var pizza Pizza
	switch pizzaName {
	case "CheesePizza":
		pizza = &CaliforniaCheesePizza{}
	case "PepperoniPizza":
		pizza = &CaliforniaPepperoniPizza{}
	}
	return pizza
}
