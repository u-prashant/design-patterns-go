package turkey

import "fmt"

type Turkey interface {
	Gobble()
	Fly()
}

type WildTurkey struct{}

func (w *WildTurkey) Gobble() {
	fmt.Println("Gobble...")
}

func (w *WildTurkey) Fly() {
	fmt.Println("Short Fly...")
}

type Adapter struct {
	Turkey Turkey
}

func (a *Adapter) Quack() {
	a.Turkey.Gobble()
}

func (a *Adapter) Fly() {
	for i := 0; i < 5; i++ {
		a.Turkey.Fly()
	}
}
