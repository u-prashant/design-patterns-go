package main

import (
	"fmt"

	"design-patterns-go/01_strategy_pattern/duck"
)

func main() {
	miniDuckSimulator := duck.NewMiniDuckSimulator()
	miniDuckSimulator.Fly()
	miniDuckSimulator.Quack()
	miniDuckSimulator.Swim()
	fmt.Println()

	modelDuck := duck.NewModelDuck()
	modelDuck.Fly()
	modelDuck.Quack()
	modelDuck.Swim()
	modelDuck.SetFlyBehaviour(duck.FlyWithWings{})
	modelDuck.Fly()
}
