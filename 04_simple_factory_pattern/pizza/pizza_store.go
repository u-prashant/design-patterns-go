package pizza

type Store struct {
	Factory SimplePizzaFactory
}

func (s Store) OrderPizza(pizzaName string) Pizza {
	pizza := s.Factory.CreatePizza(pizzaName)
	pizza.Prepare()
	pizza.Cut()
	pizza.Box()
	return pizza
}
