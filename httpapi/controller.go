package httpapi

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nibble-4bits/ondemand-go-bootcamp/usecase"
)

func getPokemonByIDController(service usecase.PokemonService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		pokemon, err := service.GetByID(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, pokemon)
	}
}

func getAllPokemonsController(service usecase.PokemonService) gin.HandlerFunc {
	return func(c *gin.Context) {
		pokemons, err := service.GetAll()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, pokemons)
	}
}
