package main

import (
	"fmt"
	"strings"
)

type kona struct{ name string }
type word struct{ words string }

type talker interface {
	talk() string
}

func (k kona) talk() string {
	return fmt.Sprint(k.name)
}

func (w word) talk() string {
	return fmt.Sprint(w.words)
}

// func shout(t kona) {
// 	louder := strings.ToUpper(t.talk() + "!!")
// 	fmt.Println(louder)
// }

// func wshout(w word) {
// 	louder := strings.ToUpper(w.talk() + "!!")
// 	fmt.Println(louder)
// }
func shout(t talker) {
	louder := strings.ToUpper("This is " + t.talk() + "!!")
	fmt.Println(louder)
}

func main() {
	t := kona{"kona"}
	shout(t)

	w := word{"nogizaka"}
	shout(w)
}
