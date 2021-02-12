package poke

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/mtslzr/pokeapi-go"
)

// PokeInfo pokemon status struct
type PokeInfo struct {
	ID    int
	Type1 string
	Type2 string
	Img   string
}

/*
GetPokemon is fetch pokemon(name, url) limit=number
 name:"bulbasaur"url:"https://pokeapi.co/api/v2/pokemon/1/"
*/
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

var pokeinfo []PokeInfo

// ShowPokemon is show Pokemon's ID, Type1, Type2 and ImageURL
func ShowPokemon(pokeList []string) error {
	pokeinfo = make([]PokeInfo, len(pokeList))

	for i, name := range pokeList {
		l, err := pokeapi.Pokemon(name)
		if err != nil {
			return err
		}

		pokeinfo[i].ID = l.ID
		pokeinfo[i].Type1 = l.Types[0].Type.Name
		if len(l.Types) == 2 {
			pokeinfo[i].Type2 = l.Types[1].Type.Name
		}
		pokeinfo[i].Img = fmt.Sprintf("https://pokeres.bastionbot.org/images/pokemon/%v.png", l.ID)
	}

	return nil
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("front/index.html")
	if err != nil {
		fmt.Println(err)
	}

	if err := t.Execute(w, pokeinfo); err != nil {
		fmt.Println(err)
	}
}
