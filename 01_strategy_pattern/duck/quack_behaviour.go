package duck

import "fmt"

type QuackBehavior interface {
	quack()
}

type Quack struct{}

func (q Quack) quack() {
	fmt.Println("Quack!")
}

type MuteQuack struct{}

func (m MuteQuack) quack() {
	fmt.Println("Mute Quack!")
}

type Squeak struct{}

func (s Squeak) quack() {
	fmt.Println("Squeak!")
}
