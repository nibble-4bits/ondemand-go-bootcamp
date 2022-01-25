package httpAPI

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/nibble-4bits/ondemand-go-bootcamp/adapter"
	"github.com/nibble-4bits/ondemand-go-bootcamp/usecase"

	"github.com/gin-gonic/gin"
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
		switch {
		case errors.Is(err, adapter.ErrPokemonNotFoundByID):
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, pokemon)
	}
}

func getAllPokemonsController(service usecase.PokemonService) gin.HandlerFunc {
	return func(c *gin.Context) {
		pokemons, err := service.GetAll()
		switch {
		case errors.Is(err, adapter.ErrPokemonsNotFound):
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, pokemons)
	}
}
