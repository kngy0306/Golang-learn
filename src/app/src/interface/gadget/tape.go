package gadget

import "fmt"

type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing: ", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Stopped!")
}

type TapeRecoder struct {
	Microphones int
}

func (t TapeRecoder) Play(song string) {
	fmt.Println("Playing: ", song)
}

func (t TapeRecoder) Recode() {
	fmt.Println("Recoding")
}

func (t TapeRecoder) Stop() {
	fmt.Println("Stopped!")
}
