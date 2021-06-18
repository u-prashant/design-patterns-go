package command

import "fmt"

type Command interface {
	Execute()
}

type NoCommand struct{}

func (n *NoCommand) Execute() {
	fmt.Println("No Command")
}

type LightOnCommand struct {
	Light Light
}

func (l *LightOnCommand) Execute() {
	l.Light.On()
}

type LightOffCommand struct {
	Light Light
}

func (l *LightOffCommand) Execute() {
	l.Light.Off()
}

type FanOnCommand struct {
	Fan Fan
}

func (f *FanOnCommand) Execute() {
	f.Fan.On()
}

type FanOffCommand struct {
	Fan Fan
}

func (f *FanOffCommand) Execute() {
	f.Fan.Off()
}
