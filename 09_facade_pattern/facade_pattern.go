package main

import (
	"design-patterns-go/09_facade_pattern/components"
	"design-patterns-go/09_facade_pattern/facade"
)

func main() {
	amp := components.Amplifier{}
	player := components.Player{}
	projector := components.Projector{}

	h := facade.NewHomeTheatre(amp, player, projector)
	h.WatchMovie("3 Idiots")
}
