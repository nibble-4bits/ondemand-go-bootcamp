package router

import (
	"log"

	"github.com/nibble-4bits/ondemand-go-bootcamp/usecase"

	"github.com/gin-gonic/gin"
)

// StartServer makes the HTTP server start listening for requests
func StartServer(service usecase.PokemonService) {
	router := gin.Default()

	v1 := router.Group("/v1")

	registerPokemonRoutes(v1, service)

	err := router.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
