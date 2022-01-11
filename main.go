package main

import (
	"fmt"

	"github.com/nibble-4bits/ondemand-go-bootcamp/adapter"
	"github.com/nibble-4bits/ondemand-go-bootcamp/data"
	"github.com/nibble-4bits/ondemand-go-bootcamp/usecase"
)

func main() {
	csvDataSource := data.NewCSVDataSource("pokemon.csv")

	pokemonAdapter := adapter.NewPokemonAdapter(csvDataSource)

	pokemonService := usecase.NewPokemonService(pokemonAdapter)

	p := pokemonService.GetByID(145)

	fmt.Printf("p: %v\n", p)
}
