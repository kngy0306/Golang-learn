package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/mtslzr/pokeapi-go"
)

type PokeInfo struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Types []string `json:"types"`
	Img   string   `json:"img"`
}

func main() {
	http.HandleFunc("/pokedex", PokeHandler)
	http.Handle("/", http.FileServer(http.Dir("index")))
	http.ListenAndServe(":3000", nil)
}

func PokeHandler(w http.ResponseWriter, r *http.Request) {
	pokemons, err := getPokemon(9)
	if err != nil {
		fmt.Println(err)
	}

	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids) < 1 {
		ids = nil
	}

	if ids != nil {
		res := []PokeInfo{}
		id, _ := strconv.Atoi(ids[0])
		for _, pokemon := range pokemons {
			if id == pokemon.ID {
				res = append(res, pokemon)
			}
			pokemons = res
		}
	}

	bytes, err := json.Marshal(pokemons)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(bytes)

	w.Write(bytes)
}

func getPokemon(limit int) ([]PokeInfo, error) {
	pokeinfo := make([]PokeInfo, limit)

	for i := 0; i < limit; i++ {
		result, err := pokeapi.Pokemon(fmt.Sprint(i + 1))
		if err != nil {
			return nil, err
		}

		pokeinfo[i].ID = result.ID
		pokeinfo[i].Name = result.Name
		pokeinfo[i].Types = append(pokeinfo[i].Types, result.Types[0].Type.Name)
		if len(result.Types) == 2 {
			pokeinfo[i].Types = append(pokeinfo[i].Types, result.Types[1].Type.Name)
		}
		pokeinfo[i].Img = fmt.Sprintf("https://pokeres.bastionbot.org/images/pokemon/%v.png", result.ID)
	}

	return pokeinfo, nil
}
