package main

import (
	"log"

	"github.com/nibble-4bits/ondemand-go-bootcamp/adapter"
	"github.com/nibble-4bits/ondemand-go-bootcamp/data"
	"github.com/nibble-4bits/ondemand-go-bootcamp/httpapi"
	"github.com/nibble-4bits/ondemand-go-bootcamp/usecase"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	csvDataSource := data.NewCSVDataSource("pokemon.csv")
	pokemonAdapter := adapter.NewPokemonAdapter(csvDataSource)
	pokemonService := usecase.NewPokemonService(pokemonAdapter)

	httpapi.StartServer(pokemonService)
}
