package main

import (
	"fmt"

	"design-patterns-go/08_adapter_pattern/duck"
	"design-patterns-go/08_adapter_pattern/turkey"
)

func main() {
	d := &duck.MallardDuck{}
	t := &turkey.WildTurkey{}
	turkeyAdapter := &turkey.Adapter{
		Turkey: t,
	}
	fmt.Println("The Turkey says...")
	t.Gobble()
	t.Fly()

	fmt.Println("The Duck says...")
	testDuck(d)

	fmt.Println("The TurkeyAdaptor says...")
	testDuck(turkeyAdapter)
}

func testDuck(d duck.Duck) {
	d.Quack()
	d.Fly()
}
