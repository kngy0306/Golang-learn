package main

import (
	"app/src/pokeAPI/app"
	"fmt"
	"net/http"
)

func main() {
	number := 3 // 取り出す数

	pokeList, _, err := app.GetPokemon(number)
	if err != nil {
		fmt.Println(err)
	}

	err = app.ShowPokemon(pokeList)
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/", app.IndexHandler)
	http.ListenAndServe(":3000", nil)
}
