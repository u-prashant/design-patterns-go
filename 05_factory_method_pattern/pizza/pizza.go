package pizza

import "fmt"

type Pizza interface {
	Prepare()
	Cut()
	Box()
}

type NYCheesePizza struct{}

func (c *NYCheesePizza) Prepare() {
	fmt.Println("Preparing NY Cheese Pizza...")
}

func (c *NYCheesePizza) Cut() {
	fmt.Println("Cutting NY Cheese Pizza...")
}

func (c *NYCheesePizza) Box() {
	fmt.Println("Boxing NY Cheese Pizza...")
}

type NYPepperoniPizza struct{}

func (p *NYPepperoniPizza) Prepare() {
	fmt.Println("Preparing NY Pepperoni Pizza...")
}

func (p *NYPepperoniPizza) Cut() {
	fmt.Println("Cutting NY Pepperoni Pizza...")
}

func (p *NYPepperoniPizza) Box() {
	fmt.Println("Boxing NY Pepperoni Pizza...")
}

type CaliforniaCheesePizza struct{}

func (c *CaliforniaCheesePizza) Prepare() {
	fmt.Println("Preparing California Cheese Pizza...")
}

func (c *CaliforniaCheesePizza) Cut() {
	fmt.Println("Cutting California Cheese Pizza...")
}

func (c *CaliforniaCheesePizza) Box() {
	fmt.Println("Boxing California Cheese Pizza...")
}

type CaliforniaPepperoniPizza struct{}

func (p *CaliforniaPepperoniPizza) Prepare() {
	fmt.Println("Preparing California Pepperoni Pizza...")
}

func (p *CaliforniaPepperoniPizza) Cut() {
	fmt.Println("Cutting California Pepperoni Pizza...")
}

func (p *CaliforniaPepperoniPizza) Box() {
	fmt.Println("Boxing California Pepperoni Pizza...")
}
