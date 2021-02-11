package main

import (
	"github.com/mtslzr/pokeapi-go"
)

// GetPokemon is fetch pokemon(name, url) limit=number
func GetPokemon(number int) ([]string, []string, error) {
	NameAry := make([]string, 0, number)
	URLAry := make([]string, 0, number)

	g, err := pokeapi.Resource("pokemon", 0, number)
	if err != nil {
		return nil, nil, err
	}

	for i := 0; i < number; i++ {
		result := g.Results[i]
		NameAry = append(NameAry, result.Name)
		URLAry = append(URLAry, result.URL)
	}

	return NameAry, URLAry, nil
}
