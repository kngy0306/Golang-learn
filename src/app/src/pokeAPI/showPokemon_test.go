package main

import (
	"testing"
)

func TestShowPokemon(t *testing.T) {
	pokemon := []string{"ヒトカゲ", "フシギダネ"}

	err := ShowPokemon(pokemon)
	if err != nil {
		t.Fatal("failed do ShowPokemon")
	}
}
