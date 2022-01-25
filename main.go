package main

import (
	"log"

	"github.com/nibble-4bits/ondemand-go-bootcamp/adapter"
	"github.com/nibble-4bits/ondemand-go-bootcamp/data"
	"github.com/nibble-4bits/ondemand-go-bootcamp/httpAPI"
	"github.com/nibble-4bits/ondemand-go-bootcamp/usecase"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	csvDataSource := data.NewCSVDataSource("pokemon.csv")

	pokemonAdapter, err := adapter.NewPokemonAdapter(csvDataSource)
	if err != nil {
		// Could handle the error here more gracefully, for example
		// we could try and fetch the pokemons from another data source.
		// For now, just exit fatally and log the error.
		log.Fatalln(err)
	}

	pokemonService := usecase.NewPokemonService(pokemonAdapter)

	httpAPI.StartServer(pokemonService)
}
