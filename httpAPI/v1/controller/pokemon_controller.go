package controller

import (
	"errors"
	"math"
	"net/http"
	"strconv"

	"github.com/nibble-4bits/ondemand-go-bootcamp/adapter"
	"github.com/nibble-4bits/ondemand-go-bootcamp/usecase"

	"github.com/gin-gonic/gin"
)

// GetPokemonByID handles the logic of getting a pokemon by ID from a data source and sending it to an http client
// as a JSON object.
//
// If an error occurs, it will send that instead.
func GetPokemonByID(service usecase.PokemonService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		pokemon, err := service.GetByID(id)
		if err != nil {
			switch {
			case errors.Is(err, adapter.ErrPokemonNotFoundByID):
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
		}

		c.JSON(http.StatusOK, pokemon)
	}
}

// GetAllPokemons handles the logic of getting all pokemons from a data source and sending it to an http client
// as a JSON array.
//
// If an error occurs, it will send that instead.
func GetAllPokemons(service usecase.PokemonService) gin.HandlerFunc {
	return func(c *gin.Context) {
		pokemons, err := service.GetAll()
		if err != nil {
			switch {
			case errors.Is(err, adapter.ErrPokemonsNotFound):
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
		}

		c.JSON(http.StatusOK, pokemons)
	}
}

func GetEvenOrOddPokemons(service usecase.PokemonService) gin.HandlerFunc {
	return func(c *gin.Context) {
		parity := c.Param("parity")
		itemCountStr := c.Param("items")
		itemsPerWorkerStr := c.Param("items_per_worker")

		itemCount, err := strconv.Atoi(itemCountStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		itemsPerWorker, err := strconv.Atoi(itemsPerWorkerStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		workers := math.Ceil(float64(itemCount) / float64(itemsPerWorker))
		pokemons, err := service.GetByParity(parity, int(workers), itemCount, itemsPerWorker)
		if err != nil {
			switch {
			case errors.Is(err, adapter.ErrPokemonsNotFound):
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			case errors.Is(err, usecase.ErrUnsupportedParityType):
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
		}

		c.JSON(http.StatusOK, pokemons)
	}
}
