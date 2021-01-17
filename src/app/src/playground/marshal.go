package marshal

import (
	"fmt"
	"strings"
)

type kona struct{}
type word struct{}

func (k kona) talk() string {
	return fmt.Sprint("This is kona")
}

func (w word) talk() string {
	return fmt.Sprint("word")
}

func shout(t kona) {
	louder := strings.ToUpper(t.talk() + "!!")
	fmt.Println(louder)
}

func wshout(w word) {
	louder := strings.ToUpper(w.talk() + "!!")
	fmt.Println(louder)
}

func main() {
	t := kona{}
	shout(t)

	w := word{}
	wshout(w)
}
