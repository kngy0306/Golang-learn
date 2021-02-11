package main

import "testing"

func TestGetPokemon(t *testing.T) {
	length := 10

	res1, res2, err := GetPokemon(length)
	if err != nil {
		t.Fatal(err)
	}

	if len(res1) != length && len(res2) != length {
		t.Fatal("failed test")
	}
}
