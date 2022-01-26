package router

import (
	"github.com/nibble-4bits/ondemand-go-bootcamp/httpAPI/controller"
	"github.com/nibble-4bits/ondemand-go-bootcamp/usecase"

	"github.com/gin-gonic/gin"
)

func getPokemonByID(r *gin.Engine, service usecase.PokemonService) {
	r.GET("/pokemons/:id", controller.GetPokemonByID(service))
}

func getAllPokemons(r *gin.Engine, service usecase.PokemonService) {
	r.GET("/pokemons", controller.GetAllPokemons(service))
}

func registerRoutes(r *gin.Engine, service usecase.PokemonService) {
	getPokemonByID(r, service)
	getAllPokemons(r, service)
}
