package duck

import "fmt"

type Duck struct {
	flyBehavior   FlyBehavior
	quackBehavior QuackBehavior
}

func (d Duck) Fly() {
	d.flyBehavior.fly()
}

func (d Duck) Quack() {
	d.quackBehavior.quack()
}

func (d Duck) Swim() {
	fmt.Println("All ducks swim!")
}

func (d *Duck) SetFlyBehaviour(behavior FlyBehavior) {
	d.flyBehavior = behavior
}

func (d *Duck) SetQuackBehaviour(behavior QuackBehavior) {
	d.quackBehavior = behavior
}

type MiniDuckSimulator struct {
	Duck
}

func NewMiniDuckSimulator() MiniDuckSimulator {
	return MiniDuckSimulator{
		Duck{
			flyBehavior:   FlyWithWings{},
			quackBehavior: Quack{},
		},
	}
}

type ModelDuck struct {
	Duck
}

func NewModelDuck() ModelDuck {
	return ModelDuck{
		Duck{
			flyBehavior:   FlyNoWay{},
			quackBehavior: Quack{},
		},
	}
}
