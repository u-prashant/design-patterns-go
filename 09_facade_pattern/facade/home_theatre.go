package facade

import (
	"fmt"

	"design-patterns-go/09_facade_pattern/components"
)

type HomeTheatre struct {
	amp       components.Amplifier
	player    components.Player
	projector components.Projector
}

func (h *HomeTheatre) WatchMovie(movie string) {
	fmt.Println("Get ready to watch a movie...")
	h.projector.On()
	h.projector.WideScreenMode()
	h.amp.On()
	h.amp.SetStreamingPlayer("CD")
	h.amp.SetVolume(5)
	h.player.On()
	h.player.Play(movie)
}

func NewHomeTheatre(amp components.Amplifier, player components.Player, projector components.Projector) *HomeTheatre {
	return &HomeTheatre{
		amp:       amp,
		player:    player,
		projector: projector,
	}
}
