package router

import (
	"github.com/nibble-4bits/ondemand-go-bootcamp/httpAPI/v1/controller"
	"github.com/nibble-4bits/ondemand-go-bootcamp/usecase"

	"github.com/gin-gonic/gin"
)

func getPokemonByID(r *gin.RouterGroup, service usecase.PokemonService) {
	r.GET("/pokemons/:id", controller.GetPokemonByID(service))
}

func getAllPokemons(r *gin.RouterGroup, service usecase.PokemonService) {
	r.GET("/pokemons", controller.GetAllPokemons(service))
}

func RegisterPokemonRoutes(r *gin.RouterGroup, service usecase.PokemonService) {
	getPokemonByID(r, service)
	getAllPokemons(r, service)
}
