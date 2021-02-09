package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, x int
	fmt.Scan(&n)
	fmt.Scan(&x)

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s := sc.Text()

	slice := strings.Split(s, " ")

	for _, str := range slice {
		tmp, _ := strconv.Atoi(str)

		if tmp != x {
			fmt.Printf("%d ", tmp)
		}
	}

	fmt.Println()
}
