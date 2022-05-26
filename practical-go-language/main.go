package main

import (
	"fmt"
)

func main() {
	ary := [5]int{1, 2, 3, 4, 5}
	for _, index := range ary {
		fmt.Println(index)
	}
}
