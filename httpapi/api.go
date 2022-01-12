package httpapi

import (
	"github.com/gin-gonic/gin"
	"github.com/nibble-4bits/ondemand-go-bootcamp/usecase"
)

func StartServer(service usecase.PokemonService) {
	r := gin.Default()

	registerRoutes(r, service)

	r.Run()
}
