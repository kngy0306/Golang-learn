package app

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
		pokeinfo[i].Type2 = l.Types[1].Type.Name
		pokeinfo[i].Img = fmt.Sprintf("https://pokeres.bastionbot.org/images/pokemon/%v.png", l.ID)
	}

	return nil
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/index.html")
	if err != nil {
		fmt.Println(err)
	}

	if err := t.Execute(w, pokeinfo); err != nil {
		fmt.Println(err)
	}
}
