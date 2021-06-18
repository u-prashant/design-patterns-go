package command

import "fmt"

type Light struct{}

func (l *Light) On() {
	fmt.Println("Switching on the light.")
}

func (l *Light) Off() {
	fmt.Println("Switching off the light.")
}

type Fan struct{}

func (f *Fan) On() {
	fmt.Println("Switching on the fan.")
}

func (f *Fan) Off() {
	fmt.Println("Switching off the fan.")
}
