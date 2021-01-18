package main

import "fmt"

type turtle struct {
	x, y int
}

func (t *turtle) Up() {
	t.y--
}

func (t *turtle) Right() {
	t.x++
}
func (t *turtle) Down() {
	t.y++
}
func (t *turtle) Left() {
	t.x--
}

func main() {
	turtle := turtle{}
	fmt.Println(turtle)

	turtle.Up()
	turtle.Up()
	turtle.Up()

	fmt.Println(turtle)
}
