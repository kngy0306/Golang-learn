package main

import (
	"fmt"
)

type item struct {
	name string
}

type character struct {
	name     string
	leftHand *item
}

func (c *character) pickup(i *item) {
	if c == nil || i == nil {
		return
	}

	fmt.Printf("%v は、%v を持った。\n", c.name, c.leftHand.name)
	c.leftHand = i
}

func (c *character) give(to *character) {
	if c == nil || to == nil {
		return
	}

	// 左手に何もないとき
	if c.leftHand == nil {
		// some print
		return
	}

	// 相手の左手が空いていないとき
	if to.leftHand != nil {
		// some print
		return
	}

	to.leftHand = c.leftHand
	c.leftHand = nil
	fmt.Printf("%v は、 %v に、 %v を与えた。\n", c.name, to.name, to.leftHand.name)
}

func (c character) String() string {
	if c.leftHand == nil {
		return fmt.Sprintf("%v は、何も所持していない。\n", c.name)
	}
	return fmt.Sprintf("%v は、 %v を所持している。\n", c.name, c.leftHand.name)
}

func main() {
	sword := &item{name: "sword"}
	knight := &character{name: "knight"}
	arthur := &character{name: "arthur"}

	knight.pickup(sword)
	knight.give(arthur)

	fmt.Print(knight)
	fmt.Print(arthur)
}
