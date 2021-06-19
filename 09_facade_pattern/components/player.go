package components

import "fmt"

type Player struct {
	movie string
}

func (p *Player) On() {
	fmt.Println("Player switched On.")
}

func (p *Player) Play(movie string) {
	p.movie = movie
	fmt.Printf("Player playing movie %s\n", p.movie)
}
