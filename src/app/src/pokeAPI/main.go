package main

import (
	"fmt"
	"net/http"
)

func main() {
	number := 9 // 取り出す数

	pokeList, _, err := GetPokemon(number)
	if err != nil {
		fmt.Println(err)
	}

	err = ShowPokemon(pokeList)
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":3000", nil)
}
