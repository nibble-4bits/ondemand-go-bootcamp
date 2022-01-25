package router

import (
	"log"

	"github.com/nibble-4bits/ondemand-go-bootcamp/usecase"

	"github.com/gin-gonic/gin"
)

func StartServer(service usecase.PokemonService) {
	r := gin.Default()

	registerRoutes(r, service)

	err := r.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
