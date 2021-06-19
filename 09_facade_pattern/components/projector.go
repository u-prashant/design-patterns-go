package components

import "fmt"

type ScreenMode int

const (
	Small ScreenMode = iota
	Medium
	Wide
)

type Projector struct {
	screenMode ScreenMode
}

func (p *Projector) On() {
	fmt.Println("Projector Switched on.")
}

func (p *Projector) WideScreenMode() {
	p.screenMode = Wide
	fmt.Println("Projector Screen Mode set to Wide.")
}
