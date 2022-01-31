package main

import (
	"log"

	"github.com/nibble-4bits/ondemand-go-bootcamp/adapter"
	"github.com/nibble-4bits/ondemand-go-bootcamp/config"
	"github.com/nibble-4bits/ondemand-go-bootcamp/data"
	routerV1 "github.com/nibble-4bits/ondemand-go-bootcamp/httpAPI/v1/router"
	"github.com/nibble-4bits/ondemand-go-bootcamp/usecase"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	csvDataSourcePokemon := data.NewCSVDataSource(config.Config.PokemonCSVPath)

	pokemonAdapter, err := adapter.NewPokemonAdapter(csvDataSourcePokemon)
	if err != nil {
		// Could handle the error here more gracefully, for example
		// we could try and fetch the pokemons from another data source.
		// For now, just exit fatally and log the error.
		log.Fatalln(err)
	}

	pokemonService := usecase.NewPokemonService(pokemonAdapter)

	// ==========================================================================

	csvDataSourceComment := data.NewCSVDataSource(config.Config.CommentCSVPath)
	httpDataSourceComment := data.NewHTTPDataSource()
	csvDataStoreComment := data.NewCSVDataStore(config.Config.CommentCSVPath)

	commentAdapter, err := adapter.NewCommentAdapter(csvDataSourceComment, httpDataSourceComment, csvDataStoreComment)
	if err != nil {
		log.Fatalln(err)
	}

	commentService := usecase.NewCommentService(commentAdapter)

	// ==========================================================================

	router := routerV1.CreateRouter()
	routerGroup := routerV1.CreateRouterGroup(router)

	routerV1.RegisterPokemonRoutes(routerGroup, pokemonService)
	routerV1.RegisterCommentRoutes(routerGroup, commentService)

	routerV1.StartServer(router)
}
