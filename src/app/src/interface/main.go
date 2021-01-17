package main

import (
	"fmt"
	"math/rand"
)

type member struct {
	name string
}

func (m member) String() string {
	return fmt.Sprint(m.name)
}

func (m member) move() string {
	return "テレビ局へ向かう。"
}

func (m member) eat() string {
	if rand.Intn(2) == 0 {
		return fmt.Sprintf("%v", "納豆を食べる。")
	}
	return fmt.Sprintf("%v", "白米を食べる。")
}

type animal interface {
	move() string
	eat() string
}

func act(a animal) {
	switch rand.Intn(2) {
	case 0:
		fmt.Printf("%v は %v をした。", a, a.eat())
	case 1:
		fmt.Printf("%v は %v をした。", a, a.move())
	}
}

func main() {

}
