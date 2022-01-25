package httpAPI

import (
	"github.com/nibble-4bits/ondemand-go-bootcamp/usecase"

	"github.com/gin-gonic/gin"
)

func getPokemonByID(r *gin.Engine, service usecase.PokemonService) {
	r.GET("/pokemon/:id", getPokemonByIDController(service))
}

func getAllPokemons(r *gin.Engine, service usecase.PokemonService) {
	r.GET("/pokemon", getAllPokemonsController(service))
}

func registerRoutes(r *gin.Engine, service usecase.PokemonService) {
	getPokemonByID(r, service)
	getAllPokemons(r, service)
}
