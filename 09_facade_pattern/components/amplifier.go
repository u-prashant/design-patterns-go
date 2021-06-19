package components

import "fmt"

type Amplifier struct {
	player string
	volume int
}

func (a *Amplifier) On() {
	fmt.Println("Amplifier switched on.")
}

func (a *Amplifier) SetStreamingPlayer(player string) {
	a.player = player
	fmt.Printf("Amplifier Streaming Player set to %s\n", a.player)
}

func (a *Amplifier) SetVolume(volume int) {
	a.volume = volume
	fmt.Printf("Amplifier Volume set to %d\n", a.volume)
}
