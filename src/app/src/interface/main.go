package main

import (
	"app/src/interface/gadget"
	"fmt"
)

// Player インターフェースは、Play()とStop()を持つ型のみ実装可
type Player interface {
	Play(string)
	Stop()
}

func playList(p Player, songs []string) {
	for _, song := range songs {
		p.Play(song)
	}
	p.Stop()
}

func TryOut(p Player) {
	p.Play("Test Track")
	p.Stop()
	recoder, ok := p.(gadget.TapeRecoder)
	if ok {
		recoder.Recode()
	} else {
		fmt.Println("Player was not a TapeRecoder")
	}
}

func main() {
	TryOut(gadget.TapeRecoder{})
	TryOut(gadget.TapePlayer{})
}
