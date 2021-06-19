package duck

import "fmt"

type Duck interface {
	Quack()
	Fly()
}

type MallardDuck struct{}

func (m *MallardDuck) Quack() {
	fmt.Println("Quacking...")
}

func (m *MallardDuck) Fly() {
	fmt.Println("Flying...")
}
